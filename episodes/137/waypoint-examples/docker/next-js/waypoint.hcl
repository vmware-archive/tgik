project = "example-nextjs"

app "example-nextjs" {
  labels = {
    "service" = "example-nextjs",
    "env" = "dev"
  }

  build {
    use "pack" {}
  }

  deploy { 
    use "docker" {}
  }
}
