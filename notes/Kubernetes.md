[TOC]

# Kubernetes 

​	Minikube是一种轻量级的Kubernetes实现，他在本地机器上创建一个vm，并部署一个仅包含一个节点的简单集群。

# 常用指令

`minikube version`	`minikube start # 开启集群`	

## 使用 kubectl 创建部署

​	与kubernetes通讯，需要用到kubectl命令行界面。`kubectl version`                                                                                           使用 kubectl 在 Kubernetes 上部署您的第一个应用程序。Deployment 指示 Kubernetes 如何创建和更新应用程序的实例。创建部署后，Kubernetes 控制平面会安排该部署中包含的应用程序实例在集群中的各个节点上运行。

## 创建一个应用程序

````
kubectl create deployment hello-node --image=k8s.gcr.io/echoserver:1.4
````

## 查看部署

```
kubectl get deployments
```

## 查看Pod

```
kubectl get pods
```

## 查看节点

```
kubectl get nodes
```

## 运行代理

```
kubectl proxy
```

## 查看应用程序的输出

```
curl http://localhost:8001/api/v1/namespaces/default/pods/$POD_NAME/proxy/
```

___

## 创建服务

```
kubectl expose deployment hello-node --type=LoadBalancer --port=8080
```

## 查看服务

```
kubectl get services
```

# Kubernetes Pod

​	Pod 是一种 Kubernetes 抽象，表示一组一个或多个应用程序容器（例如 Docker），*包括共享存储（卷）、IP 地址和有关如何运行它们的信息。*

`kubectl get pod	# 查看pod`

`kubectl describe pods	# 查看详细信息`

`kubectl proxy	# 运行代理`

`	curl http://localhost:8001/api/v1/namespaces/default/pods/$POD_NAME/proxy/	#查看输出`

`kubectl logs $POD_NAME	# 查看日志`

`kubectl exec $POD_NAME -- env	# 列出环境变量`

`kubectl exec -ti $POD_NAME -- bash	# 执行容器的bash`

`exit	# 离开`

# Node

​	Pod 总是在**Node**上运行。Node 是 Kubernetes 中的工作机器，可以是虚拟机或物理机。

​	一个节点可以有多个 Pod， Kubernetes 节点至少运行：

- Kubelet，一个负责 Kubernetes 控制平面和 Node 之间通信的进程；它管理在机器上运行的 Pod 和容器。
- 容器运行时（如 Docker）负责从注册表中提取容器镜像、解包容器并运行应用程序。

# 公开应用程序

1. ```
   $kubectl get deployments
   NAME                  READY   UP-TO-DATE   AVAILABLE   AGE
   kubernetes-bootcamp   1/1     1            1           2m35s
   ```

2. ```
   kubectl expose deployment/kubernetes-bootcamp --type="NodePort" --port 8080
   # 创建新服务并将其公开到外部流量
   ```

3. ```
   kubectl get services	# 查看服务
   NAME                  TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
   kubernetes            ClusterIP   10.96.0.1       <none>        443/TCP          4m1s
   kubernetes-bootcamp   NodePort    10.108.185.58   <none>        8080:30867/TCP   7s
   
   kubectl describe services/kubernetes-bootcamp	# 详情
   ```

4. ```
   curl ip:30867
   ```

## 获取标签索引pods、services

```
kubectl describe deployment
kubectl get pods -l app=kubernetes-bootcamp
kubectl get services -l app=kubernetes-bootcamp
```

更改标签名字

```
kubectl label pods $POD_NAME version=v1
kubectl describe pods $POD_NAME		# 查看详细的pods信息中的label中绑定了一条新的标签
kubectl get pods -l version=v1
```

## 删除服务

```
kubectl delete service -l app=kubernetes-bootcamp
```

