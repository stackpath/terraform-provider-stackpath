# Multi-Cloud LoadBalancer with AWS, GCP, and StackPath

This example provides a fully functional Terraform example that will create two nginx servers, one in AWS and one in GCP.
Then it will create a global StackPath EdgeCompute workload that uses a Traefik container to balance requests between the two nginx servers.
For a more detailed guide on how this example was created, check out the [developer documentation](https://stackpath.dev/docs/using-terraform-to-create-a-multi-cloud-load-balancer).
