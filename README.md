# Multicluster Mesh Addon

The multicluster-mesh-addon is an enhanced service mesh addon created with [addon-framework](http://github.com/open-cluster-management-io/addon-framework), it is used to manage(discovery, deploy and federate) service meshes across multiple clusters in [Open Cluster Management(OCM)](https://open-cluster-management.io/). With multicluster-mesh-addon, you can unify the configuration and operation of your services spanning from single cluster to multiple clusters in hybrid cloud.

![multicluster-mesh-addon-overview](assets/multicluster-mesh-addon.png)

## Core Concepts

To simplify the configuration and operation of service meshes, we created the following three [custom resource definitions (CRDs)](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/) than you can configure from the OCM hub cluster. Behind the scenes, the multicluster-mesh-addon translates these high level objects into low level istio resources and then applied into the managed clusters.

1. **Mesh** - a `mesh` resource is mapping to a physical service mesh in a managed cluster, it contains the desired state and status of the backend service mesh.

    For each physical service mesh in a managed cluster, a mesh resource is created in the managed cluster namespace of hub cluster. An example of mesh resource would resemble the following yaml snippet:

    ```yaml
    apiVersion: mesh.open-cluster-management.io/v1alpha1
    kind: Mesh
    metadata:
      name: cluster1-istio-system-default
    spec:
      cluster: cluster1
      controlPlane:
        components: ["base", "istiod", "istio-ingress"]
        namespace: istio-system
        profiles: ["default"]
        version: 1.13.2
      meshMemberRoll: ["istio-apps"]
      meshProvider: Upstream Istio
      meshConfig:
        trustDomain: cluster.local
    ```

2. **MeshDeployment** - `meshdeployment` resource is used to deploy physical service meshes to managed cluster(s), it supports deploying multiple physical service meshes to different managed clusters with one mesh template.

    An example of meshdeployment resource would resemble the following yaml snippet:

    ```yaml
    apiVersion: mesh.open-cluster-management.io/v1alpha1
    kind: MeshDeployment
    metadata:
      name: istio
    spec:
      clusters: ["cluster1", "cluster2"]
      controlPlane:
        components: ["base", "istiod", "istio-ingress"]
        namespace: mesh-system
        profiles: ["default"]
        version: 1.13.2
      meshMemberRoll: ["mesh-apps"]
      meshProvider: Upstream Istio
    ```

3. **MeshFederation** - `meshfederation` resource is used to federate service meshes so that the physical service meshes located in one cluster or different clusters to securely share and manage traffic between meshes while maintaining strong administrative boundaries in a multi-tenant environment.

    An example of meshfederation resource would resemble the following yaml snippet:

    ```yaml
    apiVersion: mesh.open-cluster-management.io/v1alpha1
    kind: MeshFederation
    metadata:
      name: istio-federation
    spec:
      meshPeers:
      - peers:
        - name: cluster1-istio
          cluster: cluster1
        - name: cluster2-istio
          cluster: cluster2
      trustConfig:
        trustType: Complete
    ```

## Run Demo

To run the demo with KinD cluster, make sure the following command line tools are installed:

- git
- kubectl
- [kind](https://kind.sigs.k8s.io/)
- docker
- envsubst
- [istioctl](https://istio.io/latest/docs/setup/getting-started/#download)

then execute the following command:

```bash
cd demo && ./demo
```

## Getting Started

### Prerequisites

- Ensure [docker](https://docs.docker.com/get-started) 18.03+ is installed.
- Ensure [golang](https://golang.org/doc/install) 1.17+ is installed.
- Prepare an environment of [OCM](https://open-cluster-management.io/getting-started/core/) and login to the hub cluster with `kubectl` command line tool.
- Make sure at least one managed cluster imported to the OCM hub cluster.
- For mesh federation support, make sure at least two managed clusters are imported and the cloud provider must support the network load balancer IP so that the meshes spanning across managed clusters can be connected.

### Build and Deploy

1. Build and push docker image:

    ```bash
    make docker-build docker-push IMAGE=quay.io/<your_quayio_username>/multicluster-mesh-addon:latest
    ```

2. Deploy the multicluster-mesh-addon to hub cluster:

    ```
    make deploy IMAGE=quay.io/<your_quayio_username>/multicluster-mesh-addon:latest
    ```

## How to use

1. Mesh Discovery:

    If you have installed an istio service mesh in a managed cluster, then you should also find a mesh resource created in its namespace of hub cluster:

    ```bash
    # oc -n eks1 get mesh
    NAME                         CLUSTER    VERSION   PROVIDER                PEERS    AGE
    eks1-istio-system-default    eks1       1.13.2    Upstream Istio                   50s
    ```

2. Mesh Deployment:

    To deploy new istio service mesh(es) to managed clusters, create a `meshdeployment` resource by specifying `Upstream Istio` meshProvider and selecting the managed cluster(s). For example, create the following `meshdeployment` resource to deploy new istio service mesh(es) to managed cluster `eks1` and `eks1`:

    _Note_: For now, the multicluster-mesh-addon supports [upstream Istio](https://istio.io/) mesh provider.

    ```bash
    cat << EOF | oc apply -f -
    apiVersion: mesh.open-cluster-management.io/v1alpha1
    kind: MeshDeployment
    metadata:
      name: istio
      namespace: open-cluster-management
    spec:
      clusters: ["eks1", "eks2"]
      controlPlane:
        components: ["base", "istiod", "istio-ingress"]
        namespace: istio-system
        profiles: ["default"]
        version: 1.13.2
      meshMemberRoll: ["bookinfo"]
      meshProvider: Upstream Istio
    EOF
    ```

    Then verify the service meshes are created:

    ```bash
    # oc get mesh -A
    NAMESPACE   NAME          CLUSTER   VERSION   PROVIDER                 PEERS   AGE
    eks1        eks1-istio    eks1      1.13.2    Upstream Istio                   13s
    eks2        eks2-istio    eks2      1.13.2    Upstream Istio                   13s
    ```

4. Mesh Federation:

    To federate the istio service meshes in managed clusters, create a `meshfederation` resource in hub cluster by specifying the peers of istio mesh and and trustType of `Complete`. For example, federate `eks1-istio` and `eks1-istio` created in last step by creating a `meshfederation` resource with the following command:

    ```bash
    cat << EOF | oc apply -f -
    apiVersion: mesh.open-cluster-management.io/v1alpha1
    kind: MeshFederation
    metadata:
      name: istio-federation
      namespace: open-cluster-management
    spec:
      meshPeers:
      - peers:
        - name: eks1-istio
          cluster: eks1
        - name: eks2-istio
          cluster: eks2
      trustConfig:
        trustType: Complete
    EOF
    ```

    Finally, deploy [Bookinfo application](https://istio.io/latest/docs/examples/bookinfo/) spanning across managed clusters with the [instruction]((mesh-federation-verify-istio.md)) to verify the federated meshes are working as expected.

    _Note:_ currently the verify steps have to be executed in the managed cluster, the work for the service discovery and service federation is still in progress.

## TroubleShooting

If the traffic across meshes/clusters can't be routed successfully after creating `MeshFederation` resource, follow the following steps to find the root cause for different mesh providers:

  - Make sure the services are imported to target mesh cluster by creating corresponding `ServiceEntry` resources.
  - Make sure the eastwest gateway for each mesh is created and has a public loader balancer IP allocated.
  - Make sure `ServiceEntry` resource for remote service has the loader balancer IP of eastwest gateway for the peer mesh.
  - Make sure the secret named `cacerts` for the intermediate CA is created for each mesh.

If the cross meshes/clusters traffic is still not routed successfully after all the checks above, then open the access logs for the gateways by editting the MeshDeployment resources to add the following section to the `spec`:

```yaml
  meshConfig:
    proxyConfig:
      accessLogging:
        file: /dev/stdout
```

Then check the access logs for the ingress/egress/eastwest gateway for each mesh.

## Future Work

* Services and workloads discovery
* Federate services across meshes
* Deploy application across meshes
