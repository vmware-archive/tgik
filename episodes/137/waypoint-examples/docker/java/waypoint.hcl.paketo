project = "example-java"

app "example-java" {
    build {
        use "pack" {
            builder="paketobuildpacks/builder:base"
        }
    }
    deploy {
        use "docker" {
            service_port=8080
        }
    }
}