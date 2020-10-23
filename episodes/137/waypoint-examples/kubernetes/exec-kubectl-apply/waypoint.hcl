project = "example-nodejs-exec"

app "example-nodejs-exec" {
  labels = {
    "service" = "example-nodejs-exec",
    "env" = "dev"
  }

  build {
    use "pack" {}
    registry {
        use "docker" {
          image = "example-nodejs-exec"
          tag = "1"
          local = true
        }
    }
  }

  deploy {
    use "exec" {
      # note "<TPL>" means replace this with the path to
      # the templated file post-rendering
      command = ["kubectl", "apply", "-f", "<TPL>"]

      # this file has go template syntax for {{.Input.DockerImageFull}} and
      # {{.Env}} environment variables
      template {
        path = "./example-nodejs-exec.yml"
      }
    }
  }
}
