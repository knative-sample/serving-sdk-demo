all: manager

manager:
	@echo "build k8s manager"
	go build -o bin/controller cmd/main.go
