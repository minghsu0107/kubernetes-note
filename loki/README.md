# Loki in K8s
## Reference
[Get Logs into Loki](https://grafana.com/docs/loki/latest/getting-started/get-logs-into-loki/)
[Using Loki in Grafana](https://grafana.com/docs/grafana/latest/datasources/loki/)
## Usage
To create the base64 encoded config of Loki, we can:
```bash
cat config.yaml | base64
```

`config.yaml`:
```yaml
auth_enabled: false
chunk_store_config:
  max_look_back_period: 0
ingester:
  chunk_block_size: 262144
  chunk_idle_period: 3m
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

