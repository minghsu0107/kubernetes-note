# Load Balancer Demo
 There is no LoadBalancer integrated if you are not using AWS or Google Cloud. 
## Usage
```bash
kubectl apply -f .
kubectl get pod -l app=nginx-app
kubectl get deploy -l app=nginx-app 
kubectl get service -l app=nginx-app
kubectl describe service my-service
```
## Testing
```bash
# To get inside the pod
kubectl exec -it [pod-name] -- /bin/sh

# Create test HTML page
cat <<EOF > /usr/share/nginx/html/test.html
<!DOCTYPE html>
<html>
<head>
<title>Testing..</title>
</head>
<body>
<h1 style="color:rgb(90,70,250);">Hello, Kubernetes...!</h1>
<h2>Load Balancer is working successfully. Congratulations, you passed :-) </h2>
</body>
</html>
EOF
exit
```
### Test using load-balancer-ip
```
http://load-balancer-ip
http://load-balancer-ip/test.html
```
### Testing using nodePort
```
http://nodeip:nodeport
http://nodeip:nodeport/test.html
```
## Clean Up
```bash
kubectl delete -f .
kubectl get pod
kubectl get deploy 
kubectl get service
```