## 1.3.1 (July 21, 2020)

BUG FIXES:

- Prevent various errors updating network policies. ([#14](https://github.com/terraform-providers/terraform-provider-stackpath/pull/14))
- Prevent workloads with anycast subnets from re-creating when `terraform apply` is run with no workload changes. ([#15](https://github.com/terraform-providers/terraform-provider-stackpath/pull/15))

ENHANCEMENTS:

- Tweak example files for correctness and readability. ([#13](https://github.com/terraform-providers/terraform-provider-stackpath/pull/13))

## 1.3.0 (April 01, 2020)

BUG FIXES:

- Numerous documentation and example fixes ([#10](https://github.com/terraform-providers/terraform-provider-stackpath/pull/10) and [#12](https://github.com/terraform-providers/terraform-provider-stackpath/pull/12), thanks [@prhomhyse](https://github.com/prhomhyse)!):
  - Virtual machine `user_data` should not be base64 encoded.
  - Added a full compute workload example using virtual machines.
  - Clarified the `enable_implicit_network_policy` port field's effect.
  - Clarified use of selectors in network policy and compute workload resources.
  - Clarified use of volume claims in compute workload resources. 
  - Fixed incorrect naming and documentation in the full network policy example.

ENHANCEMENTS:

- Errors from the StackPath API are presented in a much more readable format along with a request ID that can be reported to StackPath support. ([#11](https://github.com/terraform-providers/terraform-provider-stackpath/pull/11))

## 1.2.1 (March 11, 2020)

BUG FIXES:

- Correct the documentation for the virtual machine `image` field ([#7](https://github.com/terraform-providers/terraform-provider-stackpath/pull/7))

ENHANCEMENTS:

- Migrate internal code from Terraform Core to the Terraform provider SDK ([#6](https://github.com/terraform-providers/terraform-provider-stackpath/pull/6))
- Various syntactic touch-ups of example Terraform files and snippets ([#7](https://github.com/terraform-providers/terraform-provider-stackpath/pull/7))

## 1.2.0 (February 21, 2020)

FEATURES:

- **New Resource:** `stackpath_object_storage_bucket` to control storage buckets. See the resource's [documentation](https://www.terraform.io/docs/providers/stackpath/r/object_storage_bucket.html) for more information ([#3](https://github.com/terraform-providers/terraform-provider-stackpath/pull/3), thanks [@davebond](https://github.com/davebond)!)

ENHANCEMENTS:

- Various internal code cleanups ([#5](https://github.com/terraform-providers/terraform-provider-stackpath/pull/5))
- The User-Agent sent with StackPath API calls is now `HashiCorp Terraform/<terraform_version> (+https://www.terraform.io) terraform-provider-stackpath/<provider_version> (+https://www.terraform.io/docs/providers/stackpath)` ([#5](https://github.com/terraform-providers/terraform-provider-stackpath/pull/5))

## 1.1.0 (February 11, 2020)

ENHANCEMENTS:

- `stackpath_compute_workload`: Add an `enable_implicit_network_policy` property to port definitions. When `true`, the port is accessible from the public Internet ([#4](https://github.com/terraform-providers/terraform-provider-stackpath/pull/4))

## 1.0.0 (December 06, 2019)

- Initial Release
