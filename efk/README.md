# Elasticsearch, Fluentd, and Kibana Stack
[reference](https://www.digitalocean.com/community/tutorials/how-to-set-up-an-elasticsearch-fluentd-and-kibana-efk-logging-stack-on-kubernetes#step-4-—-creating-the-fluentd-daemonset)
## Prerequisites
- A working Kubernetes cluster
- **Make sure that /data/es already exists on each node**
- Turn off swap to make elasticsearch more efficient (optional)
## Usage
First, enable static volume provisioning:
```bash
kubectl apply -f elasticsearch/pv.yaml
kubectl apply -f elasticsearch/storageclass.yaml
```
Note that you need to change the node name of `.spec.nodeAffinity` to yours in `pv.yaml`.

Next, deploy all other resources.

For example, when deploying elasticsearch cluster, we can rollout the status:
```bash
kubectl rollout status sts/es-cluster -n kube-logging
```
Check elasticsearch cluster status:
```bash
kubectl port-forward es-cluster-0 9200:9200 -n kube-logging
curl http://localhost:9200/_cluster/state?pretty
```
## Using Kibana
```bash
kubectl port-forward <kibana-pod-name> 5601:5601 -n kube-logging
```
Visit http://localhost:5601 in browser.

1. Click on Discover in the left-hand navigation menu
2. Use the `logstash-*` wildcard pattern to capture all the log data in our Elasticsearch cluster
3. Configure which field Kibana will use to filter log data by time. In the dropdown, select the @timestamp field, and hit Create index pattern
4. Hit Discover in the left hand navigation menu again. You should see a histogram graph and some recent log entries
## Tuning Elasticsearch
- [heap size](https://www.elastic.co/guide/en/elasticsearch/reference/current/important-settings.html#heap-size-settings)
- [more configuration](https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html#_notes_for_production_use_and_defaults)