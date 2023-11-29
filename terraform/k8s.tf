module "k8s" {
  source            = "../../aws-modules-terraform/k8s" // replace with the path to your module
  count             = 2
  key_name          = "k8s.aureliomalheiros.tech"
  availability_zone = "us-east-1a"
  instance_type     = "t3a.small"
  tags_instance     = { "Name" = "k8s-instance-${count.index + 1}.aureliomalheiros.tech" }
}
