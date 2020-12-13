# DNS Demo
## Usage
```bash
kubectl apply -f dns-demo.yaml
```
## Testing
```bash
kubectl get all
kubectl get endpoints # return a list of <pod-ip>:<port>
```
Run curl pod:
```bash
kubectl run --rm --generator=run-pod/v1 curl --image=radial/busyboxplus:curl -i --tty
```
Inside curl pod:
```bash
nslookup svc-cluster.default.svc.cluster.local
nslookup svc-headless.default.svc.cluster.local
```

From host:
```bash
# 從 host level 使用 nslookup
$ nslookup
> server 10.233.0.3   # 切換 DNS server 到 k8s DNS svc
Default server: 10.233.0.3
Address: 10.233.0.3#53

# 查詢 "default" namespace 中的 "svc-cluster" service domain name
# (type = ClusterIP)
> svc-cluster.default.svc.cluster.local
Name:    svc-cluster.default.svc.cluster.local
Server:        10.233.0.3
Address:    10.233.0.3#53

Name:    svc-cluster.default.svc.cluster.local
Address: 10.233.49.77

# 查詢 "default" namespace 中的 "svc-headless" service domain name
# (clusterIP: None)
> svc-headless.default.svc.cluster.local
Server:        10.233.0.3
Address:    10.233.0.3#53

Name:    svc-headless.default.svc.cluster.local
Address: 10.233.103.201
Name:    svc-headless.default.svc.cluster.local
Address: 10.233.76.11
Name:    svc-headless.default.svc.cluster.local
Address: 10.233.76.10
```