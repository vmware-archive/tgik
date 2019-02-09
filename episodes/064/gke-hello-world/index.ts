import * as k8s from "@pulumi/kubernetes";
import * as pulumi from "@pulumi/pulumi";
import * as gcp from "@pulumi/gcp";
import { clusterConfig } from "./config";
import * as fs from "fs";

const name = "helloworld";

// Create a GKE cluster
const cluster = new gcp.container.Cluster(name, {
    initialNodeCount: clusterConfig.nodeCount,
    minMasterVersion: clusterConfig.minMasterVersion,
    nodeVersion: clusterConfig.nodeVersion,
    nodeConfig: {
        machineType: clusterConfig.nodeMachineType,
        oauthScopes: [
            "https://www.googleapis.com/auth/compute",
            "https://www.googleapis.com/auth/devstorage.read_only",
            "https://www.googleapis.com/auth/logging.write",
            "https://www.googleapis.com/auth/monitoring"
        ],
    },
});

// Export the Cluster name
export const clusterName = cluster.name;

// Manufacture a GKE-style kubeconfig. Note that this is slightly "different"
// because of the way GKE requires gcloud to be in the picture for cluster
// authentication (rather than using the client cert/key directly).
export const kubeconfig = pulumi.
    all([ cluster.name, cluster.endpoint, cluster.masterAuth ]).
    apply(([ name, endpoint, masterAuth ]) => {
        const context = `${gcp.config.project}_${gcp.config.zone}_${name}`;
        return `apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: ${masterAuth.clusterCaCertificate}
    server: https://${endpoint}
  name: ${context}
contexts:
- context:
    cluster: ${context}
    user: ${context}
  name: ${context}
current-context: ${context}
kind: Config
preferences: {}
users:
- name: ${context}
  user:
    auth-provider:
      config:
        cmd-args: config config-helper --format=json
        cmd-path: gcloud
        expiry-key: '{.credential.token_expiry}'
        token-key: '{.credential.access_token}'
      name: gcp
`;
    });

// Create a Kubernetes provider instance that uses our cluster from above.
const gkeClusterProvider = new k8s.Provider(name, {
    kubeconfig: kubeconfig,
});

const awsClusterProvider = new k8s.Provider(name+"aws", {
  kubeconfig: fs.readFileSync("/Users/jbeda/.kube/demo").toString(),
});

let clusters = [gkeClusterProvider, awsClusterProvider];

export let namespaces: any[] = [];
export let deployments: any[] = [];
export let services: any[] = [];
export let serviceAddresses: any[] = [];

let i = 0;
for (const clusterProvider of clusters) {
  let localName = name+i;
  // Create a Kubernetes Namespace
  const ns = new k8s.core.v1.Namespace(localName, {}, { provider: clusterProvider });

  // Export the Namespace name
   const namespaceName = ns.metadata.apply(m => m.name);
   namespaces.push(namespaceName);

  // Create a NGINX Deployment
  const appLabels = { appClass: name };
  const deployment = new k8s.apps.v1.Deployment(localName,
      {
          metadata: {
              namespace: namespaceName,
              labels: appLabels,
          },
          spec: {
              replicas: 1,
              selector: { matchLabels: appLabels },
              template: {
                  metadata: {
                      labels: appLabels,
                  },
                  spec: {
                      containers: [
                          {
                              name: name,
                              image: "gcr.io/kuar-demo/kuard-amd64:blue",
                              ports: [{ name: "http", containerPort: 8080 }]
                          }
                      ],
                  }
              }
          },
      },
      {
          provider: clusterProvider,
      }
  );

  // Export the Deployment name
  const deploymentName = deployment.metadata.apply(m => m.name);
  deployments.push(deploymentName);

  // Create a LoadBalancer Service for the NGINX Deployment
  const service = new k8s.core.v1.Service(localName,
      {
          metadata: {
              labels: appLabels,
              namespace: namespaceName,
          },
          spec: {
              type: "LoadBalancer",
              ports: [{ port: 80, targetPort: "http" }],
              selector: appLabels,
          },
      },
      {
          provider: clusterProvider,
      }
  );

  // Export the Service name and public LoadBalancer endpoint
  const serviceName = service.metadata.apply(m => m.name);
  services.push(serviceName);

  const servicePublicIP = service.status.apply(s => (s.loadBalancer.ingress[0].ip || s.loadBalancer.ingress[0].hostname))
  serviceAddresses.push(servicePublicIP);

  i++;
}