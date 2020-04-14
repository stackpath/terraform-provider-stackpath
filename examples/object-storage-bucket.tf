# Create a StackPath object storage bucket
resource "stackpath_object_storage_bucket" "bucket" {
  # A human readable label for the bucket. Labels can have the characters "a-z",
  # "0-9", an "-" and must start and end with a letter or number
  label = "my-object-storage-bucket"
  # The bucket's geogrtaphical region, "us-east-2", "us-west-1", or
  # "eu-central-1"
  region = "us-east-2"
  # Either "PUBLIC" or "PRIVATE". Defaults to "PRIVATE"
  visibility = "PRIVATE"
}
