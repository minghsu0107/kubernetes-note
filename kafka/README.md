# Kafka Cluster
[reference 1](https://kow3ns.github.io/kubernetes-zookeeper/manifests/)
[reference 2](https://kubernetes.io/blog/2017/09/kubernetes-statefulsets-daemonsets/)

We will deploy Kafka in default namespace and ZooKeeper in kafka namespace.
## Test
### Zookeeper
```bash
# show zookeeper pods
for i in 0 1 2; do kubectl exec zk-$i -n kafka -- hostname -f; done
# The servers in a ZooKeeper ensemble use natural numbers as unique identifiers
# and store each server's identifier in a file called myid in the server's data directory
for i in 0 1 2; do echo "myid zk-$i";kubectl exec zk-$i -n kafka -- cat /var/lib/zookeeper/data/myid; done
# show config
kubectl exec zk-0 -n kafka -- cat /opt/zookeeper/conf/zoo.cfg
# write world to the path /hello on the zk-0 Pod in the ensemble
kubectl exec zk-0 -n kafka zkCli.sh create /hello world
# get  the data from the zk-1 Pod
kubectl exec zk-1 -n kafka zkCli.sh get /hello
```
### Kafka
```bash
# create a topic
kubectl -n kafka exec -ti testclient -- ./bin/kafka-topics.sh --zookeeper zk-cs.kafka.svc.cluster.local:2181 --topic messages --create --partitions 1 --replication-factor 3
# list topics, should have __consumer_offsets and messages
kubectl -n kafka exec -ti testclient -- ./bin/kafka-topics.sh --list --zookeeper zk-cs.kafka.svc.cluster.local:2181
# describe a topic
kubectl -n kafka exec -ti testclient -- ./bin/kafka-topics.sh --topic messages --describe  --zookeeper zk-cs.kafka.svc.cluster.local:2181
# list broker ids
kubectl -n kafka exec -ti testclient -- ./bin/zookeeper-shell.sh zk-cs.kafka.svc.cluster.local:2181 ls /brokers/ids
# list all topics on brokers
kubectl -n kafka exec -ti testclient -- ./bin/zookeeper-shell.sh zk-cs.kafka.svc.cluster.local:2181 ls /brokers/topics
# describe a broker
kubectl -n kafka exec -ti testclient -- ./bin/zookeeper-shell.sh zk-cs.kafka.svc.cluster.local:2181 get /brokers/ids/2
```
### Consumer and Producer
```bash
# start consumer
kubectl -n kafka exec -ti testclient -- ./bin/kafka-console-consumer.sh --bootstrap-server kafka-0.kafka-hs.default.svc.cluster.local:9093 --topic messages --from-beginning
# start producer
kubectl -n kafka exec -ti testclient -- ./bin/kafka-console-producer.sh --broker-list kafka-0.kafka-hs.default.svc.cluster.local:9093,kafka-1.kafka-hs.default.svc.cluster.local:9093,kafka-2.kafka-hs.default.svc.cluster.local:9093 --topic messages
```
Send messages in producer:
```
>hello
>world
```
You should receive messages in consumer:
```
hello
world
```
### Others
One can also test the Kafka cluster using [this helper](https://github.com/rmoff/kafka-listeners/tree/master/golang).