# Exec Plugin Example with kubectl apply

A Node.js app packaged with Buildpacks and deployed with Kubernetes.

Rather than using the standard Waypoint `kubernetes` plugin that creates a
Deployment and Service with a single port, this example uses the [`exec`
plugin](https://www.waypointproject.io/plugins/exec#examples) to execute
`kubectl apply -f` with a templated file. This approach may be helpful when you
want to control precisely which Kubernetes resources are created such as a
StatefulSet, an app with multiple ports, or similar levels of Kubernetes
customizations.

This approach is very flexible, yet does have some downsides such as not
actually implementing the `waypoint destroy` command, which means resources must
be cleaned up separately. 

Reference `waypoint.hcl` and `example-nodejs-exec.yml` to see how it works. Note
the `{{.Input.DockerImageFull}}` and the loop for passing through the
Environment Variables to the Kubernetes Deployment:

```
{{range $key, $value := .Env}}
          - name: {{$key}}
            value: "{{$value}}"{{end}}
```

An alternative access syntax for .Env if you know the keys is:
```
{{index .Env "WAYPOINT_CEB_INVITE_TOKEN"}}
{{index .Env "WAYPOINT_DEPLOYMENT_ID"}}
{{index .Env "WAYPOINT_SERVER_ADDR"}}
{{index .Env "WAYPOINT_SERVER_TLS"}}
{{index .Env "WAYPOINT_SERVER_TLS_SKIP_VERIFY"}}
```

Note that the Waypoint workspace name will be used to create a Kubernetes
namespace of the same name with `{{.Workspace}}` syntax to retrieve the 
name in the template.

## Deployment steps
1. `waypoint init`
1. `waypoint up`
1. Validate that the app is available at the Deployment URL
1. Validate that k8s resources were deployed: `kubectl get deployment.apps/example-nodejs-exec service/example-nodejs-exec -n YOUR_NAMESPACE`

## Cleanup
1. `waypoint destroy`
1. Run `kubectl delete deployment.apps/example-nodejs-exec service/example-nodejs-exec -n YOUR_NAMESPACE`

If you used a non-default Workspace, then you may want to also delete the
workspace.
1. Run `kubectl delete namespaces YOUR_NAMESPACE`

## Example output
```
$ waypoint up

» Building...
Creating new buildpack-based image using builder: heroku/buildpacks:18
✓ Creating pack client
✓ Building image
 │ [exporter] Adding 1/1 app layer(s)
 │ [exporter] Reusing layer 'launcher'
 │ [exporter] Reusing layer 'config'
 │ [exporter] Adding label 'io.buildpacks.lifecycle.metadata'
 │ [exporter] Adding label 'io.buildpacks.build.metadata'
 │ [exporter] Adding label 'io.buildpacks.project.metadata'
 │ [exporter] *** Images (aa6a41914ac8):
 │ [exporter]       index.docker.io/library/example-nodejs-exec:latest
 │ [exporter] Reusing cache layer 'heroku/nodejs-engine:nodejs'
 │ [exporter] Reusing cache layer 'heroku/nodejs-engine:toolbox'
✓ Injecting entrypoint binary to image

Generated new Docker image: example-nodejs-exec:latest
✓ Tagging Docker image: example-nodejs-exec:latest => example-nodejs-exec:1

» Deploying...
✓ Rendering templates...
✓ Executing command: kubectl apply -f /var/folders/3f/pgsg1xp16rx2l9hqnqs8ft1h0000gp/T/waypoint-exec344460138/example-nodejs-exec.yml
 │ deployment.apps/example-nodejs-exec configured
 │ service/example-nodejs-exec unchanged

» Releasing...

The deploy was successful! A Waypoint deployment URL is shown below. This
can be used internally to check your deployment and is not meant for external
traffic. You can manage this hostname using "waypoint hostname."

           URL: https://rationally-helped-antelope.waypoint.run
Deployment URL: https://rationally-helped-antelope--v1.waypoint.run
```

