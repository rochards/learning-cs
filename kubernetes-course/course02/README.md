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