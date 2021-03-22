
 # 编写

```
syntax = "proto2";

package cootek.pgd.feeds;

//必须把当前项目名放进去，不然生成的*.pb.go找不到依赖包，假设当前是chubao.
option go_package = "chubao/cootek/pgd/feeds";

import "elete.proto";
import "account.proto";

message LOCATION {
    optional string latitude = 1;
    optional string longitude = 2;
}
```
 
 
 # 编译
 
 ## 命令
 ```
 protoc -I ../cootek.com/elete/sdk -I ./protobuf --go_out=plugins=grpc:../../src ./protobuf/feeds.proto 
 ```
 当前执行命令编译protobuf文件时，因为是相对路径，所以你现在要在项目的根目录$GOPATH/src/chubao，也就是当前文件目录的上级执行这个命令。

**命令说明**

1. -I ../cootek.com/elete/sdk  
指定依赖的elete.proto路径

2. -I ./protobuf  
指定要编译的proto文件的位置

3. --go_out=plugins=grpc:../../src  
plugins=grpc指定grpc
../../src 因为上面指定的包的路径为chubao/cootek/pgd/feeds，这是相对$GOPATH/src的。

4. ./protobuf/feeds.proto 
编译指定的文件

## 扩展
为了支持增加自定义tag，比如ad_platform.proto文件，引用了第三方扩展
> https://github.com/favadi/protoc-go-inject-tag

执行完protoc命令后要记得执行下扩展命令生成自定义tag
```bash
 protoc-go-inject-tag -input=./cootek/pgd/ad_platform/ad_platform.pb.go 
```

 
# 注意
## jsonpb问题
1. 为什么不用json进行格式化  
因为针对enum类型，不支持转换为string，所以用jsonpb

2. 使用pb格式化工具jsonpb会产生的问题
- 不支持json tag为 - 表示删除. 
  例如feeds.proto中的status字段，因为没用到，解决方法就是把字段直接在proto中注释掉了
- 会把int64字段转换为字符串   
   解决:改为int32


## 针对token proto的编译
- 需要使用protoc3, 请自行安装protoc3
``` protoc -I ../cootek.com/elete/sdk -I ./protobuf  --go_out=plugins=grpc:../../src/ ./protobuf/token_service.proto ```
