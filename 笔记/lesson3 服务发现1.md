## 服务发现
新建一个`serviceInfo.go`文件，编写服务发现的代码。

## 步骤
1. 首先需要连接consul

    `	reg := consul.NewRegistry(registry.Addrs("20.124.120.120:8500"))`

2. 之后获取service

    `myService, err := reg.GetService("user")`

3. 接下来随机选取一个node
```golang
    next := selector.Random(myService)
	node, err := next()
```
   selector.Random方法返回的是Next，实际上是`func() (*registry.Node, error)`，可以直接调用获取node信息。

实质上是丛Node列表中使用rand.Int()来随机选取一个node返回，如果为0，则返回ErrNoneAvailable的error。
```golang
func Random(services []*registry.Service) Next {
	nodes := make([]*registry.Node, 0, len(services))

	for _, service := range services {
		nodes = append(nodes, service.Nodes...)
	}

	return func() (*registry.Node, error) {
		if len(nodes) == 0 {
			return nil, ErrNoneAvailable
		}

		i := rand.Int() % len(nodes)
		return nodes[i], nil
	}
}
```
4. 最后打印node信息

## 输出
![img_5.png](img_5.png)