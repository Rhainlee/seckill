# seckill
参考《Go语言高并发与微服务实战》，搭建一个Golang秒杀项目

### 简介
移动端和前端应用通过网关与后端服务进行交互，进行网络请求。接入系统包括用户鉴权、负载均衡以及限流和熔断器，这是每个请求处理都需要的基础功能组件。后端核心逻辑有用户登录、秒杀处理、秒杀活动管理和系统降级等，这些服务都注册到服务注册中心，并通过配置中心进行自身业务数据的配置。链路监控时刻监控着系统的状态。最底层是缓存层的Redis以及持久化层MySQL和Zookeeper。

### 依赖基础组件
- redis
- zookeeper
- git仓库
- consul

#### 部署
- 1 部署 consul
- 2 部署 Redis,Zookeeper,MySQL。
  安装完MySQL后，可以导入主目录下的seckill.sql
- 3 新建git repo
  可以参考 https://gitee.com/cloud-source/config-repo 创建对应项目的文件，修改Redis，MySQL，Zookeeper等组件的配置
- 4 部署 Config-Service
  在yml文件中配置对应的git项目地址和consul地址，构建并运行Java程序，将config-service注册到consul上
- 5 修改bootstrap文件
  修改各个项目中的bootstrap.yml文件discover相关的consul地址和config-service的相关配置
