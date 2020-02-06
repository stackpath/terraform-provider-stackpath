---
layout: "stackpath"
page_title: "StackPath: stackpath_object_storage_bucket"
sidebar_current: "docs-stackpath-resource-object-storage-bucket"
description: |-
  An S3 compatible object storage bucket deployed to StackPath's edge network.
---

# stackpath\_object\_storage\_bucket

An S3 compatible object storage bucket deployed to StackPath's edge network.

## Example Usage

```hcl
resource "stackpath_object_storage_bucket" "my-bucket" {
  label = "my-bucket"
  region = "us-west-1"
  visibility = "PRIVATE"
}
```

## Argument Reference

* `label` - (Required) A human readable label for the bucket. Bucket label only supports (a-z, 0-9, -) and must start/end with a letter or number.
* `region` - (Required) Bucket region (us-east-2 us-west-1 eu-central-1)
* `visibility` - (Optional) PRIVATE or PUBLIC, defaults to PRIVATE

## Attributes Reference

* `endpoint_url` - S3 compatible region endpoint for the bucket, e.g. https://s3.us-east-2.stackpathstorage.com

## Import

StackPath object storage buckets can be imported by their UUID v4 formatted id. e.g.

```
$ terraform import stackpath_object_storage_bucket.bucket bdb77768-2938-4ad8-a736-be5290add801
```
