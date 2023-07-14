# Pod Demo
## Create and display Pods
```s
kubectl create -f nginx-pod.yaml
kubectl get pod
kubectl get pod -o wide
kubectl get pod nginx-pod -o yaml
kubectl describe pod nginx-pod
```
## Test & Delete
### To get inside the pod
```s
kubectl exec -it nginx-pod -- /bin/sh
```

### Create test HTML page
```s
cat <<EOF > /usr/share/nginx/html/test.html
<!DOCTYPE html>
<html>
<head>
<title>Testing..</title>
</head>
<body>
<h1 style="color:rgb(90,70,250);">Hello, Kubernetes...!</h1>
<h2>Congratulations, you passed :-) </h2>
</body>
</html>
EOF
exit
```

### Expose Pods using NodePort service
```s
kubectl expose pod nginx-pod --type=NodePort --port=80
```

### Display Service and find NodePort
Check the nodePort by:
```s
kubectl describe svc nginx-pod
```
If the nodePort is 30758, it means that we expose port 30758 on all nodes and foward all traffic to `nginx-pod`.
### Open Web-browser and access webapge using 
Visit `http://nodeip:nodeport/test.html`.

### Delete pod & svc
```s
kubectl delete svc nginx-pod
kubectl delete pod nginx-pod
```

## Deploy using k3d
### Steps
We need an agent for us to reach the pod outside the k3d cluster.

1. Create a cluster, mapping the port 30080 from agent-0 to localhost:8082
    - `k3d cluster create mycluster --agents 1 -p '8082:30080@agent:0'`
2. Create a NodePort service for it
    - `kubectl apply -f node-port-svc.yaml`
    - nodePort
        - On top of having a cluster-internal IP, expose the service on a port on each node of the cluster (the same port on each node). You'll be able to contact the service on any `<nodeIP>:nodePortaddress`. So nodePort is alse the service port which can be accessed by the node ip by others with external ip
        - `<node-ip>:<nodePort> -> <pod>:<targetPort>`
        - If we don't specify it, Kubernetes will choose it randomly for us
    - port
        - The port that the service is exposed on the service’s cluster ip (virtual ip). Port is the service port which is accessed by others with cluster ip. (can only be accessed inside the cluster)
        - `<virtual-ip>:<port> -> <pod>:<targetPort>`
    - targetPort
        - The port on the pod that the service should proxy traffic to.
3. Create `nginx-pod` and the test html
    - `kubectl create -f nginx-pod.yaml`
    - Create test HTML page (same as above)
4. Visit `localhost:8082/test.html`
### Simplest Way
Use port fowarding:
```s
kubectl port-forward nginx-pod 8888:80
```
Now, by visiting `localhost:8888`, we can reach `nginx-pod`.
