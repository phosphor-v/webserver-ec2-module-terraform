terraform {
  source = "git::https://github.com/phosphor-v/webserver-ec2-module-terraform.git//infrastructure"
}

inputs = {
  region             = "eu-north-1"
  vpc_cidr           = "192.168.0.0/16"
  network_http       = ["192.168.1.0/24"]
  db_subnet_cidrs      = ["192.168.2.0/24"]
  instance_type      = "t3.micro"
}

