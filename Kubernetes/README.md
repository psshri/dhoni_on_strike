### run the following commands

minikube start
minikube profile list
kubectl apply -f configMap.yaml
kubectl apply -f deployment.yaml

kubectl delete deployment dhoni-on-strike-deployment