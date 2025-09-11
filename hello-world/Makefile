all: build secret config deploy service

.PHONY: build secret deploy service config redeploy

build:
	cd hello && docker build -t hello:latest . && minikube image load hello:latest

secret:
	kubectl apply -f k8s/secret.yaml

deploy:
	kubectl apply -f k8s/deployment.yaml

config:
	kubectl apply -f k8s/config.yaml

redeploy:
	kubectl rollout restart deployment hello-world

service:
	kubectl apply -f k8s/service.yaml

clean:
	kubectl delete -f k8s/service.yaml
	kubectl delete -f k8s/deployment.yaml
	kubectl delete -f k8s/secret.yaml
	kubectl delete -f k8s/config.yaml
	docker rmi hello:latest || true