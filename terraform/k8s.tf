module "k8s" {
  source                 = "git@github.com:aureliomalheiros/aws-modules-terraform.git//k8s" // replace with the path to your module
  worker_count           = 1
  key_name               = "k8s.aureliomalheiros.tech"
  availability_zone      = "us-east-1a"
  instance_type          = "t3a.small"
  tags_instance          = "k8s-instance-workers"
  subnet_id              = "subnet-03bda73c3706511b9"
  vpc_security_group_ids = ["sg-030b02d63f0588700"]
}
