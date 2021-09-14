variable "ec2_instance_type" {
  type = string
  default = "t3a.xlarge"
}

variable "ami" {
  type = string
  default = "ami-0d681845843851358"
}

variable "key_name" {
  type = string
  default = "cwagent-integ-test-key"
}

variable "iam_instance_profile" {
  type = string
  default = "CloudWatchAgentServerRole"
}

variable "vpc_security_group_ids" {
  type = list(string)
  default = ["sg-013585129c1f92bf0"]
}

variable "region" {
  type = string
  default = "us-west-2"
}

variable "AWS_ACCESS_KEY_ID" {
  type = string
}

variable "AWS_SECRET_ACCESS_KEY" {
  type = string
}

variable "SSH_KEY" {
  type = string
  default = ""
}

output "instance_ip" {
  description = "The public ip for ssh access"
  value       = aws_instance.integration-test.public_dns
}