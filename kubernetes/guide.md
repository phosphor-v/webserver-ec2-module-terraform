kind create cluster --name dev-cluster
kubectl create secret generic aws-credentials \
  --from-literal=aws_access_key_id=YOUR_ACCESS_KEY \
  --from-literal=aws_secret_access_key=YOUR_SECRET_KEY -n flux-system