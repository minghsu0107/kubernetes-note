kubectl create secret generic rabbitmq-config --from-literal=erlang-cookie=rabbitmq-k8s-Dem0
kubectl create secret generic rabbitmq-admin --from-literal==user=user --from-literal==pass=pass