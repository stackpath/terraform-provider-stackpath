## 1.2.0 (February 21, 2020)

FEATURES:

- **New Resource:** `stackpath_object_storage_bucket` to control storage buckets. See the resource's [documentation](https://www.terraform.io/docs/providers/stackpath/r/object_storage_bucket.html) for more information ([#3](https://github.com/terraform-providers/terraform-provider-stackpath/pull/3), thanks [@davebond](https://github.com/davebond)!)

ENHANCEMENTS:

- Various internal code cleanups ([#5](https://github.com/terraform-providers/terraform-provider-stackpath/pull/5))
- The User-Agent sent with StackPath API calls is now `HashiCorp Terraform/<terraform_version> (+https://www.terraform.io) terraform-provider-stackpath/<provider_version> (+https://www.terraform.io/docs/providers/stackpath)` ([#5](https://github.com/terraform-providers/terraform-provider-stackpath/pull/5))

## 1.1.0 (February 11, 2020)

ENHANCEMENTS:

- `stackpath_compute_workload`: Add an `enable_implicit_network_policy` property to port definitions. When `true`, the port will have access to the public Internet ([#4](https://github.com/terraform-providers/terraform-provider-stackpath/pull/4))

## 1.0.0 (December 06, 2019)

- Initial Release
