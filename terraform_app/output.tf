output "address" {
  value       = var.host
  description = "connect with  local address or home address"
}

output "engine_internal" {
  value       = var.engine_port
  description = "Internal port route ingress traffic"
}

output "engine_external" {
  value       = var.engine_port
  description = "External port route egress traffic"
}

output "engine_image" {
  value       = var.engine_id
  description = "Fictional Engine image "
}

output "redis_igress_port" {
  value       = var.redis_port
  description = "Redis port allocate"
}

output "redis_egress_port" {
    value = var.redis_port
}

output "redis_image" {
    value = var.redis_repo_image
}

