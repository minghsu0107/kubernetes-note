# KEDA
KEDA is a Kubernetes-based Event Driven Autoscaler. With KEDA, you can drive the scaling of any container in Kubernetes based on the number of events needing to be processed.

KEDA is a single-purpose and lightweight component that can be added into any Kubernetes cluster. KEDA works alongside standard Kubernetes components like the Horizontal Pod Autoscaler and can extend functionality without overwriting or duplication. With KEDA you can explicitly map the apps you want to use event-driven scale, with other apps continuing to function. This makes KEDA a flexible and safe option to run alongside any number of any other Kubernetes applications or frameworks.
## Installation
```bash
kubectl apply -f https://github.com/kedacore/keda/releases/download/v2.0.0/keda-2.0.0.yaml
```
## Demo
First, deploy consumer. You will see that there is only one pod out there.
```bash
kubectl apply -f deploy-consumer.yaml
```
Push messages to RabbitMQ:
```bash
kubectl apply -f deploy-producer.yaml
```
Now, you can see that consumer automatically scales. You can inspect by:
```bash
kubectl get deploy/rabbitmq-consumer
kubectl get hpa
```