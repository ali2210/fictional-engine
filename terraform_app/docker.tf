resource "docker_container" "fictional_engine" {

  image = var.engine_id
  name  = "fictional_engine"
  ports {
    internal = var.engine_port
    external = var.engine_port
  }
}

resource "docker_container" "redis" {
  name  = "redis"
  image = var.redis_repo_image
  ports {
    internal = var.redis_port
    external = var.redis_port
  }
}