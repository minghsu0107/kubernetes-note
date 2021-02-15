# Distributed Tracing (OpenCensus + Jaeger)
## Open Jaeger Dashboard
```bash=
kubectl -n tracing port-forward deploy/jaeger 30188:16686 --address 0.0.0.0
```
