##### 1.本质是什么？为什么存在？

**本质：**一个**容器管理平台**，并提供了大量的功能组件去更好管理和维护容器。

```markdown
# 纯容器的部署问题
1. 业务容器数量庞大，哪些容器部署在哪些节点,使用了哪些端口,如何记录和管理容器,需要登录到每台机器去管理？
2. 跨主机通信，多个机器中的容器之间如何相互调用，iptables规则手动维护？
3. 跨主机容器调用，如何配置？写死固定IP+PORT？
4. 如何实现业务高可用，多个容器对外提供服务如何实现负载均衡？
5. 容器中的业务中断了,如何感知并重启？
6. 如何实现滚动升级，以保证业务的连续性？
# 容器化部署以及容器化大规模应用后，难以维护，k8s技术出现。
```

##### 2.k8s能做什么?

```markdown
- 自我修复：一旦一个容器崩溃，可秒重新启动另一个容器；
- 弹性伸缩：根据需要，动态调整正在运行的容器数量；
- 服务发现：服务可自动发现所依赖的服务；
- 负载均衡：如果一个服务启动多个容器，自动实现请求的负载均衡；
- 版本回退：发现新发布版本有问题，可回退到原来的版本；
- 存储编排：根据容器自身需求自动创建存储卷;
```

##### 3.如何设计一个容器管理平台？

​	1）集群架构：**管理节点**（master）分发容器到**数据节点**(node)；

```markdown
集群至少需要两台机器（master主节点,node工作节点）：master根据维护者写的yaml文件，描述容器的运行,创建具体的容器到node工作节点中。
```

​	2）自动识别目标节点的资源状态，选择最合适的节点部署新容器；

​	3）能始终确保应用的副本（**运行的容器数量**）是健康，正确的；

​	4）容器内的负载均衡,反向代理如何配置；

##### 4.核心组件及作用？

1) 控制节点（master）：集群的控制平面，负责集群决策

```markdown
1) api-server:资源操作的入口，接收用户输入的命令，提供认证，授权，api注册和发现等机制；
2) scheduler:集群资源调度，按照预定的调度策略将Pod调度到相应的node节点；
3) controller-manager:维护集群的状态，包括程序部署安排，故障检测，自动扩展和滚动更新等；【安排任务】
4) etcd:存储集群中各种资源对象的信息；【存储信息】
5) kubectl:客户端命令，和api-server交互,告诉master节点去完成相应的操作；
eg: kubectl create -f install-nginx.yml
```

2. 工作节点（node）：集群的数据平面，负责为容器提供运行环境

```markdown
1) kubelet:维护容器生命周期，控制Docker来创建，更新和销毁容器；
2) kube-proxy:提供集群内部的服务发现和负载均衡；【访问服务】
3) docker:负责节点上容器的运行及各种操作；
```

3. 相关概念

   1) pod 

   ```markdown
   pod说明：
   	- K8S 运行的最小控制单元,每个pod都有IP且删除pod IP会发生变化;
   	- pod内有一个根容器，且有一个/多个容器；
   	- pod内的所有容器共享根容器的命名空间;
   	- pod内容器的网络地址，由根容器提供
   pod的四种形态：
   	- 无状态：无需数据卷；
   	- 持久化容器内的数据，以数据卷的形式放到宿主机（node节点）；
   	- 一个pod下的多个容器共享一个卷的数据；
   	- 多个容器共享多个卷的数据；
   ```

   <img src="/Users/jacktaoyang/Library/Application Support/typora-user-images/image-20230129175201340.png" alt="image-20230129175201340"  />

   2) 其他名词

   ```markdown
   Controller：控制器，管理Pod的开启状态，数量等;
   Service：Pod对外服务的统一入口;
   Label：标签，对Pod分类;
   NameSpace：命名空间（资源组），区分Pod的运行环境;
   ```

##### 5.k8s组件工作流程

<img src="/Users/jacktaoyang/Library/Application Support/typora-user-images/image-20230129173216961.png" alt="image-20230129173216961" style="zoom:50%;" />

##### 6.环境搭建

###### 6.1 环境初始化

```markdown
1.docker环境配置
2.配置hosts
```

###### 6.2 安装k8s集群初始化工具（mater&node）

```markdown
# kubelet组件：在具体的机器上增删改查pod,主节点和node节点都可以运行;
# kubeadm：自动拉取k8s基础镜像的工具
# kubectl：管理,维护k8s客户端和服务端交互的命令行工具(会一直监控etcd数据库变化)
```

###### 6.3 k8s集群初始化（master）

```markdown
# kubeadm init 命令
```

###### 6.4 集群状态就绪配置

```markdown
# 部署Flannel网络插件
# 1.下载网络插件的yaml及配置文件
git clone --depth 1 https://github.com/coreos/flannel.git
# 2.在k8s主节点上应用yaml(基于yaml创建pod)
# 3.需要修改pod运行网络时，修改配置文件[flannel-master/Docmentation/kube-flannel.yml]
# 4.配置k8s命令补全
  yum install bash-completion -y
  source /usr/share/bash-completion/bash_completion
  source < (kubectl completion bash)
  echo "source <(kubectl completion bash)" >> ~/.bashrc
```

> 1.修改配置文件的网络插件信息

<img src="/Users/jacktaoyang/Library/Application Support/typora-user-images/image-20230130164951959.png" alt="image-20230130164951959" style="zoom:50%;" />

> 2.跨主机通信，k8s集群配置宿主机的物理网卡信息（容器也得通过宿主机的物理网卡通信）

![image-20230130165535148](/Users/jacktaoyang/Library/Application Support/typora-user-images/image-20230130165535148.png)

> 3.基于kubectl命令，应用yml文件创建pod资源

```markdown
kubectl create -f yml文件路径
```

