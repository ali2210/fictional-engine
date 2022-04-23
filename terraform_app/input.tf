variable "password" {
  type        = string
  default     = ""
  sensitive   = true
  description = "Redis password"
}

variable "host" {
  type        = string
  default     = "127.0.0.1"
  description = "Host Address"
}


variable "engine" {
  type    = string
  default = "Fictional-engine started ...."
}

variable "engine_port" {
  type    = string
  default = 3000
}

variable "engine_id" {
  type    = string
  default = "3c3da61c4be0fb9e93fc84fb702b6a1e01d8c43ffae4d2ea88192803c2b44191"
}

variable "redis_port" {
  type    = string
  default = 6379
}

variable "redis_repo_image"{
    type = string
    default = "redis:latest"
}