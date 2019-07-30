# Grab the latest ubuntu AMI from the API. We will use this
# when creating a new EC2 Instance.
data "aws_ami" "ubuntu" {
  most_recent = true

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["099720109477"]
}

# Create a new VPC using the 10.0.0.0/16 CIDR block
resource "aws_vpc" "main" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true

  tags = {
    Name = "main"
  }
}

# Create a new subnet for the created VPC
resource "aws_subnet" "main" {
  vpc_id                  = "${aws_vpc.main.id}"
  cidr_block              = "10.0.1.0/24"
  map_public_ip_on_launch = true

  tags = {
    Name = "main"
  }
}

# Create a new internet gateway for the VPC
resource "aws_internet_gateway" "main" {
  vpc_id = "${aws_vpc.main.id}"
  tags = {
    Name = "main"
  }
}

# Manage the default route table of the VPC and
# add a route for 0.0.0.0/0 that sends traffic
# to the managed internet gateway.
resource "aws_default_route_table" "main" {
  default_route_table_id = "${aws_vpc.main.default_route_table_id}"
  tags = {
    "Name" = "main"
  }
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = "${aws_internet_gateway.main.id}"
  }
}

# Create a new security group that allows inbound http requests
resource "aws_security_group" "allow_inbound_http" {
  name        = "allow-inbound-http"
  description = "Allow inbound HTTP traffic"
  vpc_id      = "${aws_vpc.main.id}"

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# Create a new security group that allows outbound traffic
resource "aws_security_group" "allow_outbound_traffic" {
  name        = "allow-outbound-traffic"
  description = "Allow all outbound traffic"
  vpc_id      = "${aws_vpc.main.id}"

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# Define an output value of the IP of the EC2 instance
output "aws-nginx-ip" {
  value = "${aws_instance.web_server_01.public_ip}"
}

# Create a new instance of the latest Ubuntu 14.04 on an
# t2.micro node with an AWS Tag naming it "web-server-01"
resource "aws_instance" "web_server_01" {
  ami           = "${data.aws_ami.ubuntu.id}"
  instance_type = "t2.micro"
  subnet_id     = "${aws_subnet.main.id}"
  user_data     = <<EOT
#cloud-config
# update apt on boot
package_update: true
# install nginx
packages:
- nginx
write_files:
- content: |
    <!DOCTYPE html>
    <html>
    <head>
      <title>StackPath - Amazon Web Services Instance</title>
      <meta http-equiv="Content-Type" content="text/html;charset=UTF-8">
      <style>
        html, body {
          background: #000;
          height: 100%;
          width: 100%;
          padding: 0;
          margin: 0;
          display: flex;
          justify-content: center;
          align-items: center;
          flex-flow: column;
        }
        img { width: 250px; }
        svg { padding: 0 40px; }
        p {
          color: #fff;
          font-family: 'Courier New', Courier, monospace;
          text-align: center;
          padding: 10px 30px;
        }
      </style>
    </head>
    <body>
      <img src="https://www.stackpath.com/content/images/logo-and-branding/stackpath-logo-standard-screen.svg">
      <p>This request was proxied from <strong>Amazon Web Services</strong></p>
    </body>
    </html>

  path: /usr/share/app/index.html
  permissions: '0644'
runcmd:
- cp /usr/share/app/index.html /usr/share/nginx/html/index.html
EOT

  tags = {
    Name = "web-server-01"
  }

  vpc_security_group_ids = [
    aws_security_group.allow_inbound_http.id,
    aws_security_group.allow_outbound_traffic.id,
  ]
}
