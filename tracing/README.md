# Distributed Tracing (OpenCensus + Jaeger)
## Open Jaeger Dashboard
```bash=
kubectl -n tracing port-forward deploy/jaeger 30188:16686 --address 0.0.0.0
```
## Note
- [Opencensus Collector Configuration](https://github.com/census-instrumentation/opencensus-service/blob/master/exporter/README.md)
- [Jaeger Data Source in Grafana](https://grafana.com/docs/grafana/latest/datasources/jaeger/)
    - set Jaeger URL: `http://jaeger.tracing:16686`.
- Managing Elasticsearch backend
    - Automatically deletes older indices
    - [reference](https://github.com/jaegertracing/jaeger/tree/master/plugin/storage/es)
    - [docker image](https://hub.docker.com/r/jaegertracing/jaeger-es-index-cleaner)
    - `docker run -e TIMEOUT=120 jaegertracing/jaeger-es-index-cleaner 3 http://<es-host>:9200`
        - delete any indices older than 3 days
        - Environment variable TIMEOUT that sets the timeout in seconds for indices deletion