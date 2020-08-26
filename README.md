# Desafio Curso Microserviços - k8s

## Parte 1 - nginx

Pasta nginx

## Parte 2 - mysql

Pasta mysql

## Parte 3 - Server Golang

Pasta go

Imagem:
https://hub.docker.com/r/gmeneguz/cursomicrosservicos_desafio_go_http_server

## Parte 4 - Server Golang com Kubernetes

Para rodar no projeto no k8s:

```
cd go
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
# Minikube
minikube service go-service

```
