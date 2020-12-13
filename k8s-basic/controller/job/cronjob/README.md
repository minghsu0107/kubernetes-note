# Cronjob
## Usage
```bash
kubectl apply -f cronjob1.yaml
```
```bash
kubectl get cronjob
kubectl get pod -l=job-name=<job_name>
kubectl logs <pod_name>
```
## concurrencyPolicy
Choose Forbid if you don't want concurrent executions of your Job. When its time to trigger a Job as per the schedule and a Job instance is already running, the current iteration is skipped.

If you choose Replace as the concurrency policy, the current running Job will be stopped and a new Job will be spawned. 

Specifying Allow will let multiple Job instances run concurrently.