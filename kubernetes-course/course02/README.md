# Kubernetes Tutorial for Beginners [FULL COURSE in 4 Hours]

Course link: https://www.youtube.com/watch?v=X48VuDVv0do  
Git with a list of CLI commands using `kubectl`: https://gitlab.com/nanuchi/youtube-tutorial-series/-/blob/master/basic-kubectl-commands/cli-commands.md

**Kubernetes**, also known as **K8s**, is an open source container orchestration tool.

## Tools needed to follow the course

- [minikube](https://minikube.sigs.k8s.io/docs/start/) - it will be used to create a single node k8s cluster on your machine. It will install all the necessary tools, including docker on the node.
- [kubectl](https://kubernetes.io/docs/tasks/tools/) - to manage the kubernetes cluster.

## K8s Architecture

Whenever you feel lost when encounter terms like Node, Pod, Deployment, or Service,  refer to the [K8s Glossary](https://kubernetes.io/docs/reference/glossary/?fundamental=true)

Here we have a picture showing a basic K8s cluster:

![Kubernetes Cluster](https://kubernetes.io/images/docs/kubernetes-cluster-architecture.svg "Kubernetes Cluster")

- **Pod** - a group of one or more containers. Think of Pod as a wrapper of a container. In K8s you deploy Pods not containers directly.
- **Node** - a worker machine (virtual or physical). Each Node must have the following components installed: 
    - **kubelet**: an agent/process that manages the Pods on the Node.
    - **kube-proxy**: a network proxy to ensure networks rules on the Node and also provide communication to the Pods.
    - **container runtime** e.g., Docker: represented and abstracted in the picture above by the Container Runtime Interface (CRI).
- **Control Plane** - previously known as **master node** is intended to run only K8s processes, and it's the "brain" of K8s cluster. The Control Plane runs:
    - **kube-api-server**: also known as **API Server**, it exposes the K8s API. When you use `kubectl` tool to manage the cluster, it  interacts with this component.
    - **scheduler** or **kube-scheduler**: basically this component takes a request from the API Server that new Pods are available to creation. The schedule then evaluates and selects a node for them to run on, and makes a request to the kubelet on the selected node to create the Pod.
    - **Controller Manager** or **kube-controller-manager**: it's a component that runs **controller** processes. Controllers are control loops that watch the state of your cluster, then make or request changes where needed. Ex.: if a Pod "dies" it will request the scheduler to create a new Pod. It's a very complex component to understand and I highly recommend exploring the documentation for [controllers](https://kubernetes.io/docs/concepts/architecture/controller/).
    - **cloud-controller-manager**: it also runs **controller** processes but they are specific to a cloud provider, like AWS.
    - **etcd**: a key-value store of the cluster state. Every information about K8s cluster, configuration, current state, desired state1 is stored in etcd.

## Layers of abstraction

- Service;
- ReplicaSet;
- Pod;
- Container

## About the YAML files

Each configuration file consists of three parts:
- **metadata** - defined by you. Ex.:
```yaml
metadata:
  name: nginx-deployment
  labels: ...
```
- **specification** - defined by you. Ex.:
```yaml
spec:
  replicas: 2
  selector: ...
```
- **status** - this one is generated and added automatically by K8s. It holds the status of your deployment and uses it to reach the desired state getting information from etcd Ex.:
```yaml
status:
  availableReplicas: 2
  conditions: ...
  replicas: ...
```
to check the information above just type `kubectl get deployment <deployment-name> -o yaml`.

## About the mongo-deploy folder

The intent of the YAML files in this folder is to accomplish the diagram below. You will access the MongoDB Express from your browser and the request will follow this flow:

```mermaid
---
title: Simple K8s cluster
---
flowchart LR;
    id1[[Browser]] --> id2(Mongo Express External Service);
    id2 --> id3{{Mongo Express Pod}};
    id3 --> id4(Mongo DB Internal Service);
    id4 --> id5{{MongoDB Pod}}
```

Type `chmod +x deploy.sh` to give the file execution permission, and then run it to deploy the applications. The same might be done for the `destroy.sh` file.

