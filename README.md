# k8s playground
This project aims to deploy a simple project (which is visitor counter) on kubernetes.

### Deploy steps
`kubectl apply -f pv.yaml`

`kubectl apply -f pvc.yaml`

`kubectl apply -f secret.yaml`

`kubectl apply -f redis-deployment.yaml`

`kubectl apply -f redis-service.yaml`

`kubectl apply -f config-map.yaml`

`kubectl apply -f visitors-deployment.yaml`

`kubectl apply -f visitors-service.yaml`

### Test
#### using net-utils
`kubectl run -i -t net-utils --image=saman2000h/net-utils:1.2`

bash-5.1# `curl visitors-service/api`
#### using port-forward
`kubectl port-forward service/visitors-service 8888:80`

`curl localhost:8888/api`
