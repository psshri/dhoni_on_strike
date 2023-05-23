### run the following commands

minikube start
minikube profile list
kubectl apply -f configMap.yaml
kubectl apply -f deployment.yaml

kubectl delete deployment dhoni-on-strike-deployment
kubectl delete configmap dhoni-on-strike-configmap

minikube stop
minikube status


there is an issue with running the locally stored image, resolve this out