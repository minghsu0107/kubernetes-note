# Loki in K8s
## Reference
[Get Logs into Loki](https://grafana.com/docs/loki/latest/getting-started/get-logs-into-loki/)
[Using Loki in Grafana](https://grafana.com/docs/grafana/latest/datasources/loki/)
## Usage
To create the base64 encoded config of Loki, we can:
```bash
cat config.yaml | base64
```
### Promtail Configuration
`config.yaml`:
```yaml
auth_enabled: false
chunk_store_config:
  max_look_back_period: 0
ingester:
  # When this threshold is exceeded the head chunk block will be cut and compressed inside the chunk
  chunk_block_size: 262144
  # How long chunks should sit in-memory with no updates before being flushed 
  # if they don't hit the max block size. This means that half-empty chunks will still be flushed 
  # after a certain period as long as they receive no further activity.
  chunk_idle_period: 3m
  # How long chunks should be retained in-memory after they've been flushed.
  chunk_retain_period: 1m
  lifecycler:
    ring:
      kvstore:
        store: inmemory
      # The number of ingesters to write to and read from
      replication_factor: 1
  max_transfer_retries: 0
limits_config:
  enforce_metric_name: false
  reject_old_samples: true
  reject_old_samples_max_age: 168h
schema_config:
  configs:
  - from: "2018-04-15"
    index:
      period: 168h
      prefix: index_
    object_store: filesystem
    schema: v11
    store: boltdb
server:
  http_listen_port: 3100
storage_config:
  boltdb:
    directory: /data/loki/index
  filesystem:
    directory: /data/loki/chunks
table_manager:
  retention_deletes_enabled: false
  retention_period: 0
```

See more config [here](https://grafana.com/docs/loki/latest/configuration/#schema_config).
### Import Loki In Grafana
Go to datasource and add loki with url `http://loki-headless.logging.svc.cluster.local:3100`.

Now, go to `explore` and you start query lgos in the Loki dashboard.