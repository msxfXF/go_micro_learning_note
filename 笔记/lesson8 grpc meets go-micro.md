`protoc --micro_out=../ --go_out=../ *.proto `

**特别坑的一个点！**

这里rpc服务名要大写！！！不然找不到

```protobuf
service UserService{
  rpc GetUserList(getUserListReq) returns (getUserListResp);  //RPC服务名要大写！
}
```

运行完生成命令后，需要实现micro文件的接口即可
```go
func (impl *UserServiceImpl) GetUserList(ctx context.Context, req *userService.GetUserListReq, resp *userService.GetUserListResp) error {
	var i int64
	for i = 0; i < req.Size; i++ {
		resp.Data = append(resp.Data, &userService.UserModel{Id: i, Name: "userName_" + strconv.Itoa(int(i))})
	}
	return nil
}
```

同时，注册服务时使用micro生成的service，，需要传入ServiceImpl.UserServiceImpl实现类的指针。
```go
	func main() {
	csRegistry := consul.NewRegistry(registry.Addrs("20.124.120.120:8500"))
	mService := micro.NewService(
		micro.Address(":8003"),
		micro.Name("Grpc"),
		micro.Registry(csRegistry),
	)
	mService.Init()
	err := userService.RegisterUserServiceHandler(mService.Server(), new(ServiceImpl.UserServiceImpl))
	if err != nil {
		log.Fatal(err)
	}
	err = mService.Run()
	if err != nil {
		log.Fatal(err)
	}

}
```


测试时同样需要连接到服务发现consul，userService.NewUserService()实例化对象，并调用方法。
```go
csRegistry := consul.NewRegistry(registry.Addrs("20.124.120.120:8500"))
	gService := micro.NewService()
	gService.Init(micro.Registry(csRegistry))
	uService := userService.NewUserService("Grpc", gService.Client())
	list, err := uService.GetUserList(context.Background(), &userService.GetUserListReq{Size: 5})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(list)
  ```
  
