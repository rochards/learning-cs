#!/bin/bash

kubectl apply -f mongodb-secret.yaml
# if kubectl get secret mongodb-secret > /dev/null 2>&1; then
#     echo "'mongodb-secret' applied successfully."
# else
#     echo "Error applying 'mongodb-secret'."
#     exit 1
# fi

kubectl apply -f mongodb-deploy.yaml

kubectl apply -f mongodb-configmap.yaml

kubectl apply -f mongoexpress-deploy.yaml