package stackpath

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-stackpath/stackpath/api/object_storage/client/buckets"
	"github.com/terraform-providers/terraform-provider-stackpath/stackpath/api/object_storage/models"
)

// Create bucket and update visibility
func TestObjectStorageBucketBasic(t *testing.T) {

	bucket := &models.GetBucketOKBodyBucket{}
	labelSuffix := strconv.Itoa(int(time.Now().Unix()))

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck: func() {
			testAccPreCheck(t)
		},
		CheckDestroy: testAccObjectStorageBucketCheckDestroy(),
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testObjectStorageBucketBasic(labelSuffix, "us-east-2"),
				Check: resource.ComposeTestCheckFunc(
					testAccObjectStorageBucketCheckExists("stackpath_object_storage_bucket.bucket", bucket),
					resource.TestCheckResourceAttr("stackpath_object_storage_bucket.bucket", "label", fmt.Sprintf("acc-test-%s", labelSuffix)),
					resource.TestCheckResourceAttr("stackpath_object_storage_bucket.bucket", "region", "us-east-2"),
					resource.TestCheckResourceAttr("stackpath_object_storage_bucket.bucket", "visibility", "PRIVATE"),
					resource.TestCheckResourceAttr("stackpath_object_storage_bucket.bucket", "endpoint_url", "https://s3.us-east-2.stackpathstorage.com"),
				),
			},
			resource.TestStep{
				Config: testObjectStorageBucketPublic(labelSuffix, "us-east-2"),
				Check: resource.ComposeTestCheckFunc(
					testAccObjectStorageBucketCheckExists("stackpath_object_storage_bucket.bucket", bucket),
					resource.TestCheckResourceAttr("stackpath_object_storage_bucket.bucket", "label", fmt.Sprintf("acc-test-%s", labelSuffix)),
					resource.TestCheckResourceAttr("stackpath_object_storage_bucket.bucket", "region", "us-east-2"),
					resource.TestCheckResourceAttr("stackpath_object_storage_bucket.bucket", "visibility", "PUBLIC"),
					resource.TestCheckResourceAttr("stackpath_object_storage_bucket.bucket", "endpoint_url", "https://s3.us-east-2.stackpathstorage.com"),
				),
			},
		},
	})
}

func TestObjectStorageBucketBasicRegionChange(t *testing.T) {

	bucket1 := &models.GetBucketOKBodyBucket{}
	bucket2 := &models.GetBucketOKBodyBucket{}
	labelSuffix := strconv.Itoa(int(time.Now().Unix()))

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck: func() {
			testAccPreCheck(t)
		},
		CheckDestroy: testAccObjectStorageBucketCheckDestroy(),
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testObjectStorageBucketBasic(labelSuffix, "us-east-2"),
				Check: resource.ComposeTestCheckFunc(
					testAccObjectStorageBucketCheckExists("stackpath_object_storage_bucket.bucket", bucket1),
					resource.TestCheckResourceAttr("stackpath_object_storage_bucket.bucket", "label", fmt.Sprintf("acc-test-%s", labelSuffix)),
					resource.TestCheckResourceAttr("stackpath_object_storage_bucket.bucket", "region", "us-east-2"),
					resource.TestCheckResourceAttr("stackpath_object_storage_bucket.bucket", "visibility", "PRIVATE"),
					resource.TestCheckResourceAttr("stackpath_object_storage_bucket.bucket", "endpoint_url", "https://s3.us-east-2.stackpathstorage.com"),
				),
			},
			resource.TestStep{
				Config: testObjectStorageBucketBasic(labelSuffix, "us-west-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccObjectStorageBucketCheckDestroyed(bucket1),
					testAccObjectStorageBucketCheckExists("stackpath_object_storage_bucket.bucket", bucket2),
					resource.TestCheckResourceAttr("stackpath_object_storage_bucket.bucket", "label", fmt.Sprintf("acc-test-%s", labelSuffix)),
					resource.TestCheckResourceAttr("stackpath_object_storage_bucket.bucket", "region", "us-west-1"),
					resource.TestCheckResourceAttr("stackpath_object_storage_bucket.bucket", "visibility", "PRIVATE"),
					resource.TestCheckResourceAttr("stackpath_object_storage_bucket.bucket", "endpoint_url", "https://s3.us-west.stackpathstorage.com"),
				),
			},
		},
	})
}

func TestObjectStorageBucketBasicLabelChange(t *testing.T) {

	bucket1 := &models.GetBucketOKBodyBucket{}
	bucket2 := &models.GetBucketOKBodyBucket{}
	labelSuffix1 := strconv.Itoa(int(time.Now().Unix()))
	labelSuffix2 := strconv.Itoa(int(time.Now().Unix()) + 1)

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck: func() {
			testAccPreCheck(t)
		},
		CheckDestroy: testAccObjectStorageBucketCheckDestroy(),
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testObjectStorageBucketBasic(labelSuffix1, "us-east-2"),
				Check: resource.ComposeTestCheckFunc(
					testAccObjectStorageBucketCheckExists("stackpath_object_storage_bucket.bucket", bucket1),
					resource.TestCheckResourceAttr("stackpath_object_storage_bucket.bucket", "label", fmt.Sprintf("acc-test-%s", labelSuffix1)),
					resource.TestCheckResourceAttr("stackpath_object_storage_bucket.bucket", "region", "us-east-2"),
					resource.TestCheckResourceAttr("stackpath_object_storage_bucket.bucket", "visibility", "PRIVATE"),
					resource.TestCheckResourceAttr("stackpath_object_storage_bucket.bucket", "endpoint_url", "https://s3.us-east-2.stackpathstorage.com"),
				),
			},
			resource.TestStep{
				Config: testObjectStorageBucketBasic(labelSuffix2, "us-east-2"),
				Check: resource.ComposeTestCheckFunc(
					testAccObjectStorageBucketCheckDestroyed(bucket1),
					testAccObjectStorageBucketCheckExists("stackpath_object_storage_bucket.bucket", bucket2),
					resource.TestCheckResourceAttr("stackpath_object_storage_bucket.bucket", "label", fmt.Sprintf("acc-test-%s", labelSuffix2)),
					resource.TestCheckResourceAttr("stackpath_object_storage_bucket.bucket", "region", "us-east-2"),
					resource.TestCheckResourceAttr("stackpath_object_storage_bucket.bucket", "visibility", "PRIVATE"),
					resource.TestCheckResourceAttr("stackpath_object_storage_bucket.bucket", "endpoint_url", "https://s3.us-east-2.stackpathstorage.com"),
				),
			},
		},
	})
}

func testAccObjectStorageBucketCheckExists(name string, bucket *models.GetBucketOKBodyBucket) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("resource not found: %s: available resources: %v", name, s.RootModule().Resources)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID set: %s", name)
		}

		config := testAccProvider.Meta().(*Config)
		resp, err := config.objectStorage.Buckets.GetBucket(&buckets.GetBucketParams{
			BucketID: rs.Primary.ID,
			StackID:  config.StackID,
			Context:  context.Background(),
		}, nil)
		if err != nil {
			return fmt.Errorf("Could not retrieve object storage bucket: %v", err)
		}
		bucket = resp.GetPayload().Bucket
		return nil
	}
}

func testAccObjectStorageBucketCheckDestroyed(bucket *models.GetBucketOKBodyBucket) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config := testAccProvider.Meta().(*Config)
		_, err := config.objectStorage.Buckets.GetBucket(&buckets.GetBucketParams{
			BucketID: bucket.ID,
			StackID:  config.StackID,
			Context:  context.Background(),
		}, nil)
		if err == nil {
			return fmt.Errorf("Bucket still exists")
		}
		return nil
	}
}

func testAccObjectStorageBucketCheckDestroy() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config := testAccProvider.Meta().(*Config)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "stackpath_object_storage_bucket" {
				continue
			}
			_, err := config.objectStorage.Buckets.GetBucket(&buckets.GetBucketParams{
				BucketID: rs.Primary.ID,
				StackID:  config.StackID,
				Context:  context.Background(),
			}, nil)
			if c, ok := err.(interface{ Code() int }); ok && c.Code() != http.StatusNotFound {
				return fmt.Errorf("Object storage bucket still exists: %v HTTP %d", rs.Primary.ID, c.Code())
			}
		}
		return nil
	}
}

func testObjectStorageBucketBasic(suffix string, region string) string {
	return fmt.Sprintf(`
resource "stackpath_object_storage_bucket" "bucket" {
	label = "acc-test-%s"
	region = "%s"
}
`, suffix, region)
}

func testObjectStorageBucketPublic(suffix string, region string) string {
	return fmt.Sprintf(`
resource "stackpath_object_storage_bucket" "bucket" {
	label = "acc-test-%s"
	region = "%s"
	visibility = "PUBLIC"
}
`, suffix, region)
}
