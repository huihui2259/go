1.在运行项目之前，需要修改conf/conf.go，将ip和端口修改为自己的ip和port
2.cd goDemo; go build ./main.go; ./main

模块及作用
conf: 配置模块
controller: 类似于java中的controller，接收前端数据并调用service接口
entity: 各结构体定义
mysql,redis: 仅连接mysql与redis，返回连接后的连接
router: 路由的分组和注册
service: 主要的逻辑处理模块
utils: 通用工具模块，包括通用错误定义处理等
.yaml: 上云的文件，使用`kubectl create -f xxx.yaml`，注意：若将服务上云的话，需要更改conf，将本地的ip变为真正ip
Dockerfile: 构建镜像文件，使用`docker build -t`
