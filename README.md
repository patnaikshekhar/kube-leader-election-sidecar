# Leader Election Sidecar for Kubernetes

This is a sample sidecar which can be injected into your apps that require leader election.

The sidecar exposes a simple http interface (localhost) which can be polled to check if
the container is the leader. It uses the leader election and resource lock functionalities
in the client go library.

```sh
kubectl apply -f examples/k8s.yaml
```

The example consists of a sample nodejs application that uses the sidecar. It communicates with the sidecar over HTTP.
