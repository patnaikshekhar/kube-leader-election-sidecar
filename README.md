# Leader Election Sidecar for Kubernetes

This is a sample sidecar which can be injected into your apps that require leader election.

The sidecar exposes a simple http interface (localhost) which can be polled to check if
the container is the leader. It uses the leader election and resource lock functionalities
in the client go library.


