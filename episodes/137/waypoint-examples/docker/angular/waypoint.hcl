project = "angular-example"

app "angular" {
    build {
        use "pack" {}
    }

    deploy {
        use "docker" {
        }
    }
}
