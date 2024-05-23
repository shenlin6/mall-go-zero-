## 1. 生成api代码

# 首先切入到api目录
# 输入指令： goctl api go -api user.api -dir . -style=goZero
# 再完善api/internal/svc/serviceContext.go(桥梁)还有config 和yaml文件

## 2. 生成rpc代码

# 1. 首先切入到rpc目录
# 2. 输入指令：goctl rpc protoc user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.
# 3. 再完善 rpc/internal/svc/serviceContext.go(桥梁)  还有config 和yaml文件 
#  注意yaml文件和config.go一定要对应上
# 4. 完善 rpc的业务逻辑（logic层）


#### 启动rpcui：( localhost:8080 为rpc服务地址)
# 指令: grpcui -plaintext localhost:8080 


# 如果出现下面的情况：

# grpcui -plaintext localhost:8080
# Failed to compute set of methods to expose: server does not support the reflection API

# 需要再yaml文件中，把Mode切换到dev模式

# 指令：切换到grpc webui 可视化端口
# grpcui -plaintext localhost:8080




#### 实现 order 调用其他 rpc服务

## 使用订单ID查询到用户ID再查询用户信息

## go-zero中通过rpc调用其他服务：

# 1.配置config：添加 zrpc.RpcClientConf //连接其他微服务的RPC客户端（可以起名为： UserRPC）

# 2. 修改yaml文件：通过 RpcClientConf结构体来配置

# 3. 修改svc/serviceContext.go（修改桥梁）

# 4. 生成model层代码:
# 转到 order 目录执行: goctl model mysql datasource -url="root:wssl20050419.@tcp(127.0.0.1:3306)/db2" -table="order" -dir=./model -cache=true

# 4. 配置完毕后:编写业务逻辑


#### 使用consul作为注册中心

### 服务注册
1. 修改配置(config结构体和yaml文件)
- 在config文件引入github文件:  "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
2. 服务启动的时候将服务注册到consul:
在user.go中写入: consul.RegisterService(c.ListenOn,c.Consul)

### 服务发现
1. 修改配置（修改yaml，无需修改config）:配置Target:
2. 在order.go文件里(启动程序)匿名导入: _ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"