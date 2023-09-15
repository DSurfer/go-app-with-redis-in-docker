//declaring variables for auth and instance settings

variable "user_name" {
  type = string
  sensitive = false
}

variable "password" {
  type = string
  sensitive = true
}

variable "api_key" {
  type = string
  description = "Token for your project"
  sensitive = true
}

variable "image_flavor" {
    type = string
    default = "Ubuntu-18.04-Standard"
}

variable "compute_flavor" {
  type = string
  default = "Basic-1-1-10"
}

variable "key_pair_name" {
  type = string
}

variable "availability_zone_name" {
  type = string
  default = "MS1"
}
