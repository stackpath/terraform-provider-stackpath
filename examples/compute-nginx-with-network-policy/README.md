# EdgeCompute with Nginx

This example provides a demo of how to launch an nginx container using StackPath EdgeCompute and levering network policies to open up access to the web server.

## Getting Started

To use this example you must have a pair or API credentials for your StackPath account and the ID of your StackPath stack.
You can follow our [getting started guide](https://stackpath.dev/docs/getting-started) to create API credentials.

> You must also have terraform installed and configured to use this extension. For more information visit the [terraform documentation](https://terraform.io/)

Once you have terraform installed with the custom StackPath plugin, you can execute the plan with the following command.

```shell
$ terraform plan
```

This command will provide a plan of all the resources it needs to create in the StackPath platform.

To create the resources execute the following command.

```shell
$ terraform apply
```

Once terraform successfully creates the resources in the API, [log into the portal](https://control.stackpath.com/) and visit the workloads tab.
You should see a newly provisioned nginx container running in Seattle and New York!
