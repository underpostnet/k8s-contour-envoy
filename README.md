![contour](https://www.nexodev.org/uploads/cloud/francisco-verdugo/k8s/fd5ed-contour1.png)

#### CLI project install

`go mod init main`

`go get github.com/dixonwille/wmenu/v5`

#### Usage

- Create cluster `kind create cluster --name <name-cluster>`

- Verify clusters, context, and users `kubectl config view` and save data `kubectl config view --minify --flatten --context=<name-context> > out.txt`

- Get a list of resource types `kubectl api-resources`

- Verify contexts (cluster name, user, and namespace that you can access when you invoke) `kubectl config get-contexts`

- Select a context `kubectl config use-context <context-name>`

- Get all current nodes `kubectl get nodes`

- Download contour installation manifests `wget -O contour.yaml https://projectcontour.io/quickstart/v1.12.0/contour.yaml`

- Test manifest `kubectl apply -f contour.yaml --dry-run=client`

- Get node info `kubectl describe node`

- Install manifest `kubectl apply -f contour.yaml`

- Get all namespaces pods `kubectl get pods --all-namespaces` | `kubectl get pods --all-namespaces -o wide --field-selector spec.nodeName=<node-name>`

- Get logs a pod `kubectl logs <podname> --namespace <some_namespace>`

- Get projectcontour namespace info `kubectl -n projectcontour get deployment,daemonset,service` | `kubectl -n projectcontour get all,ingress`

- Set external ip service `kubectl patch svc <name-service> -n <namespace> -p '{\"spec\": {\"type\": \"LoadBalancer\", \"externalIPs\":[\"192.168.1.83\"]}}'`

- Fix pod `kubectl patch pod <name-pod> -n <namespace> -p '{\"apiVersion\": \"networking.k8s.io/v1beta1\"}'`

- Get pod yaml `kubectl get pod -n <namespace> <pod-name> -o yaml > out.txt`

- Delete pod `kubectl delete pod <pod_name> -n <namespace>`

- Restart deployment `kubectl rollout restart deployment <deployment_name> -n <namespace>`

- Delete kind cluster `kind delete cluster --name kind-underpost-cluster`

- Delete cluster `kubectl config delete-cluster <name-cluster>`

- Delete contex `kubectl config delete-context <name-cluster>`

- Delete users `kubectl config unset users.<name-cluster>|<admin-cluster>`

#### Doc

- https://tanzu.vmware.com/developer/guides/service-routing-contour-to-ingress-and-beyond/

- https://blog.markvincze.com/how-to-use-envoy-as-a-load-balancer-in-kubernetes/

- https://projectcontour.io/contour_v190/
