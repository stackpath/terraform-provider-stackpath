package stackpath

import (
	"context"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stackpath/terraform-provider-stackpath/stackpath/api/workload/workload_client/workloads"
	"github.com/stackpath/terraform-provider-stackpath/stackpath/api/workload/workload_models"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

var (
	IPv4IPFamilies      = []string{"IPv4"}
	DualStackIPFamilies = []string{"IPv4", "IPv6"}
)

func TestComputeWorkloadContainers(t *testing.T) {
	t.Parallel()

	workload := &workload_models.V1Workload{}
	nameSuffix := strconv.Itoa(int(time.Now().Unix()))

	// By design, the StackPath API does not return image pull secrets to the
	// user when retrieving a workload. Expect to see an empty secret in the
	// result and suppress the diff error.
	//emptyImagePullSecrets := regexp.MustCompile("(.*)image_pull_credentials.0.docker_registry.0.password:(\\s*)\"\" => \"secret registry password\"(.*)")

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		PreCheck: func() {
			testAccPreCheck(t)
		},
		CheckDestroy: testAccComputeWorkloadCheckDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testComputeWorkloadConfigContainerBasic(nameSuffix, nil),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckContainerPort(workload, "app", "http", "TCP", 80, false),
					testAccComputeWorkloadCheckContainerEnvVar(workload, "app", "MY_ENVIRONMENT_VARIABLE", "value"),
					testAccComputeWorkloadCheckTarget(workload, "us", "cityCode", "in", 1, "AMS"),
					testAccComputeWorkloadCheckInterface(workload, 0, "default", true, "", "", IPv4IPFamilies),
				),
			},
			{
				Config: testComputeWorkloadConfigContainerAddPorts(nameSuffix, nil),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckContainerPort(workload, "app", "http", "TCP", 80, false),
					testAccComputeWorkloadCheckContainerPort(workload, "app", "https", "TCP", 443, true),
					testAccComputeWorkloadCheckContainerEnvVar(workload, "app", "MY_ENVIRONMENT_VARIABLE", "some value"),
					testAccComputeWorkloadCheckTarget(workload, "us", "cityCode", "in", 2, "AMS"),
					testAccComputeWorkloadCheckInterface(workload, 0, "default", true, "", "", IPv4IPFamilies),
				),
			},
			{
				Config: testComputeWorkloadConfigContainerRemoveEnvVar(nameSuffix, nil),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckContainerPort(workload, "app", "http", "TCP", 80, false),
					testAccComputeWorkloadCheckContainerPortNotExist(workload, "app", "https"),
					testAccComputeWorkloadCheckContainerEnvVarNotExist(workload, "app", "MY_ENVIRONMENT_VARIABLE"),
					testAccComputeWorkloadCheckTarget(workload, "us", "cityCode", "in", 2, "AMS"),
					testAccComputeWorkloadCheckInterface(workload, 0, "default", true, "", "", IPv4IPFamilies),
				),
			},
			{
				Config: testComputeWorkloadConfigContainerAddProbes(nameSuffix, nil),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckContainerPort(workload, "app", "http", "TCP", 80, false),
					testAccComputeWorkloadCheckContainerPortNotExist(workload, "app", "https"),
					testAccComputeWorkloadCheckContainerEnvVarNotExist(workload, "app", "MY_ENVIRONMENT_VARIABLE"),
					testAccComputeWorkloadCheckTarget(workload, "us", "cityCode", "in", 2, "AMS"),
					testAccComputeWorkloadCheckInterface(workload, 0, "default", true, "", "", IPv4IPFamilies),
				),
			},
			{
				ExpectNonEmptyPlan: true,
				Config:             testComputeWorkloadConfigContainerImagePullCredentials(nameSuffix, nil),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckImagePullCredentials(workload, "docker.io", "my-registry-user", "developers@stackpath.com"),
					testAccComputeWorkloadCheckInterface(workload, 0, "default", true, "", "", IPv4IPFamilies),
				),
			},
			{
				ExpectNonEmptyPlan: true,
				Config:             testComputeWorkloadConfigAutoScalingConfiguration(nameSuffix, nil),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckTargetAutoScaling(workload, "us", "cpu", 2, 4, 50),
					testAccComputeWorkloadCheckInterface(workload, 0, "default", true, "", "", IPv4IPFamilies),
				),
			},
			// TODO: there's a ordering issue where the order of the containers is shuffled when being read in from the API
			//   Need to ensure consistent ordering of containers when reading in state.
			//
			// {
			// 	Config: testComputeWorkloadConfigContainerAddContainer(),
			// 	Check: resource.ComposeTestCheckFunc(
			// 		testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
			// 		testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
			// 		testAccComputeWorkloadCheckContainerImage(workload, "app-2", "nginx:v1.15.0"),
			// 	),
			// },
		},
	})
}

func TestComputeWorkloadContainersWithOneToOneNATEnabled(t *testing.T) {
	t.Parallel()

	workload := &workload_models.V1Workload{}
	nameSuffix := "nat-" + strconv.Itoa(int(time.Now().Unix()))
	oneToOneNAT := true

	// By design, the StackPath API does not return image pull secrets to the
	// user when retrieving a workload. Expect to see an empty secret in the
	// result and suppress the diff error.
	//emptyImagePullSecrets := regexp.MustCompile("(.*)image_pull_credentials.0.docker_registry.0.password:(\\s*)\"\" => \"secret registry password\"(.*)")

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		PreCheck: func() {
			testAccPreCheck(t)
		},
		CheckDestroy: testAccComputeWorkloadCheckDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testComputeWorkloadConfigContainerBasic(nameSuffix, &oneToOneNAT),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckContainerPort(workload, "app", "http", "TCP", 80, false),
					testAccComputeWorkloadCheckContainerEnvVar(workload, "app", "MY_ENVIRONMENT_VARIABLE", "value"),
					testAccComputeWorkloadCheckTarget(workload, "us", "cityCode", "in", 1, "AMS"),
					testAccComputeWorkloadCheckInterface(workload, 0, "default", true, "", "", IPv4IPFamilies),
				),
			},
			{
				Config: testComputeWorkloadConfigContainerAddPorts(nameSuffix, &oneToOneNAT),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckContainerPort(workload, "app", "http", "TCP", 80, false),
					testAccComputeWorkloadCheckContainerPort(workload, "app", "https", "TCP", 443, true),
					testAccComputeWorkloadCheckContainerEnvVar(workload, "app", "MY_ENVIRONMENT_VARIABLE", "some value"),
					testAccComputeWorkloadCheckTarget(workload, "us", "cityCode", "in", 2, "AMS"),
					testAccComputeWorkloadCheckInterface(workload, 0, "default", true, "", "", IPv4IPFamilies),
				),
			},
			{
				Config: testComputeWorkloadConfigContainerRemoveEnvVar(nameSuffix, &oneToOneNAT),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckContainerPort(workload, "app", "http", "TCP", 80, false),
					testAccComputeWorkloadCheckContainerPortNotExist(workload, "app", "https"),
					testAccComputeWorkloadCheckContainerEnvVarNotExist(workload, "app", "MY_ENVIRONMENT_VARIABLE"),
					testAccComputeWorkloadCheckTarget(workload, "us", "cityCode", "in", 2, "AMS"),
					testAccComputeWorkloadCheckInterface(workload, 0, "default", true, "", "", IPv4IPFamilies),
				),
			},
			{
				Config: testComputeWorkloadConfigContainerAddProbes(nameSuffix, &oneToOneNAT),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckContainerPort(workload, "app", "http", "TCP", 80, false),
					testAccComputeWorkloadCheckContainerPortNotExist(workload, "app", "https"),
					testAccComputeWorkloadCheckContainerEnvVarNotExist(workload, "app", "MY_ENVIRONMENT_VARIABLE"),
					testAccComputeWorkloadCheckTarget(workload, "us", "cityCode", "in", 2, "AMS"),
					testAccComputeWorkloadCheckInterface(workload, 0, "default", true, "", "", IPv4IPFamilies),
				),
			},
			{
				ExpectNonEmptyPlan: true,
				Config:             testComputeWorkloadConfigContainerImagePullCredentials(nameSuffix, &oneToOneNAT),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckImagePullCredentials(workload, "docker.io", "my-registry-user", "developers@stackpath.com"),
					testAccComputeWorkloadCheckInterface(workload, 0, "default", true, "", "", IPv4IPFamilies),
				),
			},
			{
				ExpectNonEmptyPlan: true,
				Config:             testComputeWorkloadConfigAutoScalingConfiguration(nameSuffix, &oneToOneNAT),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckTargetAutoScaling(workload, "us", "cpu", 2, 4, 50),
					testAccComputeWorkloadCheckInterface(workload, 0, "default", true, "", "", IPv4IPFamilies),
				),
			},
			// TODO: there's a ordering issue where the order of the containers is shuffled when being read in from the API
			//   Need to ensure consistent ordering of containers when reading in state.
			//
			// {
			// 	Config: testComputeWorkloadConfigContainerAddContainer(),
			// 	Check: resource.ComposeTestCheckFunc(
			// 		testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
			// 		testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
			// 		testAccComputeWorkloadCheckContainerImage(workload, "app-2", "nginx:v1.15.0"),
			// 	),
			// },
		},
	})
}

func TestComputeWorkloadContainersAdditionalVolume(t *testing.T) {
	t.Parallel()

	workload := &workload_models.V1Workload{}
	nameSuffix := strconv.Itoa(int(time.Now().Unix()))

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		PreCheck: func() {
			testAccPreCheck(t)
		},
		CheckDestroy: testAccComputeWorkloadCheckDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testComputeWorkloadConfigContainerAddVolumeMounts(nameSuffix, nil),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo-volume", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadAdditionalVolume(workload, "volume", "10Gi"),
					testAccComputeWorkloadContainerVolumeMount(workload, "app", "volume", "/var/log"),
				),
			},
		},
	})
}

func TestComputeWorkloadContainersIPv4Only(t *testing.T) {
	t.Parallel()

	workload := &workload_models.V1Workload{}
	nameSuffix := "ipv4-" + strconv.Itoa(int(time.Now().Unix()))

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		PreCheck: func() {
			testAccPreCheck(t)
		},
		CheckDestroy: testAccComputeWorkloadCheckDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testComputeWorkloadConfigContainerForIPFamilies(nameSuffix, "default", "", "", IPv4IPFamilies),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckContainerPort(workload, "app", "http", "TCP", 80, false),
					testAccComputeWorkloadCheckContainerEnvVar(workload, "app", "MY_ENVIRONMENT_VARIABLE", "value"),
					testAccComputeWorkloadCheckTarget(workload, "us", "cityCode", "in", 1, "AMS"),
					testAccComputeWorkloadCheckInterface(workload, 0, "default", true, "", "", IPv4IPFamilies),
				),
			},
		},
	})
}

func TestComputeWorkloadContainersIPv6DualStack(t *testing.T) {
	t.Parallel()

	workload := &workload_models.V1Workload{}
	nameSuffix := "ipv6-" + strconv.Itoa(int(time.Now().Unix()))

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		PreCheck: func() {
			testAccPreCheck(t)
		},
		CheckDestroy: testAccComputeWorkloadCheckDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testComputeWorkloadConfigContainerForIPFamilies(nameSuffix, "default", "", "", DualStackIPFamilies),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckContainerPort(workload, "app", "http", "TCP", 80, false),
					testAccComputeWorkloadCheckContainerEnvVar(workload, "app", "MY_ENVIRONMENT_VARIABLE", "value"),
					testAccComputeWorkloadCheckTarget(workload, "us", "cityCode", "in", 1, "AMS"),
					testAccComputeWorkloadCheckInterface(workload, 0, "default", true, "", "", DualStackIPFamilies),
				),
			},
		},
	})
}

func TestComputeWorkloadContainersIPv6DualStackWithSubnets(t *testing.T) {
	t.Parallel()

	workload := &workload_models.V1Workload{}
	nameSuffix := "ipv6-" + strconv.Itoa(int(time.Now().Unix()))

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		PreCheck: func() {
			testAccPreCheck(t)
		},
		CheckDestroy: testAccComputeWorkloadCheckDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testComputeWorkloadConfigContainerForIPFamilies(nameSuffix, "tf-test-dual-vpc", "subnet-ipv4", "subnet-ipv6", DualStackIPFamilies),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckContainerPort(workload, "app", "http", "TCP", 80, false),
					testAccComputeWorkloadCheckContainerEnvVar(workload, "app", "MY_ENVIRONMENT_VARIABLE", "value"),
					testAccComputeWorkloadCheckTarget(workload, "us", "cityCode", "in", 1, "AMS"),
					testAccComputeWorkloadCheckInterface(workload, 0, "tf-test-dual-vpc", true, "subnet-ipv4", "subnet-ipv6", DualStackIPFamilies),
				),
			},
		},
	})
}

func TestComputeContainersEnhancedContainerControls(t *testing.T) {
	t.Parallel()

	workload := &workload_models.V1Workload{}
	nameSuffix := "ecc-" + strconv.Itoa(int(time.Now().Unix()))

	// By design, the StackPath API does not return image pull secrets to the
	// user when retrieving a workload. Expect to see an empty secret in the
	// result and suppress the diff error.
	//emptyImagePullSecrets := regexp.MustCompile("(.*)image_pull_credentials.0.docker_registry.0.password:(\\s*)\"\" => \"secret registry password\"(.*)")

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		PreCheck: func() {
			testAccPreCheck(t)
		},
		CheckDestroy: testAccComputeWorkloadCheckDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testComputeWorkloadConfigContainerBasic(nameSuffix, nil),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckContainerPort(workload, "app", "http", "TCP", 80, false),
					testAccComputeWorkloadCheckContainerEnvVar(workload, "app", "MY_ENVIRONMENT_VARIABLE", "value"),
					testAccComputeWorkloadCheckTarget(workload, "us", "cityCode", "in", 1, "AMS"),
					testAccComputeWorkloadCheckInterface(workload, 0, "default", true, "", "", IPv4IPFamilies),
				),
			},
			{
				Config: testComputeWorkloadConfigContainerSecurityContextCapabilities(nameSuffix, nil),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckCapabilities(workload, "app", []string{"NET_ADMIN"}, []string{"NET_BROADCAST"}),
					testAccComputeWorkloadCheckSecurityContext(workload, "app", true /*priv */, false /*ro*/, true /*nonroot*/, "101", ""),
				),
			},
			{
				ExpectNonEmptyPlan: true, // This flag is confusing
				Config:             testComputeWorkloadConfigContainerSecurityContextClearCapabilities(nameSuffix, nil),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckCapabilities(workload, "app", []string{}, []string{}),
					testAccComputeWorkloadCheckSecurityContext(workload, "app", true /*priv */, false /*ro*/, true /*nonroot*/, "101", ""),
				),
			},
			{
				ExpectNonEmptyPlan: true,
				Config:             testComputeWorkloadConfigContainerRuntimeSettings(nameSuffix, nil),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckContainerRuntimeExists(workload, true),
					testAccComputeWorkloadCheckContainerRuntimeSecurityContext(workload,
						"60",  // termination,
						true,  // share namespace
						"999", // run_as_user
						"991", // run_as_group
						true,  // run_as_non_root
						[]string{"42"},
					),
					testAccComputeWorkloadCheckRuntimeSysctl(workload, map[string]string{
						"net.core.rmem_max":     "10065408",
						"net.core.rmem_default": "1006540",
					}),
					testAccComputeWorkloadCheckRuntimeHostAliases(workload,
						map[string][]string{
							"192.168.3.4": {"domain.com"},
						}),
					testAccComputeWorkloadCheckRuntimeDNSConfig(workload,
						[]string{"8.8.8.8"},
						[]string{"domain.com"},
						map[string]string{
							"timeout": "10",
						},
					),
				),
			},
			{
				ExpectNonEmptyPlan: true,
				Config:             testComputeWorkloadConfigContainerRuntimeClearSettings(nameSuffix, nil),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.foo", workload),
					testAccComputeWorkloadCheckContainerImage(workload, "app", "nginx:latest"),
					testAccComputeWorkloadCheckContainerRuntimeExists(workload, true),
					testAccComputeWorkloadCheckContainerRuntimeSecurityContext(workload,
						"60",           // termination,
						true,           // share namespace
						"999",          // run_as_user
						"991",          // run_as_group
						true,           // run_as_non_root
						[]string{"43"}, // make sure it actually changed
					),
					testAccComputeWorkloadCheckRuntimeSysctl(workload, map[string]string{
						"net.core.rmem_max":     "10065408",
						"net.core.rmem_default": "1006540",
					}),
					// clearing host settings
					testAccComputeWorkloadCheckRuntimeHostAliases(workload,
						map[string][]string{}),
					// cleared options
					testAccComputeWorkloadCheckRuntimeDNSConfig(workload,
						[]string{"8.8.8.8"},
						[]string{"domain2.com"}, // changed to verify we made a change
						nil,
					),
				),
			},
		},
	})

}

func TestComputeWorkloadVirtualMachines(t *testing.T) {
	t.Parallel()

	workload := &workload_models.V1Workload{}
	nameSuffix := strconv.Itoa(int(time.Now().Unix()))

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		PreCheck: func() {
			testAccPreCheck(t)
		},
		CheckDestroy: testAccComputeWorkloadCheckDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testComputeWorkloadConfigVirtualMachineBasic(nameSuffix, nil),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.bar", workload),
					testAccComputeWorkloadCheckVirtualMachineImage(workload, "app", "stackpath-edge/centos-7:v201905012051"),
					testAccComputeWorkloadCheckVirtualMachinePort(workload, "app", "http", "TCP", 80),
					testAccComputeWorkloadCheckTarget(workload, "us", "cityCode", "in", 1, "AMS"),
					testAccComputeWorkloadCheckInterface(workload, 0, "default", true, "", "", IPv4IPFamilies),
				),
			},
		},
	})
}

func TestComputeWorkloadVirtualMachinesIPv6DualStack(t *testing.T) {
	t.Parallel()

	workload := &workload_models.V1Workload{}
	nameSuffix := "ipv6-" + strconv.Itoa(int(time.Now().Unix()))

	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		PreCheck: func() {
			testAccPreCheck(t)
		},
		CheckDestroy: testAccComputeWorkloadCheckDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testComputeWorkloadConfigVirtualMachineForIPFamilies(nameSuffix, "", "", DualStackIPFamilies),
				Check: resource.ComposeTestCheckFunc(
					testAccComputeWorkloadCheckExists("stackpath_compute_workload.bar", workload),
					testAccComputeWorkloadCheckVirtualMachineImage(workload, "app", "stackpath-edge/centos-7:v201905012051"),
					testAccComputeWorkloadCheckVirtualMachinePort(workload, "app", "http", "TCP", 80),
					testAccComputeWorkloadCheckTarget(workload, "us", "cityCode", "in", 1, "AMS"),
					testAccComputeWorkloadCheckInterface(workload, 0, "default", true, "", "", DualStackIPFamilies),
				),
			},
		},
	})
}

func testAccComputeWorkloadCheckDestroy() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config := testAccProvider.Meta().(*Config)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "stackpath_compute_workload" {
				continue
			}

			resp, err := config.edgeCompute.Workloads.GetWorkload(&workloads.GetWorkloadParams{
				StackID:    config.StackID,
				WorkloadID: rs.Primary.ID,
				Context:    context.Background(),
			}, nil)
			// Since compute workloads are deleted asynchronously, we want to look at the fact that
			// the deleteRequestedAt timestamp was set on the workload. This field is used to indicate
			// that the workload is being deleted.
			if err == nil && resp.Payload.Workload.Metadata.DeleteRequestedAt == nil {
				return fmt.Errorf("compute workload still exists: %v", rs.Primary.ID)
			}
		}

		return nil
	}
}

func testAccComputeWorkloadCheckInterface(
	workload *workload_models.V1Workload,
	interfaceIndex int,
	networkName string,
	enableOneToOneNAT bool,
	ipv4Subnet string,
	ipv6Subnet string,
	ipFamilies []string,
) resource.TestCheckFunc {
	return func(_ *terraform.State) error {
		interfaces := workload.Spec.NetworkInterfaces
		if interfaceIndex < 0 {
			return fmt.Errorf("invalid interface index to check: %d", interfaceIndex)
		}
		if interfaceIndex >= len(interfaces) {
			return fmt.Errorf("could not find the interface index %d/%d", interfaceIndex, len(interfaces))
		}
		inter := interfaces[interfaceIndex]
		if inter.Network != networkName {
			return fmt.Errorf("invalid network on interface %d. got=%s want=%s", interfaceIndex, inter.Network, networkName)
		}
		if inter.EnableOneToOneNat != enableOneToOneNAT {
			return fmt.Errorf("invalid enableOneToOneNat on interface %d. got=%v want=%v", interfaceIndex, inter.EnableOneToOneNat, enableOneToOneNAT)
		}
		if inter.Subnet != ipv4Subnet {
			return fmt.Errorf("invalid subnet on interface %d. got=%v want=%v", interfaceIndex, inter.Subnet, ipv4Subnet)
		}
		if inter.IPV6Subnet != ipv6Subnet {
			return fmt.Errorf("invalid ipv6Subnet on interface %d. got=%v want=%v", interfaceIndex, inter.IPV6Subnet, ipv6Subnet)
		}
		if len(inter.IPFamilies) > 0 {
			ipFamiliesStrList := make([]string, len(inter.IPFamilies))
			for i, ipFamily := range inter.IPFamilies {
				ipFamiliesStrList[i] = string(*ipFamily)
			}
			if !reflect.DeepEqual(ipFamiliesStrList, ipFamilies) {
				return fmt.Errorf("invalid ipFamilies on interface %d. got=%v want=%v", interfaceIndex, ipFamiliesStrList, ipFamilies)
			}
		}
		return nil
	}
}

func testAccComputeWorkloadContainerVolumeMount(workload *workload_models.V1Workload, containerName, volumeSlug, mountPath string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		container, found := workload.Spec.Containers[containerName]
		if !found {
			return fmt.Errorf("container not found: %s", containerName)
		}
		var mount *workload_models.V1InstanceVolumeMount
		for _, v := range container.VolumeMounts {
			if v.Slug == volumeSlug {
				mount = v
				break
			}
		}
		if mount == nil {
			return fmt.Errorf("container volume mount not found: %s", volumeSlug)
		} else if mount.MountPath != mountPath {
			return fmt.Errorf("mount path '%s' does not match expected '%s'", mount.MountPath, mountPath)
		}
		return nil
	}
}

func testAccComputeWorkloadAdditionalVolume(workload *workload_models.V1Workload, volumeName, size string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		var volume *workload_models.V1VolumeClaim
		for _, v := range workload.Spec.VolumeClaimTemplates {
			if v.Name == volumeName {
				volume = v
				break
			}
		}
		if volume == nil {
			return fmt.Errorf("could not find volume: %s", volumeName)
		} else if volume.Spec.Resources.Requests["storage"] != size {
			return fmt.Errorf("volume size '%s' does not match expected '%s'", volume.Spec.Resources.Requests["storage"], size)
		}
		return nil
	}
}

func testAccComputeWorkloadCheckVirtualMachinePort(workload *workload_models.V1Workload, vmName, portName, protocol string, portNum int32) resource.TestCheckFunc {
	return func(*terraform.State) error {
		if vm, found := workload.Spec.VirtualMachines[vmName]; !found {
			return fmt.Errorf("virtual machine was not found: %s", vmName)
		} else if port, found := vm.Ports[portName]; !found {
			return fmt.Errorf("virtual machine port not found: %s", portName)
		} else if port.Protocol != protocol {
			return fmt.Errorf("virtual machine port protocol '%s' does not match expected '%s'", port.Protocol, protocol)
		} else if port.Port != portNum {
			return fmt.Errorf("virtual machine port '%d' does not match expected '%d'", port.Port, portNum)
		}
		return nil
	}
}

func testAccComputeWorkloadCheckVirtualMachineImage(workload *workload_models.V1Workload, name, image string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		if vm, found := workload.Spec.VirtualMachines[name]; !found {
			return fmt.Errorf("virtual machine was not found: %s", name)
		} else if vm.Image != image {
			return fmt.Errorf("virtual machine image '%s' does not match expected '%s'", vm.Image, image)
		}
		return nil
	}
}

func testAccComputeWorkloadCheckNoImagePullCredentials(workload *workload_models.V1Workload) resource.TestCheckFunc {
	return func(*terraform.State) error {
		if workload.Spec.ImagePullCredentials != nil {
			return fmt.Errorf("unexpected image pull credentials set on the workload")
		}
		return nil
	}
}

func testAccComputeWorkloadCheckImagePullCredentials(workload *workload_models.V1Workload, server, username, email string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		if workload.Spec.ImagePullCredentials == nil {
			return fmt.Errorf("no image pull credentials set on the workload")
		} else if creds := workload.Spec.ImagePullCredentials[0]; creds.DockerRegistry.Server != server {
			return fmt.Errorf("image pull server '%s' does not match expected value '%s'", creds.DockerRegistry.Server, server)
		} else if creds.DockerRegistry.Email != email {
			return fmt.Errorf("image pull email '%s' does not match expected value '%s'", creds.DockerRegistry.Email, email)
		} else if creds.DockerRegistry.Username != username {
			return fmt.Errorf("image pull username '%s' does not match expected value '%s'", creds.DockerRegistry.Username, username)
		}
		return nil
	}
}

func testAccComputeWorkloadCheckTarget(workload *workload_models.V1Workload, targetName, scope, operator string, minReplicas int32, values ...string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		sort.Strings(values)
		if target, found := workload.Targets[targetName]; !found {
			return fmt.Errorf("target not found: %s", targetName)
		} else if target.Spec.DeploymentScope != scope {
			return fmt.Errorf("target scope '%s' does not match expected scope '%s'", target.Spec.DeploymentScope, scope)
		} else if deployment := target.Spec.Deployments; deployment.MinReplicas != minReplicas {
			return fmt.Errorf("target min replicas '%d' does not match expected '%d'", target.Spec.Deployments.MinReplicas, minReplicas)
		} else if selector := target.Spec.Deployments.Selectors[0]; selector.Operator != operator {
			return fmt.Errorf("target selector operator '%s' does not match expected operator '%s'", selector.Operator, operator)
		} else if sort.Strings(selector.Values); !reflect.DeepEqual(values, selector.Values) {
			return fmt.Errorf("target selector values %v do not match expected values %v", selector.Values, values)
		}
		return nil
	}
}

func testAccComputeWorkloadCheckContainerEnvVar(workload *workload_models.V1Workload, containerName, envVar, value string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		containerSpec, found := workload.Spec.Containers[containerName]
		if !found {
			return fmt.Errorf("container not found: %s", containerName)
		} else if envVarSpec, found := containerSpec.Env[envVar]; !found {
			return fmt.Errorf("environment variable not found: %s", envVar)
		} else if envVarSpec.Value != value {
			return fmt.Errorf(`environment variable '%s="%s"' does not match expected value '%s'`, envVar, envVarSpec.Value, value)
		}
		return nil
	}
}

func testAccComputeWorkloadCheckContainerEnvVarNotExist(workload *workload_models.V1Workload, containerName, envVar string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		containerSpec, found := workload.Spec.Containers[containerName]
		if !found {
			return fmt.Errorf("container not found: %s", containerName)
		} else if _, found := containerSpec.Env[envVar]; found {
			return fmt.Errorf("unexpected environment variable found: %s", envVar)
		}
		return nil
	}
}

func testAccComputeWorkloadCheckContainerPort(workload *workload_models.V1Workload, containerName, portName, protocol string, port int32, enableImplicitNetworkPolicy bool) resource.TestCheckFunc {
	return func(*terraform.State) error {
		containerSpec, found := workload.Spec.Containers[containerName]
		if !found {
			return fmt.Errorf("container not found: %s", containerName)
		} else if portSpec, found := containerSpec.Ports[portName]; !found {
			return fmt.Errorf("port not found: %s", portName)
		} else if portSpec.Port != port {
			return fmt.Errorf("port number '%d' does not match expected port '%d'", portSpec.Port, port)
		} else if portSpec.Protocol != protocol {
			return fmt.Errorf("port protocol '%s' does not match expected protocol '%s'", portSpec.Protocol, protocol)
		} else if portSpec.EnableImplicitNetworkPolicy != enableImplicitNetworkPolicy {
			return fmt.Errorf("port enable implicit network policy '%t' does not match expected enable implicit network policy '%t'", portSpec.EnableImplicitNetworkPolicy, enableImplicitNetworkPolicy)
		}
		return nil
	}
}

func testAccComputeWorkloadCheckContainerPortNotExist(workload *workload_models.V1Workload, containerName, portName string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		containerSpec, found := workload.Spec.Containers[containerName]
		if !found {
			return fmt.Errorf("container not found: %s", containerName)
		} else if _, found := containerSpec.Ports[portName]; found {
			return fmt.Errorf("unexpected port found: %s", portName)
		}
		return nil
	}
}

func testAccComputeWorkloadCheckContainerImage(workload *workload_models.V1Workload, containerName, image string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		if containerSpec, found := workload.Spec.Containers[containerName]; !found {
			return fmt.Errorf("container not found: %s", containerName)
		} else if containerSpec.Image != image {
			return fmt.Errorf("container image '%s' does not match expected '%s'", containerSpec.Image, image)
		}
		return nil
	}
}

func testAccComputeWorkloadCheckExists(name string, spec *workload_models.V1Workload) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("resource not found: %s: available resources: %v", name, s.RootModule().Resources)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID set: %s", name)
		}

		config := testAccProvider.Meta().(*Config)
		found, err := config.edgeCompute.Workloads.GetWorkload(&workloads.GetWorkloadParams{
			WorkloadID: rs.Primary.ID,
			StackID:    config.StackID,
			Context:    context.Background(),
		}, nil)
		if err != nil {
			return fmt.Errorf("could not retrieve workload: %v", err)
		}

		*spec = *found.Payload.Workload

		return nil
	}
}

func testAccComputeWorkloadCheckTargetAutoScaling(workload *workload_models.V1Workload, targetName, metric string, minReplicas, maxReplicas, averageUtilization int32) resource.TestCheckFunc {
	return func(*terraform.State) error {
		target, found := workload.Targets[targetName]
		if !found {
			return fmt.Errorf("target not found: %s", targetName)
		} else if target.Spec.Deployments.MinReplicas != minReplicas {
			return fmt.Errorf("expected %d min replicas, got %d", minReplicas, target.Spec.Deployments.MinReplicas)
		} else if target.Spec.Deployments.MaxReplicas != maxReplicas {
			return fmt.Errorf("expected %d max replicas, got %d", maxReplicas, target.Spec.Deployments.MaxReplicas)
		} else if target.Spec.Deployments.ScaleSettings == nil {
			return fmt.Errorf("expected scale settings to be non-nil")
		}

		for _, m := range target.Spec.Deployments.ScaleSettings.Metrics {
			if m.Metric != metric {
				continue
			} else if m.AverageUtilization != averageUtilization {
				return fmt.Errorf("expected average utilization to be %d, got %d", averageUtilization, m.AverageUtilization)
			}
		}
		return nil
	}
}

func testAccComputeWorkloadCheckCapabilities(workload *workload_models.V1Workload, containerName string, add []string, drop []string) resource.TestCheckFunc {

	return func(*terraform.State) error {
		container, found := workload.Spec.Containers[containerName]
		if !found {
			return fmt.Errorf("container not found: %s", containerName)
		} else if container.SecurityContext == nil {
			return fmt.Errorf("container '%s' does not contain a security context", containerName)
		} else if container.SecurityContext.Capabilities == nil && (add != nil || drop != nil) {
			return fmt.Errorf("container '%s' does not have capability requests: %+v", containerName, container.SecurityContext)
		} else if container.SecurityContext.Capabilities != nil && add == nil && drop == nil {
			return fmt.Errorf("container '%s' expected to have empty capabilities, but has object (%v %v)",
				containerName,
				container.SecurityContext.Capabilities.Add,
				container.SecurityContext.Capabilities.Drop,
			)
		} else if container.SecurityContext.Capabilities != nil && !reflect.DeepEqual(container.SecurityContext.Capabilities.Add, add) {
			return fmt.Errorf("container '%s' ADD caps not the same (%v)!=(%v)",
				containerName, container.SecurityContext.Capabilities.Add, add)
		} else if container.SecurityContext.Capabilities != nil && !reflect.DeepEqual(container.SecurityContext.Capabilities.Drop, drop) {
			return fmt.Errorf("container '%s' DROP caps not the same (%v)!=(%v)",
				containerName, container.SecurityContext.Capabilities.Drop, drop)
		}

		return nil
	}
}

func testAccComputeWorkloadCheckSecurityContext(workload *workload_models.V1Workload, containerName string,
	allowPrivEscalation, readOnly, nonRoot bool, uid, gid string) resource.TestCheckFunc {

	return func(*terraform.State) error {
		container, found := workload.Spec.Containers[containerName]
		if !found {
			return fmt.Errorf("container not found: %s", containerName)
		} else if container.SecurityContext == nil {
			return fmt.Errorf("container '%s' does not contain a security context", containerName)

		} else if container.SecurityContext.AllowPrivilegeEscalation != allowPrivEscalation {
			return fmt.Errorf("container '%s' does not have correct allowPrivilegeEscalation", containerName)
		} else if container.SecurityContext.ReadOnlyRootFilesystem != readOnly {
			return fmt.Errorf("container '%s' does not have correct rootOnlyFilesystem", containerName)

		} else if container.SecurityContext.RunAsNonRoot != nonRoot {
			return fmt.Errorf("container '%s' does not have correct runAsNonRoot", containerName)
		} else if container.SecurityContext.RunAsUser != uid {
			return fmt.Errorf("container '%s' does not have correct runAsUser: %s", containerName, container.SecurityContext.RunAsUser)
		} else if container.SecurityContext.RunAsGroup != gid {
			return fmt.Errorf("container '%s' does not have correct runAsGroup: %s", containerName, container.SecurityContext.RunAsGroup)
		}

		return nil
	}
}

func testAccComputeWorkloadCheckContainerRuntimeExists(workload *workload_models.V1Workload, exists bool) resource.TestCheckFunc {
	return func(*terraform.State) error {
		if workload.Spec.Runtime == nil {
			return fmt.Errorf("expected runtime settings, but had none")
		} else if exists && workload.Spec.Runtime.Containers == nil {
			return fmt.Errorf("expected runtime container settings, but had none")
		} else if !exists && workload.Spec.Runtime.Containers != nil {
			return fmt.Errorf("did not expect to have container runtime settings, but not nil")
		}
		return nil
	}

}

// Precondition: testAccComputeWorkloadCheckContainerRuntimeExists has run
func testAccComputeWorkloadCheckContainerRuntimeSecurityContext(workload *workload_models.V1Workload,
	terminationGrace string, shareNamespace bool,
	runAsUser string, runAsGroup string, nonRoot bool,
	supplemental []string) resource.TestCheckFunc {

	return func(*terraform.State) error {
		if workload.Spec.Runtime.Containers.SecurityContext == nil {
			return fmt.Errorf("expected container runtime security context")
		}
		secContext := workload.Spec.Runtime.Containers.SecurityContext

		if secContext.RunAsUser != runAsUser {
			return fmt.Errorf("runtime: run_as_user %s != %s", secContext.RunAsUser, runAsUser)
		} else if secContext.RunAsGroup != runAsGroup {
			return fmt.Errorf("runtime: run_as_group %s != %s", secContext.RunAsGroup, runAsGroup)
		} else if secContext.RunAsNonRoot != nonRoot {
			return fmt.Errorf("runtime: run_as_non_root should be %v, but it was %v", nonRoot, secContext.RunAsNonRoot)
		} else if !reflect.DeepEqual(secContext.SupplementalGroups, supplemental) {
			return fmt.Errorf("runtime: supplemental groups, '%v' != '%v'", secContext.SupplementalGroups, supplemental)
		}
		return nil
	}
}

// This test has precondition of verifying correct stsate of the runtime settings themselves
func testAccComputeWorkloadCheckRuntimeSysctl(workload *workload_models.V1Workload, expected map[string]string) resource.TestCheckFunc {

	return func(*terraform.State) error {
		if workload.Spec.Runtime.Containers.SecurityContext == nil {
			return fmt.Errorf("expected runtime sysctl settings, but had none")
		} else if sysctls := workload.Spec.Runtime.Containers.SecurityContext.Sysctls; len(expected) > 0 && (sysctls == nil || len(sysctls) == 0) {
			return fmt.Errorf("expected sysctl overrides, but had none")
		} else if len(expected) == 0 && len(sysctls) > 0 {
			return fmt.Errorf("expected empty sysctl overrides, but had some: %v", sysctls)
		} else {
			for _, sysctl := range sysctls {
				if val, ok := expected[sysctl.Name]; !ok {
					return fmt.Errorf("unexpected sysctl override %s=%s", sysctl.Name, sysctl.Value)
				} else if sysctl.Value != val {
					return fmt.Errorf("sysctl override did not match for %s: %s != %s", sysctl.Name, sysctl.Value, val)
				}
			}
		}

		return nil
	}
}

// This test has precondition of verifying correct stsate of the runtime settings themselves
func testAccComputeWorkloadCheckRuntimeHostAliases(workload *workload_models.V1Workload,
	aliases map[string][]string,
) resource.TestCheckFunc {

	return func(*terraform.State) error {
		containerData := workload.Spec.Runtime.Containers
		if aliases == nil {
			aliases = map[string][]string{}
		}
		if len(aliases) == 0 && (containerData.HostAliases != nil && len(containerData.HostAliases) > 0) {
			return fmt.Errorf("expected empty hostaliases, but had %d", len(containerData.HostAliases))
		} else if len(aliases) > 0 && (containerData.HostAliases == nil || len(containerData.HostAliases) == 0) {
			return fmt.Errorf("expected non-empty host aliases, but they were empty")
		} else {

			for _, hostAlias := range containerData.HostAliases {
				if exp, ok := aliases[hostAlias.IP]; !ok {
					return fmt.Errorf("did not expect to have alias for %s", hostAlias.IP)
				} else {
					sort.Sort(sort.StringSlice(hostAlias.Hostnames))
					sort.Sort(sort.StringSlice(exp))
					if !reflect.DeepEqual(exp, hostAlias.Hostnames) {
						return fmt.Errorf("aliases for %s were not equal: %v != %v",
							hostAlias.IP,
							hostAlias.Hostnames,
							exp,
						)
					}
				}
			}

		}
		return nil
	}
}

// Checks the resolver_config sub-resource (dns config in API)
func testAccComputeWorkloadCheckRuntimeDNSConfig(workload *workload_models.V1Workload,
	nameservers []string,
	search []string,
	options map[string]string,
) resource.TestCheckFunc {

	return func(*terraform.State) error {
		containerData := workload.Spec.Runtime.Containers
		if containerData.DNSConfig == nil && len(nameservers) > 0 {
			return fmt.Errorf("did not expect to have dns config, but found it")
		} else {
			dnsConfig := containerData.DNSConfig

			if !reflect.DeepEqual(dnsConfig.Nameservers, nameservers) {
				return fmt.Errorf("nameservers not equal: %v != %v",
					dnsConfig.Nameservers, nameservers)
			}
			if !reflect.DeepEqual(dnsConfig.Searches, search) {
				return fmt.Errorf("search options not equal: %v != %v",
					dnsConfig.Searches, search,
				)
			}
			if dnsConfig.Options == nil && len(options) > 0 {
				return fmt.Errorf("expected to have dns options but had none")
			}

			for _, option := range dnsConfig.Options {
				if val, ok := options[option.Name]; !ok {
					return fmt.Errorf("unexpected resolver option '%s'", option.Name)
				} else if val != option.Value {
					return fmt.Errorf("unexpected value for resolver option '%s': %s != %s",
						option.Name, option.Value, val,
					)
				}
			}
		}
		return nil
	}

}

func printSlice(a []string) string {
	q := make([]string, len(a))
	for i, s := range a {
		q[i] = fmt.Sprintf("%q", s)
	}
	return fmt.Sprintf("[%s]", strings.Join(q, ", "))
}

func getInterface(network, ipv4Subnet, ipv6Subnet string, enableNAT *bool, ipFamilies []string) string {
	var config string
	generatedConfig := ""

	if enableNAT != nil {
		generatedConfig = generatedConfig + fmt.Sprintf(`
      enable_one_to_one_nat = %v
     `, *enableNAT)
	}

	if ipFamilies != nil {
		generatedConfig = generatedConfig + fmt.Sprintf(`
      ip_families = %s
     `, printSlice(ipFamilies))
	}

	if ipv4Subnet != "" {
		generatedConfig = generatedConfig + fmt.Sprintf(`
      subnet = "%s"
     `, ipv4Subnet)
	}

	if ipv6Subnet != "" {
		generatedConfig = generatedConfig + fmt.Sprintf(`
      ipv6_subnet = "%s"
     `, ipv6Subnet)
	}

	config = fmt.Sprintf(`
    network_interface {
      network = "%s"
%s
    }`, network, generatedConfig)

	return config
}

func testComputeWorkloadConfigContainerBasic(suffix string, enableNAT *bool) string {
	return fmt.Sprintf(`
resource "stackpath_compute_workload" "foo" {
  name = "My Compute Workload - %s"
  slug = "my-compute-workload-%s"
  %s

  container {
    name  = "app"
    image = "nginx:latest"
    resources {
      requests = {
        cpu    = "1"
        memory = "2Gi"
      }
    }
    port {
      name     = "http"
      port     = 80
      protocol = "TCP"
    }
    env {
      key   = "MY_ENVIRONMENT_VARIABLE"
      value = "value"
    }
  }

  target {
    name         = "us"
    min_replicas = 1
    selector {
      key      = "cityCode"
      operator = "in"
      values = [
        "AMS",
      ]
    }
  }
}`, suffix, suffix, getInterface("default", "", "", enableNAT, nil))
}

func testComputeWorkloadConfigContainerForIPFamilies(suffix, network, ipv4Subnet, ipv6Subnet string, ipFamilies []string) string {
	vpcConfig := ""

	if ipv4Subnet != "" || ipv6Subnet != "" {
		vpcConfig = fmt.Sprintf(`
resource "stackpath_compute_vpc_network" "net" {
	name = "%[1]s"
	slug = "%[1]s"
	root_subnet = "10.0.0.0/9"
	ip_families = ["IPv4", "IPv6"]
	ipv6_subnet = "fc00::/64"
}`, network)
	}

	if ipv4Subnet != "" {
		vpcConfig = vpcConfig + fmt.Sprintf(`
resource "stackpath_compute_vpc_network_subnet" "ipv4subnet" {
	name = "%[1]s"
	slug = "%[1]s"
	network_id = stackpath_compute_vpc_network.net.slug
	prefix = "11.0.0.0/9"
}`, ipv4Subnet)
	}

	if ipv6Subnet != "" {
		vpcConfig = vpcConfig + fmt.Sprintf(`
resource "stackpath_compute_vpc_network_subnet" "ipv6subnet" {
	name = "%[1]s"
	slug = "%[1]s"
	network_id = stackpath_compute_vpc_network.net.slug
	prefix = "fc01::/116"
`, ipv6Subnet)
		// add depends_on tag to mimic sequence of ipv6subnet resource
		// creation after creation of ipvsubnet in case both ipv4subnet and ipv4subnet
		// subnets are being created
		// This is required due to bug on createnetworksubnet api where it gives ""
		// when we try to create multiple subnets in same vpc at same time.
		// This is happening as it tries to update network policy based on subnet prefix
		// of newly created subnet and somehow it end up with version conflicts on etcd for
		// network policy
		if ipv4Subnet != "" {
			vpcConfig = vpcConfig + fmt.Sprintf(`
	depends_on = [stackpath_compute_vpc_network_subnet.ipv4subnet]
}`)
		} else {
			vpcConfig = vpcConfig + fmt.Sprintf(`
}`)
		}
	}

	config := fmt.Sprintf(`
resource "stackpath_compute_workload" "foo" {
  name = "My Compute Workload - %s"
  slug = "my-compute-workload-%s"
  %s

  container {
    name  = "app"
    image = "nginx:latest"
    resources {
      requests = {
        cpu    = "1"
        memory = "2Gi"
      }
    }
    port {
      name     = "http"
      port     = 80
      protocol = "TCP"
    }
    env {
      key   = "MY_ENVIRONMENT_VARIABLE"
      value = "value"
    }
  }

  target {
    name         = "us"
    min_replicas = 1
    selector {
      key      = "cityCode"
      operator = "in"
      values = [
        "AMS",
      ]
    }
  }
`, suffix, suffix, getInterface(network, ipv4Subnet, ipv6Subnet, nil, ipFamilies))

	if vpcConfig != "" {
		config = vpcConfig + config
		config = config + fmt.Sprintf(`
depends_on = [stackpath_compute_vpc_network_subnet.ipv6subnet]
}`)
	} else {
		config = config + fmt.Sprintf(`
}`)
	}

	return config
}

func testComputeWorkloadConfigContainerAddPorts(suffix string, enableNAT *bool) string {
	return fmt.Sprintf(`
resource "stackpath_compute_workload" "foo" {
  name = "My Compute Workload - %s"
  slug = "my-compute-workload-%s"
  %s

  container {
    name  = "app"
    image = "nginx:latest"
    resources {
      requests = {
        cpu    = "1"
        memory = "2Gi"
      }
    }
    port {
      name                           = "http"
      port                           = 80
      protocol                       = "TCP"
      enable_implicit_network_policy = false
    }
    port {
      name                           = "https"
      port                           = 443
      protocol                       = "TCP"
      enable_implicit_network_policy = true
    }
    env {
      key   = "MY_ENVIRONMENT_VARIABLE"
      value = "some value"
    }
  }

  target {
    name         = "us"
    min_replicas = 2
    selector {
      key      = "cityCode"
      operator = "in"
      values = [
        "AMS",
      ]
    }
  }
}`, suffix, suffix, getInterface("default", "", "", enableNAT, nil))
}

func testComputeWorkloadConfigContainerRemoveEnvVar(suffix string, enableNAT *bool) string {
	return fmt.Sprintf(`
resource "stackpath_compute_workload" "foo" {
  name = "My Compute Workload - %s"
  slug = "my-compute-workload-%s"
  %s

  container {
    name  = "app"
    image = "nginx:latest"
    resources {
      requests = {
        cpu    = "1"
        memory = "2Gi"
      }
    }
    port {
      name     = "http"
      port     = 80
      protocol = "TCP"
    }
  }

  target {
    name         = "us"
    min_replicas = 2
    selector {
      key      = "cityCode"
      operator = "in"
      values = [
        "AMS",
      ]
    }
  }
}`, suffix, suffix, getInterface("default", "", "", enableNAT, nil))
}

func testComputeWorkloadConfigContainerAddProbes(suffix string, enableNAT *bool) string {
	return fmt.Sprintf(`
resource "stackpath_compute_workload" "foo" {
  name = "My Compute Workload - %s"
  slug = "my-compute-workload-%s"
  %s

  container {
    name  = "app"
    image = "nginx:latest"
    resources {
      requests = {
        cpu    = "1"
        memory = "2Gi"
      }
    }
    port {
      name     = "http"
      port     = 80
      protocol = "TCP"
    }
    liveness_probe {
      period_seconds = 60
      success_threshold = 1
      failure_threshold = 4
      initial_delay_seconds = 60
      http_get {
        port = 80
        path = "/"
        scheme = "HTTP"
        http_headers = {
          "content-type" = "application/json"
        }
      }
    }
    readiness_probe {
      period_seconds = 60
      success_threshold = 1
      failure_threshold = 4
      initial_delay_seconds = 60
      timeout_seconds = 75
      tcp_socket {
        port = 80
      }
    }
  }

  target {
    name         = "us"
    min_replicas = 2
    selector {
      key      = "cityCode"
      operator = "in"
      values = [
        "AMS",
      ]
    }
  }
}`, suffix, suffix, getInterface("default", "", "", enableNAT, nil))
}

func testComputeWorkloadConfigContainerImagePullCredentials(suffix string, enableNAT *bool) string {
	return fmt.Sprintf(`
resource "stackpath_compute_workload" "foo" {
  name = "My Compute Workload - %s"
  slug = "my-compute-workload-%s"
  %s

  image_pull_credentials {
    docker_registry {
      server   = "docker.io"
      username = "my-registry-user"
      password = "secret registry password"
      email    = "developers@stackpath.com"
    }
  }

  container {
    name  = "app"
    image = "nginx:latest"
    resources {
      requests = {
        cpu    = "1"
        memory = "2Gi"
      }
    }
  }

  target {
    name         = "us"
    min_replicas = 1
    selector {
      key      = "cityCode"
      operator = "in"
      values = [
        "AMS",
      ]
    }
  }
}`, suffix, suffix, getInterface("default", "", "", enableNAT, nil))
}

func testComputeWorkloadConfigContainerAddContainer(suffix string, enableNAT *bool) string {
	return fmt.Sprintf(`
resource "stackpath_compute_workload" "foo" {
  name = "My Compute Workload - %s"
  slug = "my-compute-workload-%s"
  %s

  container {
    name  = "app"
    image = "nginx:latest"
    resources {
      requests = {
        cpu    = "1"
        memory = "2Gi"
      }
    }
  }

  container {
    name  = "app-2"
    image = "nginx:v1.15.0"
    resources {
      requests = {
        cpu    = "1"
        memory = "2Gi"
      }
    }
  }

  target {
    name         = "us"
    min_replicas = 1
    selector {
      key      = "cityCode"
      operator = "in"
      values = [
        "AMS",
      ]
    }
  }
}`, suffix, suffix, getInterface("default", "", "", enableNAT, nil))
}

func testComputeWorkloadConfigVirtualMachineForIPFamilies(suffix, ipv4Subnet, ipv6Subnet string, ipFamilies []string) string {
	return fmt.Sprintf(`
resource "stackpath_compute_workload" "bar" {
  name = "My Terraform Compute VM Workload - %s"
  slug = "terraform-vm-workload-%s"
  %s

  virtual_machine {
    name  = "app"
    image = "stackpath-edge/centos-7:v201905012051"
    port {
      name     = "http"
      port     = 80
      protocol = "TCP"
    }
    resources {
      requests = {
        cpu    = "1"
        memory = "2Gi"
      }
    }
    user_data = <<EOT
package_update: true
packages:
- nginx
EOT
  }

  target {
    name         = "us"
    min_replicas = 1
    selector {
      key      = "cityCode"
      operator = "in"
      values = [
        "AMS",
      ]
    }
  }
}`, suffix, suffix, getInterface("default", ipv4Subnet, ipv6Subnet, nil, ipFamilies))
}

func testComputeWorkloadConfigVirtualMachineBasic(suffix string, enableNAT *bool) string {
	return fmt.Sprintf(`
resource "stackpath_compute_workload" "bar" {
  name = "My Terraform Compute VM Workload - %s"
  slug = "terraform-vm-workload-%s"
  %s

  virtual_machine {
    name  = "app"
    image = "stackpath-edge/centos-7:v201905012051"
    port {
      name     = "http"
      port     = 80
      protocol = "TCP"
    }
    resources {
      requests = {
        cpu    = "1"
        memory = "2Gi"
      }
    }
    user_data = <<EOT
package_update: true
packages:
- nginx
EOT
  }

  target {
    name         = "us"
    min_replicas = 1
    selector {
      key      = "cityCode"
      operator = "in"
      values = [
        "AMS",
      ]
    }
  }
}`, suffix, suffix, getInterface("default", "", "", enableNAT, nil))
}

func testComputeWorkloadConfigContainerAddVolumeMounts(suffix string, enableNAT *bool) string {
	return fmt.Sprintf(`
resource "stackpath_compute_workload" "foo-volume" {
  name = "My Compute Workload Volume - %s"
  slug = "my-compute-workload-volume-%s"
  %s

  container {
    name  = "app"
    image = "nginx:latest"
    resources {
      requests = {
        cpu    = "1"
        memory = "2Gi"
      }
    }
    volume_mount {
      slug = "volume"
      mount_path = "/var/log"
    }
  }

  target {
    name         = "us"
    min_replicas = 1
    selector {
      key      = "cityCode"
      operator = "in"
      values = [
        "AMS",
      ]
    }
  }

  volume_claim {
    name = "volume"
    slug = "volume"
    resources {
      requests = {
        "storage" = "10Gi"
      }
    }
  }
}`, suffix, suffix, getInterface("default", "", "", enableNAT, nil))
}

func testComputeWorkloadConfigAutoScalingConfiguration(suffix string, enableNAT *bool) string {
	return fmt.Sprintf(`
resource "stackpath_compute_workload" "foo" {
  name = "My Compute Workload - %s"
  slug = "my-compute-workload-%s"
  %s

  image_pull_credentials {
    docker_registry {
      server   = "docker.io"
      username = "my-registry-user"
      password = "secret registry password"
      email    = "developers@stackpath.com"
    }
  }

  container {
    name  = "app"
    image = "nginx:latest"
    resources {
      requests = {
        cpu    = "1"
        memory = "2Gi"
      }
    }
    port {
      name     = "http"
      port     = 80
      protocol = "TCP"
    }

  }

  target {
    name         = "us"
    min_replicas = 2
    max_replicas = 4
    scale_settings {
      metrics {
        metric = "cpu"
        average_utilization = 50
      }
    }
    selector {
      key      = "cityCode"
      operator = "in"
      values = [
        "AMS",
      ]
    }
  }
}`, suffix, suffix, getInterface("default", "", "", enableNAT, nil))
}

// Security context without requesting capability changes
func testComputeWorkloadConfigContainerSecurityContextClearCapabilities(suffix string, enableNAT *bool) string {
	return fmt.Sprintf(`
resource "stackpath_compute_workload" "foo" {
  name = "My Compute Workload - %s"
  slug = "my-compute-workload-%s"
  %s

  container {
    name  = "app"
    image = "nginx:latest"
    resources {
      requests = {
        cpu    = "1"
        memory = "2Gi"
      }
    }
    port {
      name     = "http"
      port     = 80
      protocol = "TCP"
    }
    env {
      key   = "MY_ENVIRONMENT_VARIABLE"
      value = "value"
    }
	security_context {

		allow_privilege_escalation = true
		run_as_non_root = true
		run_as_user = "101"

	}
  }

  target {
    name         = "us"
    min_replicas = 1
    selector {
      key      = "cityCode"
      operator = "in"
      values = [
        "AMS",
      ]
    }
  }
}`, suffix, suffix, getInterface("default", "", "", enableNAT, nil))
}

func testComputeWorkloadConfigContainerSecurityContextCapabilities(suffix string, enableNAT *bool) string {
	return fmt.Sprintf(`
resource "stackpath_compute_workload" "foo" {
  name = "My Compute Workload - %s"
  slug = "my-compute-workload-%s"
  %s

  container {
    name  = "app"
    image = "nginx:latest"
    resources {
      requests = {
        cpu    = "1"
        memory = "2Gi"
      }
    }
    port {
      name     = "http"
      port     = 80
      protocol = "TCP"
    }
    env {
      key   = "MY_ENVIRONMENT_VARIABLE"
      value = "value"
    }
	security_context {

		allow_privilege_escalation = true
		run_as_non_root = true
		run_as_user = "101"

		capabilities {
			add = [
				"NET_ADMIN",
			]

			drop = [
				"NET_BROADCAST",
			]
		}
	}
  }

  target {
    name         = "us"
    min_replicas = 1
    selector {
      key      = "cityCode"
      operator = "in"
      values = [
        "AMS",
      ]
    }
  }
}`, suffix, suffix, getInterface("default", "", "", enableNAT, nil))
}

func testComputeWorkloadConfigContainerRuntimeSettings(suffix string, enableNAT *bool) string {
	return fmt.Sprintf(`
resource "stackpath_compute_workload" "foo" {
  name = "My Compute Workload - %s"
  slug = "my-compute-workload-%s"
  %s

  container {
    name  = "app"
    image = "nginx:latest"
    resources {
      requests = {
        cpu    = "1"
        memory = "2Gi"
      }
    }
    port {
      name     = "http"
      port     = 80
      protocol = "TCP"
    }
    env {
      key   = "MY_ENVIRONMENT_VARIABLE"
      value = "value"
    }
  }

  target {
    name         = "us"
    min_replicas = 1
    selector {
      key      = "cityCode"
      operator = "in"
      values = [
        "AMS",
      ]
    }
  }

  container_runtime_environment {
	termination_grace_period_seconds = 60
	share_process_namespace = true
	security_context {
		run_as_user = "999"
		run_as_group = "991"
		run_as_non_root = true

		supplemental_groups = [ 
			"42", 
		]

		sysctl = {
			"net.core.rmem_max" = "10065408"
			"net.core.rmem_default" = "1006540"
		}
	}



	dns {

		host_aliases {
			address = "192.168.3.4"
			hostnames = [ "domain.com" ]
		}

		resolver_config {
			nameservers = [ "8.8.8.8" ]
			search = [ "domain.com" ]
			options = {
				timeout = "10"
			}
		}
	}


  }
}`, suffix, suffix, getInterface("default", "", "", enableNAT, nil))
}

func testComputeWorkloadConfigContainerRuntimeClearSettings(suffix string, enableNAT *bool) string {
	return fmt.Sprintf(`
resource "stackpath_compute_workload" "foo" {
  name = "My Compute Workload - %s"
  slug = "my-compute-workload-%s"
  %s

  container {
    name  = "app"
    image = "nginx:latest"
    resources {
      requests = {
        cpu    = "1"
        memory = "2Gi"
      }
    }
    port {
      name     = "http"
      port     = 80
      protocol = "TCP"
    }
    env {
      key   = "MY_ENVIRONMENT_VARIABLE"
      value = "value"
    }
  }

  target {
    name         = "us"
    min_replicas = 1
    selector {
      key      = "cityCode"
      operator = "in"
      values = [
        "AMS",
      ]
    }
  }

  container_runtime_environment {
	termination_grace_period_seconds = 60
	share_process_namespace = true
	security_context {
		run_as_user = "999"
		run_as_group = "991"
		run_as_non_root = true

		supplemental_groups = [ 
			"43", 
		]

		sysctl = {
			"net.core.rmem_max" = "10065408"
			"net.core.rmem_default" = "1006540"
		}
	}



	dns {

		resolver_config {
			nameservers = [ "8.8.8.8" ]
			search = [ "domain2.com" ]
		}
	}


  }
}`, suffix, suffix, getInterface("default", "", "", enableNAT, nil))
}
