# gss (go sample service)

A toy api for local experimenting with a golang code running inside a minikube pod with open telemetry tracing written to a text file. The accompanying frontend is [here](https://github.com/eugenekatsov/gss-fe)

## Setup

You'll need the following:
* I use docker desktop [docker](https://www.docker.com/products/docker-desktop/)
* to run commands against your cluster [kubectl](https://kubernetes.io/docs/tasks/tools/)
* for local kubernetes [minikube](https://minikube.sigs.k8s.io/docs/)
* Charts are a great way to deploy and manage releases to your minikube cluster [helm](https://helm.sh)
* to watch pod logs [stern](https://github.com/stern/stern)


## Getting Started

1. Clone this repo
2. Docker build and tag, then push to docker hub
4. Update the repository name and tag in the helm chart values.yaml with the docker hub repo name and tag
5. `minikube start` to start your k8s environment
6. kubectl create secret to access your image, in our case it's named `regcred` see [here](https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/#create-a-secret-by-providing-credentials-on-the-command-line)
7. `helm install <CHOOSE A RELEASE NAME> ./helm` from the root repo directory
8. Verify the pods are up and healthy: `kubectl get pods` or `describe pods`
9. Run `minikube service <YOUR SERVICE NAME> --url` to get the url
10. The frontend has some buttons for the fake CRUD calls
11. To see some logs use `stern <YOUR POD NAME>`
12. Enable metrics with `minikube addons enable metrics-server`
13. To have a general overview of your cluster: `minikube dashboard`