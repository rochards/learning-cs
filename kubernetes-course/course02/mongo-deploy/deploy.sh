#!/bin/bash

kubectl apply -f mongodb-secret.yaml
if kubectl get secret mongodb-secret > /dev/null 2>&1; then
    echo "Secret applied successfully."
else
    echo "Error applying Secret. Deployment aborted."
    exit 1
fi

kubectl apply -f mongodb-deploy.yaml
if kubectl get deployment  mongodb-deployment > /dev/null 2>&1; then
    echo "Deployment applied successfully."
else
    echo "Error applying Deployment."
fi