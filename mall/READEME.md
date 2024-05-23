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


#### RPC调用metadata

## go-zero添加client拦截器

order服务的search接口中添加拦截器，添加一些requestID token userID等拦截器

# 关键点：
1. 存入metadata时机：             searchLogic.go里面，在调用RPC方法之前
2. 如何存:                       通过自定义拦截器拦截需要传入的metadata
3. 拦截其中如何通过context传值     在svc中初始化RPC服务的时候传入拦截器 
4. context存值取值操作            自定义类型传入key并设置好value（防止协作开发他人写相同key（数据冲突）
  

## go-zero添加server拦截器

# 关键点:
1. 什么时候添加拦截器： 在user.go中：服务启动之前调用(s.Start()之前)，千万勿忘添加自定义拦截器
2. 拦截器业务怎么写： 点开user.go中的s.AddUnaryInterceptors，找到对应函数，传入传出值复制到自定义拦截器上
3. 服务器拦截器怎么取值:  在handler(实际RPC方法)调用之前或者之后进行业务逻辑处理
　　　　　　　　　　　　　其中:取出元数据： metadata.FromIncomingContext(ctx)


#### 错误处理

1. 自定义错误格式

2. 业务代码中要按照需求返回自定义错误（用起来）

3. 处理：告诉go-zero框架，处理一下自定义错误（重要）
   在启动客户端之前处理掉这些错误,go-zero框架中有 httpx包: httpx.SetErrorHandlerCtx() ，拿来主义直接用

   #### go-zero中的goctl模板(暂时只了解
)

   模板的用处: 用来生成代码，goctl指令生成代码时使用（类似于:）goctl api go -api user.api -dir . -style=goZero

   ### goctl template用法:

   goctl env 查看版本信息： 我的：OCTL_HOME=C:\Users\shone\.goctl

   初始化模板: goctl template init