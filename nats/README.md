# NATS Streaming
To minimize the single point of failure, NATS Streaming server can be run in Fault Tolerance mode. It works by having a group of servers with one acting as the active server (accessing the store) and handling all communication with clients, and all others acting as standby servers.

It is important to note that is not possible to run Nats Streaming as Fault Tolerance mode and Clustering mode at the same time.

However, we need NFS to deploy NATS Streaming with FT mode. Therefore, we demo clustering mode here.
## Usage
1. Apply `storageclass.yaml` and `pv.yaml`.
2. Apply `stan.yaml`.
3. Check logs: `kubectl logs stan-0 -c stan`.
## Note
- [Grafana dashboard for NATS server](https://grafana.com/grafana/dashboards/2279)

The configuration for FT mode:
```javascript
http: 8222

cluster {
    port: 6222
        routes [
        nats://stan-0.stan:6222
        nats://stan-1.stan:6222
        nats://stan-2.stan:6222
    ]
    cluster_advertise: $CLUSTER_ADVERTISE
    connect_retries: 10
}

streaming {
    id: test-cluster
    store: file
    dir: /data/stan/store
    ft_group_name: "test-cluster"
    file_options {
        buffer_size: 32mb
        sync_on_flush: false
        slice_max_bytes: 512mb
        parallel_recovery: 64
    }
    store_limits {
        max_channels: 10
        max_msgs: 0
        max_bytes: 256gb
        max_age: 1h
        max_subs: 128
    }  
}
```
If everything was setup properly, one of the servers will be the active node.