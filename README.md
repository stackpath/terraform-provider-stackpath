# The StackPath Terraform Provider Has Moved!

<img src="https://www.stackpath.com/content/images/logo-and-branding/stackpath-monogram-reversed-screen.svg" width="145" alt="StackPath"> 

:heavy_plus_sign:

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px" alt="HashiCorp Terraform">
<br><br>

The StackPath Terraform provider is now an [official Terraform provider](https://www.terraform.io/docs/providers/stackpath/index.html), and its code now lives at [https://github.com/terraform-providers/terraform-provider-stackpath](https://github.com/terraform-providers/terraform-provider-stackpath). Our original releases will stay here for those who need them, but please open issues and pull requests against the official repository. 

Thanks, everyone!

 * SP// 

---

Terraform Provider For StackPath
==================

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

Maintainers
------------------

This terraform provider plugin is maintained by the Engineering team at [StackPath](https://www.stackpath.com/).

Requirements
------------------

- [Terraform](https://www.terraform.io/downloads.html) 0.10.x+
- [Go](https://golang.org/doc/install) 1.11+ (to build the provider plugin)

Building The Provider
------------------

```sh
$ mkdir -p $GOPATH/src/github.com/terraform-providers; cd $GOPATH/src/github.com/terraform-providers
$ git clone git@github.com:terraform-providers/terraform-provider-template
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-template
$ make build
```

Installing the provider
------------------

After downloading the latest release from GitHub, move the binary into the third party plugin directory on your workstation.
Third-party plugins (both providers and provisioners) can be manually installed into the user plugins directory, located at `%APPDATA%\terraform.d\plugins\<OS>_<ARCH>` on Windows and `~/.terraform.d/plugins/<OS>_<ARCH>` on other systems.

Using macOS as an example:

```shell
// TODO add curl command for latest release
$ mv ./terraform-provider-stackpath_$VERSION ~/.terraform.d/plugins/darwin_amd64/
```

Once the plugin has been installed, run `terraform init` to have terraform discover the StackPath plugin.

Using the provider
------------------

Before you can use the StackPath provider, you will need to configure the provider with the stack ID and API credentials that should be used for managing resources. See StackPath's [getting started guide](https://stackpath.dev/docs/getting-started) for more information on finding your stack ID and API credentials.

```terraform
provider "stackpath" {
  # only allow version 0.1 of the StackPath provider to be used
  version = "~> 0.1"
  # ID of the stack that resources should be created in
  stack_id = "{{ stack-id }}"
  # The API credentials that should be used for authentication
  client_id = "{{ client-id }}"
  client_secret = "{{ client-secret }}"
}
```

Developing the Provider
------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.11+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-template
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```
