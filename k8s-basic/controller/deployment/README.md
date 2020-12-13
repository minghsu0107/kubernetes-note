# Deployment Demo
## Deployment

```s
# 透過 "--record" 參數來保留後續的 revision history
kubectl apply -f nginx-deployment.yaml --record
```

接著透過 kubectl 來查詢剛剛佈署的 deployment 相關資訊:
```s
# NAME： 列出了在目前 namespace 中的 deployment 清單
# DESIRED： 使用者所宣告的 desired status
# CURRENT： 表示目前有多少個 pod 副本在運行
# UP-TO-DATE： 表示目前有多個個 pod 副本已經達到 desired status
# AVAILABLE： 表示目前有多個 pod 副本已經可以開始提供服務
# AGE： 顯示目前 pod 運行的時間
$ kubectl get deployment
NAME               DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
nginx-deployment   3         3         3            0           15s
```

即時監控 deployment 佈署的狀況:
```s
kubectl rollout status deployment/nginx-deployment
```
```s
# 若rollout fail，則上一個指令回傳的 exit code 不等於 0
echo $?
1

# 若上一個指令的 exit code 為 0 表示 deployment 目前為 complete 狀態
echo $?
0
```

Check replicaset:
```s
# deployment controller 主動為 replicaset 新增了一個 "pod-template-hash" label
# ReplicaSet 的名稱會是 [DEPLOYMENT-NAME]-[POD-TEMPLATE-HASH-VALUE]
kubectl get rs
NAME                          DESIRED   CURRENT   READY     AGE
nginx-deployment-67594d6bf6   3         3         3         9m

# 檢視 ReplicaSet 的細節
kubectl describe rs/nginx-deployment-67594d6bf6
```
```s
# deployment controller 同時也幫 pod 增加了一個 "pod-template-hash" label，hash value 也是相同的
kubectl get pod --show-labels
```

在另外一個 namespace 中建立相同的 Deployment依然會拿到同樣 value 的 pod-template-hash:
```s
# 建立一個新的 namespace，名稱為 deployment-test
kubectl create namespace deployment-test
namespace/deployment-test created

# 將同樣的 deployment 宣告佈署到 deployment-test namespace 中
kubectl -n deployment-test apply -f nginx-deployment.yaml

kubectl -n deployment-test get deployment
kubectl -n deployment-test get rs
kubectl -n deployment-test get pod --show-labels

# 完成實驗，移除 namespace
kubectl delete ns/deployment-test
```
原因是因為 `.spec.template` 並沒有改變，因此透過 hash 產生出來的 value 當然也不會改變。

## 更新 Deployment
1. 修改 replica 的數量 (`.spec.replicas`)
2. 修改 template 的內容 (`.spec.template`)

```s
# 將 container image 的版本改成 nginx:1.9.1
kubectl set image deployment/nginx-deployment nginx=nginx:1.9.1 --record
kubectl set resources deployment/nginx-deployment -c='*-nginx-container' --limits=cpu=500m,memory=5Gi

kubectl rollout status deployment/nginx-deployment
```
```s
# replace deployment with new-nginx.yaml
kubectl replace -f new-nginx.yaml --record
```
```s
# edit directly
kubectl edit deployment/nginx-deployment --record
```

k8s 在運作機制上會遵守以下原則：

- Deployment ensures that only a certain number of Pods are down while they are being updated. By default, it ensures that at least 75% of the desired number of Pods are up (25% max unavailable).
- Deployment ensures that only a certain number of Pods are created above the desired number of Pods. By default, it ensures that at most 125% of the desired number of Pods are up (25% max surge).1
## 暫停 & 恢復 Deployment
```s
kubectl rollout pause deployment/nginx-deployment
kubectl rollout resume deployment/nginx-deployment
```
若是 rollout 目前已經暫停，那就連 rollback 也無法做了，因為 rollback 也會觸發 rollout 的發生。
## Rollout
```s
# 透過 rollout history 指令可以看出曾經下過什麼指令
kubectl rollout history deployment/nginx-deployment

# 檢視 rollout hisitory 2nd revision 的細節
kubectl rollout history deployment/nginx-deployment --revision=2
```

Undo the last deployment:
```s
kubectl rollout undo deployment/nginx-deployment
```

Rolling back to the 2nd revision:
```s
kubectl rollout undo deployment/nginx-deployment --to-revision=2
```

## Scaling
### Scale out
```s
# 設定目前的 pod repplica 為 10
kubectl scale deployment/nginx-deployment --replicas=10
```
### Horizontal Pod Autoscaling
```s
kubectl autoscale deployment nginx-deployment --min=10 --max=20 --cpu-percent=80
horizontalpodautoscaler.autoscaling/nginx-deployment autoscaled

kubectl get hpa
NAME               REFERENCE                     TARGETS         MINPODS   MAXPODS   REPLICAS   AGE
nginx-deployment   Deployment/nginx-deployment   <unknown>/80%   10        20        0          5s
```