# Java Getting Started

|Title|Description|
|---|---|
|Pack|Multiple (read notes below)|
|Cloud|Multiple|
|Language|Java|
|Docs|[Docker](https://www.waypointproject.io/plugins/docker)|
|Tutorial|[HashiCorp Learn](https://learn.hashicorp.com/tutorials/waypoint/get-started-docker)|

This is an example Java Spring application that can be deployed with Waypoint.

Waypoint defaults to using Heroku buildpacks if you do not specify a [builder variable](https://waypointproject.io/plugins/pack#builder) in the `waypoint.hcl` configuration section for `pack` within the `build` section. This example uses Heroku buildpacks by default.

This example also supports using [Cloud Foundry Paketo Buildpacks](https://paketo.io/docs/) and [Google Cloud Platform (GCP) Buildpacks](https://github.com/GoogleCloudPlatform/buildpacks).

# Deploying the example application.

1. Install a Waypoint Server up and ensure `waypoint context verify` is successful.
1. `waypoint init`
1. `waypoint up`
1. Visit a URL provided in the `waypoint` output.

## Configuring Waypoint for Cloud Foundry Paketo or GCP Buildpacks.

Copy the contents of `waypoint.hcl.paketo` or `waypoint.hcl.gcp` file into `waypoint.hcl` or run a command like `cp waypoint.hcl.paketo waypoint.hcl` from the base directory. Then you can run the `waypoint` commands as described above.

There are several `waypoint.hcl` adjustments to enable non-Heroku Buildpacks.
1. The `builder` variable of the `pack` build plugin should specify the a Paketo buildpacks builder image like `paketobuildpacks/builder:base` or  `gcr.io/buildpacks/builder:v1`. If no builder image is specified, Waypoint uses the Heroku builder image `heroku/buildpacks:18` by default.
1. The `service_port` variable of the deploy plugin explicity specifies port `8080` which is commonly used by Paketo and GCP buildpacks. The default Waypoint `service_port` is `3000`.

Here is a full `waypoint.hcl` example that works with Paketo buildpacks.

```
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
```

Here is a full `waypoint.hcl` example that works with GCP buildpacks.

```
project = "example-java"

app "example-java" {
    build {
        use "pack" {
            builder="gcr.io/buildpacks/builder:v1"
        }
    }
    deploy {
        use "docker" {
            service_port=8080
        }
    }
}
```
