# Redis Cluster
[reference](https://rancher.com/blog/2019/deploying-redis-cluster)
## Usage
1. Apply `storageclass.yaml` and `pv.yaml`.
2. Apply `redis-sts.yaml` and `redis-svc.yaml`.
3. Check the pod deployment. If all work normally, we can start joining all nodes together:

First, obtain ips and ports of all nodes:
```bash
kubectl get pods -l app=redis-cluster -o jsonpath='{range.items[*]}{.status.podIP}:6379 '
```
Then join all nodes:
```bash
kubectl exec -it redis-cluster-0 -- redis-cli --cluster create --cluster-replicas 1 10.42.0.159:6379 10.42.1.88:6379 10.42.0.160:6379 10.42.1.89:6379 10.42.0.161:6379 10.42.1.90:6379
```

Finally, check the cluster info:
```bash
kubectl exec -it redis-cluster-0 -- redis-cli cluster info
```
```bash
for x in $(seq 0 5); do echo "redis-cluster-$x"; kubectl exec redis-cluster-$x -- redis-cli role; echo; done
```
## Note
- [Rewriting/compacting append-only files](https://redislabs.com/ebook/part-2-core-concepts/chapter-4-keeping-data-safe-and-ensuring-performance/4-1-persistence-options/4-1-3-rewritingcompacting-append-only-files/)
- [Grafana Dashboard for Redis Exporter](https://grafana.com/grafana/dashboards/763/revisions)