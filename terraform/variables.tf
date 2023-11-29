variable "tags" {
  description = "A map of tags to add to all resources"
  type        = map(string)
  default = {
    Name        = "k8s-vpc"
    Component   = "k8s"
    Environment = "study"
  }
}
