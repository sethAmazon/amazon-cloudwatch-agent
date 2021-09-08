resource "aws_instance" "integration-test" {
  ami           = var.ami
  instance_type = var.ec2_instance_type
  key_name = var.key_name
  iam_instance_profile = var.iam_instance_profile
  vpc_security_group_ids = var.vpc_security_group_ids
  provisioner "remote-exec" {
    inline = [
      "cloud-init status --wait"
    ]
    connection {
      type = "ssh"
      user = "ubuntu"
      private_key = var.SSH_KEY
      host = self.public_dns
    }
  }
}