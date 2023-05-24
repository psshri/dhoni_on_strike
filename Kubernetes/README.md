### run the following commands

minikube start
minikube profile list
minikube delete --profile <profile-name>
kubectl apply -f configMap.yaml
kubectl apply -f deployment.yaml

kubectl delete deployment dhoni-on-strike-deployment
kubectl delete configmap dhoni-on-strike-configmap

minikube stop
minikube status
