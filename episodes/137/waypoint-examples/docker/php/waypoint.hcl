project = "example-php"

app "example-php" {
  labels = {
    "service" = "example-php",
    "env" = "dev"
  }

  build {
    use "pack" {}
  }

  deploy {
    use "docker" {}
  }
}
