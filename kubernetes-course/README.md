# Kubernetes Course - Full Beginners Tutorial (Containerize Your Apps!)

Click https://www.youtube.com/watch?v=d6WC5n9G_sM to access the full Youtube course.

Refer to the glossary in this link https://kubernetes.io/docs/reference/glossary/?fundamental=true whenever you encounter terms like Pod, Node, etc.

`kubectl` is the CLI for Kubernetes cluster management. To learn how to use it always consult the documentation https://kubernetes.io/docs/reference/kubectl/. There is also a great quick-reference available on https://kubernetes.io/docs/reference/kubectl/quick-reference/ and https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands

## Tools needed to follow the course

- [kubectl](https://kubernetes.io/docs/tasks/tools/) - to manage the kubernetes cluster
- [minikube](https://minikube.sigs.k8s.io/docs/start/) - it will be used to create a single node cluster on your machine. It will install all the necessary tools, including docker on the node.
- [Docker Desktop](https://docs.docker.com/desktop/) - it will be necessary because some Docker components, like custom images, will be created outside minikube

## Connection to your minikube node

If you would like to connect to the node machine just type: `minikube ssh`  
And to see all the docker containers running on the node, type: `docker ps`  
Running  `docker ps` right after creating a minikube node with `minikube start`, you'll see the default containers that would be present in the Control Plane and workers node:
```bash
k8s_storage-provisioner_storage-provisioner_kube-system_32152284-6c6f-48a1-b4f8-31c9eaa08157_3
k8s_coredns_coredns-5dd5756b68-kkq7w_kube-system_42f25373-2ec8-43f7-bf6f-59b829a6dbe0_1
k8s_kube-proxy_kube-proxy-4crpc_kube-system_9eb4e31d-7eef-465f-a663-8c03198a1e53_1
k8s_POD_coredns-5dd5756b68-kkq7w_kube-system_42f25373-2ec8-43f7-bf6f-59b829a6dbe0_1
k8s_POD_kube-proxy-4crpc_kube-system_9eb4e31d-7eef-465f-a663-8c03198a1e53_1
k8s_POD_storage-provisioner_kube-system_32152284-6c6f-48a1-b4f8-31c9eaa08157_1
k8s_etcd_etcd-minikube_kube-system_9aac5b5c8815def09a2ef9e37b89da55_1
k8s_kube-apiserver_kube-apiserver-minikube_kube-system_55b4bbe24dac3803a7379f9ae169d6ba_1
k8s_kube-controller-manager_kube-controller-manager-minikube_kube-system_7da72fc2e2cfb27aacf6cffd1c72da00_1
k8s_kube-scheduler_kube-scheduler-minikube_kube-system_75ac196d3709dde303d8a81c035c2c28_1
k8s_POD_kube-apiserver-minikube_kube-system_55b4bbe24dac3803a7379f9ae169d6ba_1
k8s_POD_kube-scheduler-minikube_kube-system_75ac196d3709dde303d8a81c035c2c28_1
k8s_POD_etcd-minikube_kube-system_9aac5b5c8815def09a2ef9e37b89da55_1
k8s_POD_kube-controller-manager-minikube_kube-system_7da72fc2e2cfb27aacf6cffd1c72da00_1
```
some of them are very easy to spot, like **etcd**, **kube-apiserver**, **kube-scheduler**, **kube-controller-manager** for the Control Plane. **kube-proxy** would be present on a worker node:
![Kubernetes Cluster](https://kubernetes.io/images/docs/components-of-kubernetes.svg "Kubernetes Cluster")

and a more detailed picture:
![Kubernetes Cluster](https://kubernetes.io/images/docs/kubernetes-cluster-architecture.svg "Kubernetes Cluster")

## Some key considerations

### About Kubernetes in its objects in general
- It's better to create Pods using [Deployments](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/) rather than just `kubectl run <pod-name> --image=<image-name>`. Deployments provide more control over Pod lifecycle management and replica sets.
- Creating a Deployment also creates a [ReplicaSet](https://kubernetes.io/docs/concepts/workloads/controllers/replicaset/). A ReplicaSet's purpose is to maintain a stable set of replica Pods running at any given time. 
- When creating Pods through Deployments, the Pod's name is like **nginx-deployment-6d6565499c-tng44**, where:
    - **nginx-deployment** -> is the Deployment's name.
    - **nginx-deployment-6d6565499c** -> is the ReplicaSet's name.
    - **nginx-deployment-6d6565499c-tng44** -> is the Pod's name.
- A [Service](https://kubernetes.io/docs/concepts/services-networking/service/) is the method for exposing a single or group of Pods over a network, with `ClusterIP` being the default type.

### About web-hello folder
- To deploy this app on a Kubernetes cluster, first you will have to build the docker image and then push it to DockerHub. Or just download the already built image [here](https://hub.docker.com/repository/docker/rochards/web-hello/general)

## kubectl commands I used

- `kubectl cluster-info` - to get information about the cluster.
- `kubectl get nodes` - to get info about the machines (nodes)
- `kubectl get namespaces` - to list all available namespaces
- `kubectl get pods` - to list pods in the default namespace
- `kubectl get pods --namespace=<choosed>` - to list pods in the namespace called choosed. Ex.: `kubectl get pods --namespace=kube-system`
- `kubectl run <pod-name> --image=<image-name>` - to create and run a new pod in the default namespace with an image available on Docker Hub (by default). Ex.: `kubectl run nginx --image=nginx`
- `kubectl describe pod <pod-name>` - to list a lot more information about the pod. Ex.: `kubectl describe pod nginx`
- `kubectl delete pod <pod-name>` - to delete the running pod. Ex.: `kubectl delete pod nginx`
- `kubectl delete deployment <deployment-name>` - to delete the running deployment. Ex.: `kubectl delete deployment nginx-deployment`
- `kubectl create deployment <deployment-name> --image=<image-name>` - to create a deployment in the default namespace from an image available on Docker Hub (by default). Ex.: `kubectl create deployment nginx-deployment --image=nginx`
- `kubectl get deployment` - to list deployments
- `kubectl describe deployment <deployment-name>` - to list a lot more information about the deployment: Ex.: `kubectl describe deployment nginx-deployment`
- `kubectl get service` - to list services
- `kubectl describe service <service-name>` - to list a lot more information about services. Ex.: `kubectl describe service nginx-service`
- `kubectl expose deployment <deployment-name> --port=<number>` - to create a Service for the specified deployment on the choosed port. Ex.: `kubectl expose deployment k8s-web-hello --port=3000` -> in this case `port` and `targetPort` are both set to 3000, and by the default the Service type will be `ClusterIP`
- `kubectl scale deployment <deployment-name> --replicas=<quantity>` - to change the numbers of Pods of your deployment .Ex.: `kubectl scale deployment k8s-web-hello --replicas=4`
- `kubectl apply -f <file.yaml>` - to apply the configuration of a `.yaml` file. Ex.: `kubectl apply -f web-hello-deployment.yaml`