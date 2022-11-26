#### Doc

- https://tanzu.vmware.com/developer/guides/service-routing-contour-to-ingress-and-beyond/

- https://projectcontour.io/contour_v190/

#### Usage

- Create cluster `kind create cluster --name <name-cluster>`

- Verify clusters, context, and users `kubectl config view`

- Verify contexts (cluster name, user, and namespace that you can access when you invoke) `kubectl config get-contexts`

- Select a context `kubectl config use-context <context-name>`

- Get all current nodes `kubectl get nodes`

- Download contour installation manifests `wget -O contour.yaml https://projectcontour.io/quickstart/v1.12.0/contour.yaml`

- Test manifest `kubectl apply -f contour.yaml --dry-run=client`

- Get node info `kubectl describe node`

- Install manifest `kubectl apply -f contour.yaml`

- Get projectcontour namespace info `kubectl -n projectcontour get deployment,daemonset,service`
