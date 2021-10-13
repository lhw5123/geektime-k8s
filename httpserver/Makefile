export tag=v1.0
root:
	export ROOT=github.com/lhw5123/geektime-k8s

build:
	echo "building httpserver"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .

release: build
	echo "building httpserver container"
	docker build -t geektime-k8s-hev1n/httpserver:${tag} .

push: release
	echo "pusing geektime-k8s-hev1n/httpserver"
	docker push geektime-k8s-hev1n/httpserver:v1.0