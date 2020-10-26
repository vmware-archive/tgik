project = "example-nodejs"

app "example-nodejs" {
  labels = {
    "service" = "example-nodejs",
    "env"     = "dev"
  }

  build {
    use "pack" {}

    registry {
      use "docker" {
        image = "<my-container-registry>.azurecr.io/example-nodejs"
        tag   = "latest"
      }
    }
  }

  deploy {
    use "azure-container-instance" {
      resource_group = "DefaultResourceGroup-EUS"
      location       = "eastus"
      ports          = [8080]

      capacity {
        memory    = "512"
        cpu_count = 1
      }

    }
  }

}
