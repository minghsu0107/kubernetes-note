# Secret Demo

## Creating Secret using Kubectl & Consuming it from "volumes" inside Pod
```s
echo -n 'admin' > username.txt
echo -n 'pa$$w00rd' > password.txt

kubectl create secret generic nginx-secret-vol --from-file=username.txt --from-file=password.txt
```
```s
kubectl get secrets
kubectl describe secrets nginx-secret-vol
```
```s
kubectl create -f nginx-pod-secret-vol.yaml
```
```s
# Display
kubectl get po
kubectl get secrets
kubectl describe pod nginx-pod-secret-vol
```
```s
# Validate from "inside" the pod
kubectl exec nginx-pod-secret-vol -it /bin/sh
cd /etc/confidential
ls 
cat username.txt
cat password.txt
exit
```
(OR)
```s
# Validate from "outside" the pod
kubectl exec nginx-pod-secret-vol ls /etc/confidential
kubectl exec nginx-pod-secret-vol cat /etc/confidential/username.txt
kubectl exec nginx-pod-secret-vol cat /etc/confidential/password.txt
```

## Creating Secret "manually" using YAML file & Consuming it from "environment variables" inside Pod

```s
# Encoding secret
echo -n 'admin' | base64
echo -n 'pa$$w00rd' | base64
```
```s
kubectl create -f redis-secret-env.yaml
kubectl get secret
kubectl describe secret redis-secret-env
```
```s
kubectl create -f redis-pod-secret-env.yaml
```
```s
# Display
kubectl get pods
kubectl get secrets
kubectl describe pod redis-pod-secret-env
```
```s
# Validate from "inside" the pod
kubectl exec redis-pod-secret-env -it /bin/sh
env | grep  SECRET
exit
```
(OR)
```s
# Validate from "outside" the pod
kubectl exec redis-pod-secret-env env | grep SECRET
```

## Cleanup
```s
# Delete secrets
kubectl delete secrets nginx-secret-vol redis-secret-env

# Delete pods
kubectl delete pods nginx-pod-secret-vol redis-pod-secret-env

# Validate
kubectl get pods
kubectl get secrets
```
 