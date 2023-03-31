#### 一.Docker技术

##### 1.1 容器技术

1.技术原理：基于**命名空间**实现**资源隔离**。

```markdown
命名空间包括：进程空间，网络空间，文件系统空间。
	1)进程空间(PID)：Linux内核分配的可用Pid号码范围；
	2)网络空间(PORT)：网络接口，网卡，port等;
	3)文件系统空间(SYSTEM)：/opt/...，/usr/...,  /home等文件路径
```

2.容器作用

```markdown
1)创建独立的namespace(独立的环境资源);
2)对资源进行限制；
3)直接使用宿主机的硬件配置。
```

##### 1.2 Docker (vs) 容器

​	Docker:基于Linux内核的Cgroups,Namespace等技术，对进程进行封装隔离，属于操作系统层面的虚拟化技术。（**封装隔离的进程也可称为容器**，docker就是管理容器的工具软件）

![image-20230129113158976](/Users/jacktaoyang/Library/Application Support/typora-user-images/image-20230129113158976.png)

##### 1.3 生命周期本质（两大要素）

​	1)镜像构建（**代码运行环境构建**）

​	2)基于镜像，运行容器（**在隔离的容器里运行项目代码**）

![image-20230129145520958](/Users/jacktaoyang/Library/Application Support/typora-user-images/image-20230129145520958.png)

##### 1.4 优势

```markdown
1) 基于docker镜像，彻底解决环境部署，迁移难和环境不一致问题；
2) docker镜像与宿主机公用一个内核，可实现轻量级发布&敏捷开发；
3) 基于docker镜像交付的devops部署流水线，能够实现环境一致性；
```

#### 二.Docker命令

##### 2.1 拉取镜像

```markdown
docker pull 镜像名称:TAG 
```

##### 2.2 运行镜像

```markdown
docker run -it -v 本地项目绝对路径:镜像中的绝对路径 镜像名称 bash

-it : 以终端的形式运行

-v : 绑定 本地路径 和 镜像路径（两者保持同步）
```

##### 2.3 构建项目target

```makefile
#生成target包，供docker build 使用
make build 
```

##### 2.4 构建镜像

```makefile
# 通过Dockerfile文件构建镜像
docker build -f Dockerfile -t 镜像名称 .
# 按规则重命名镜像
docker tag 镜像名称 按规则重命名(镜像名称：TAG)
```

##### 2.5 镜像操作

```makefile
# 查看所有镜像
docker images
# 删除本地镜像
docker image rm 镜像ID
```

##### 2.6 登录docker上传镜像到指定仓库

```makefile
docker login 仓库hosts(域名/ip) -u username -p pwd
docker push 镜像名称:TAG
```
