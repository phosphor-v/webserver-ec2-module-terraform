terraform {
  source = "git::https://github.com/phosphor-v/webserver-ec2-module-terraform.git//infrastructure"
}

inputs = {
  region             = "eu-north-1"
  vpc_cidr           = "192.168.0.0/16"
  network_http       = ["192.168.1.0/24"]
  db_subnet_cidrs      = ["192.168.2.0/24"]
  instance_type      = "t3.micro"
  db_instance_class  = "db.t3.micro"
  db_name            = "mydatabase"
  db_user            = "admin"
  db_password        = "password"
}

remote_state {
  backend = "s3"
  config = {
    bucket         = "your-s3-bucket"
    key            = "terraform.tfstate"
    region         = "us-east-1"
  }
}
