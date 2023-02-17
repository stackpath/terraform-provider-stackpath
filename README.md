<a href="https://terraform.io">
    <img src=".github/terraform_logo.svg" alt="Terraform logo" title="Terraform" align="right" height="50" />
</a>

# Terraform Provider for StackPath

The StackPath provider is a plugin for Terraform to interact with resources on the StackPath edge platform. It is publicly available on the [Terraform registry](https://registry.terraform.io/providers/stackpath/stackpath/latest). Please see the [official documentation](https://registry.terraform.io/providers/stackpath/stackpath/latest/docs) to get started.

## Provider development

### Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.12+ (to run acceptance tests)
- [Go](https://golang.org/doc/install) 1.16+ (to build the provider plugin)

### Building the provider

Run the following command to build the provider:

```sh
$ make build
```

### Generating the StackPath API client code

We need to install swagger cli tool to generate StackPath API client code
(See: https://goswagger.io/install.html)
Note: Current client code is generated using go-swagger version v0.27.0, so it is
recommended that we use same version to avoid client code differences being introduced
due to latest version of swagger unless we decide to bump up client code versioning.

Once swagger cli is installed, Run the following command to generate client code:

```sh
$ make generate
```

### Installing the built provider for Terraform < 0.13

Run the following command to install the built provider:

```sh
$ make install
```

Once the plugin has been installed, run `terraform init` to have terraform discover the StackPath plugin.

### Installing the built provider for Terraform >= 0.13

Make sure that your terraform config is setup to mirror local providers:
Something like this should be in your `~/.terraformrc`

```
provider_installation {
  filesystem_mirror {
    path    = "/Users/USERNAME/terraform-providers/"
    include = ["local/providers/*"]
  }
  direct {
    exclude = ["local/providers/*"]
  }
}
```

Run the following command to install the built provider:

```sh
$ make install-13
```

Once the plugin has been installed, run `terraform init` to have terraform discover the StackPath plugin.

### Testing the provider

Run the following command to run the provider's unit tests:

```sh
$ make test
```

In order to run the full suite of acceptance tests, run `make testacc`. You must declare a valid StackPath stack ID or slug, API client ID, and API client secret in the `STACKPATH_STACK_ID`, `STACKPATH_CLIENT_ID`, and `STACKPATH_CLIENT_SECRET` environment variables:

_Note:_ Acceptance tests create real resources, and often cost money to run.

```sh
$ STACKPATH_STACK_ID=my-stack-id STACKPATH_CLIENT_ID=my-client-id STACKPATH_CLIENT_SECRET=my-client-secret make testacc
```
