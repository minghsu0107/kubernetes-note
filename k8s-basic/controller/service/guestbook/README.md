# Fullstack Demo
## Steps
1. Set up a "Redis master"
    - Create redis-master "deployment"
	- Create redis-master "service"
2. Set up a "Redis slave"
	- Create redis-master "deployment"
	- Create redis-master "slave"
3. Set up the "guestbook web frontend"
    - Create guestbook web frontend "deployment"
	- Expose frontend on an external IP address (LoadBalancer)
## Usage
```bash
kubectl apply -f .
```
Then visit http://<frontend-clusterIP>
## Note
The application code specify the redis master as follows:
```php
$host = 'redis-master';
  if (getenv('GET_HOSTS_FROM') == 'env') {
    $host = getenv('REDIS_MASTER_SERVICE_HOST');
  }
```
Since we use coreDNS, the redis master hostname will be `redis-master`.