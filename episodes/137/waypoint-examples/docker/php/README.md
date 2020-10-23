# Waypoint PHP Example

|Title|Description|
|---|---|
|Pack|Cloud Native Buildpack|
|Cloud|Local|
|Language|PHP|
|Docs|[Docker](https://www.waypointproject.io/plugins/docker)|
|Tutorial|[HashiCorp Learn](https://learn.hashicorp.com/tutorials/waypoint/get-started-docker)|

A barebones Laravel 8 app, which can easily be deployed by Waypoint.

This repository provides an .env file to make it easy to deploy the example,
however these environment variables should be managed using `waypoint config`.

Laravel `TrustProxies` has been set to `*` as an example to allow Laravel to
generate the correct SSL asset URL.
