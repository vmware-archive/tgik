project = "example-nodejs"

app "example-nodejs" {

  build {
    use "pack" {}
    registry {
        use "docker" {
          image = "nodejs-example"
          tag = "1"
          local = true
        }
    }
 }

  deploy { 
    use "nomad" {
      datacenter = "dc1"
    }
  }

}
