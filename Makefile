
build: 
	GO111MODULE=on GOOS=linux CGO_ENABLED=0 go build -v && \
	docker build -t patnaikshekhar/leader-election-test:1.1 ./example && \
	docker build -t patnaikshekhar/leader-election-sidecar:1.1 .

push:
	docker push patnaikshekhar/leader-election-test:1.1 && \
	docker push patnaikshekhar/leader-election-sidecar:1.1

debug: build clean runp1 runp2

clean:
	-docker rm -vf le-p1-sidecar le-p1-test le-p2-sidecar le-p2-test

runp1:
	docker run -d --name le-p1-sidecar -v $$HOME/.kube/config:/config -p 8080:8080 --env POD_NAME=p1 --env POD_NAMESPACE=default patnaikshekhar/leader-election-sidecar:1 && \
	docker run -d --name le-p1-test --env LEADER_ELECTION_SERVER=localhost:8080 patnaikshekhar/leader-election-test:1

runp2:
	docker run -d --name le-p2-sidecar -v $$HOME/.kube/config:/config -p 8081:8080 --env POD_NAME=p2 --env POD_NAMESPACE=default patnaikshekhar/leader-election-sidecar:1 && \
	docker run -d --name le-p2-test --env LEADER_ELECTION_SERVER=localhost:8081 patnaikshekhar/leader-election-test:1