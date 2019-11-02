FROM alpine

COPY kube-leader-election-sidecar .

CMD ["./kube-leader-election-sidecar"]