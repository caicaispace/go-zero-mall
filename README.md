# go-zero 微服务框架商城

### doc

- https://zhuanlan.zhihu.com/p/461604538
- https://go-zero.dev/cn/

#### goctl

```
$ GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get -u github.com/zeromicro/go-zero/tools/goctl
```

- api 生成

```
goctl api go -api ./api/search.api -dir ./api
./cmd
```

- rpc 生成

```
goctl rpc proto -src ./rpc/search.proto -dir ./rpc
```

### cmd.sh

- dockerfile  生成

```shell
./cmd.sh gen dockerfile gateway
```

- docker image build

```shell
./cmd.sh docker build gateway
```

- docker image push

```shell
./cmd.sh docker push gateway ttsimple
```

#### etcd

[docker搭建etcd服务](https://juejin.cn/post/6844904070059130894)

- 安装

```shell
# docker pull appcelerator/etcd
// 注意！ 生产环境一定不能在所有网卡无密暴露端口！
# docker run -d -p 2379:2379 -p 2380:2380 appcelerator/etcd --listen-client-urls http://0.0.0.0:2379 --advertise-client-urls http://0.0.0.0:2379
```

- 测试

```go
package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	var (
		config clientv3.Config
		client *clientv3.Client
		err error
	)

	// 创建配置对象，指定server地址并设置超时时间
	// 这里因为我用的是windows系统 docker安装在虚拟中
	// 所以地址填的是虚拟机ip
	config = clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}

	if client, err = clientv3.New(config); err != nil {
		// 只是测试一下，有错误就直接panic吧
		panic(err)
	}

	_, err = client.Put(context.TODO(), "/user/Roki", "hello! etcd")
	if err != nil {
		panic(err)
	}

	response, err := client.Get(context.TODO(), "/user/Roki")
	if err != nil {
		panic(err)
	}

	for _, kv := range response.Kvs {
		fmt.Println(string(kv.Key), string(kv.Value))
	}
}
```