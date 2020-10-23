project = "example-nodejs"

app "example-nodejs" {
  labels = {
      "service" = "example-nodejs",
      "env" = "dev"
  }

  build {
    use "pack" {}
    registry {
        use "docker" {
          image = "joshrosso/example-nodejs"
          tag = "tgik"
        }
    }
 }

  deploy { 
    use "kubernetes" {
    }
  }

  release {
    use "kubernetes" {
    }
  }
}
