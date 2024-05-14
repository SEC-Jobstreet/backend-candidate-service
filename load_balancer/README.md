# Load balancer

## Tools
- Docker
- Kubernetes (minikube)

## Set up:
- Install Docker
- Install Minikube
- Install kubectl

### Set up database:
```bash
cd load_balancer
kubectl apply db/
```

### Set up service:
```bash
kubectl apply service/
```

### To get the service url:
```bash
minikube service job-service --url
```

## To test if that work
```bash
# Get all pods
kubectl get pods
# Select 1 pod name
kubetctl delete pod <name>
```
New pods will be created to match the replica numbers

## Testing the load balancing:
- First, you have the correct url of that service with **minikube service job-service --url**
- Then you can do this command
```
#curl url/pods
curl http://192.168.49.2:30183/pods
```
- Do this curl several time, it will show the pod name of each pod that the load balancer is calling to