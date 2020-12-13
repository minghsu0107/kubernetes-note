
# ConfigMaps Demo
Overview:
1. Creating Configmap from "multiple files" & Consuming it inside Pod from "volumes" 
    - Create Configmap `nginx-configmap-vol` from "multiple files"
    - Consume `nginx-configmap-vol` configmap inside Pod from "volumes" 
    - Create | Display | Validate

2. Creating Configmap from "literal values" & Consuming it inside Pod from "environment variables" 
    -  Create configmap `redis-configmap-env` from "literal values"
    -  Consume `redis-configmap-env` configmap inside pod from environment Variables inside pod
    - Create | Display | Validate
3. Creating Configmap from `configmaps.yaml` & Consuming it inside Pod from "environment variables"
    - Create configmap `configmaps.yaml`
    - Consum configmap `spcial-config` and `env-config` inside Pod
4. Cleanup
    - Delete configmaps
    - Delete pods
    - Validate

## 1
Creating Configmap from "multiple files" & Consuming it inside Pod from "volumes" 

Create Configmap "nginx-configmap-vol" from "multiple files":
```s
echo -n 'Non-sensitive data inside file-1' > file-1.txt
echo -n 'Non-sensitive data inside file-2' > file-2.txt
```
```s
kubectl create configmap nginx-configmap-vol --from-file=file-1.txt --from-file=file-2.txt
```
```s
kubectl get configmaps
kubectl describe configmaps nginx-configmap-vol
```

Consume above "nginx-configmap-vol" configmap inside Pod from "volumes" in `nginx-pod-configmap-vol.yaml`.

Create | Display | Validate:
```s
kubectl create -f nginx-pod-configmap-vol.yaml
```
```s
kubectl get po
kubectl get configmaps
kubectl describe pod nginx-pod-configmap-vol
```
```s
# validate from "inside" the pod
kubectl exec nginx-pod-configmap-vol -it /bin/sh
cd /etc/non-sensitive-data
ls 
cat file-a.txt
cat file-b.txt
exit
```
(OR)
```s
# Validate from "outside" the pod
kubectl exec nginx-pod-configmap-vol ls /etc/non-sensitive-data
kubectl exec nginx-pod-configmap-vol cat /etc/non-sensitive-data/file-a.txt
kubectl exec nginx-pod-configmap-vol cat /etc/non-sensitive-data/file-b.txt
```

## 2
Creating Configmap from "literal values" & Consuming it inside Pod from "environment variables".


Create configmap `redis-configmap-env` from "literal values"
```s
kubectl create configmap redis-configmap-env --from-literal=file.1=file.a --from-literal=file.2=file.b
```
```s
kubectl get configmap
kubectl describe configmap redis-configmap-env
```

Consume `redis-configmap-env` configmap inside `redis-pod-configmap-env.yaml`.

Create | Display | Validate:
```s
kubectl create -f redis-pod-configmap-env.yaml
```
```s
kubectl get pods
kubectl get configmaps
kubectl describe pod redis-pod-configmap-env
```
```s
kubectl exec redis-pod-configmap-env -it /bin/sh
env | grep  FILE
exit
```
(OR)
```s
# Validate from "outside" the pod
kubectl exec redis-pod-configmap-env env | grep FILE
```
## 3
```s
kubectl create -f configmaps.yaml
```
```s
kubectl create -f box-pod-configmap-yml.yaml
```
```s
kubectl describe pod dapi-test-pod
```
## 4
Cleanup.
```s
# Delete configmaps
kubectl delete configmaps nginx-configmap-vol redis-configmap-env special-config env-config
```
```s
# Delete pods
kubectl delete pods nginx-pod-configmap-vol redis-pod-configmap-env dapi-test-pod
```
```s
# Validate
kubectl get pods
kubectl get configmaps
```

 