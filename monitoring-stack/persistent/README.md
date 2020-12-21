# Persistent Monitoring Stack

This tutorial demos how we can save Prometheus and Grafana data in persistent volume.

Something worth noting:

1. We should fix Grafana user's permission (UID=472) when mounting volume
2. Make sure the mounted host folder is empty