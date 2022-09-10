`protoc --micro_out=../ --go_out=../ *.proto `

**特别坑的一个点！**

这里rpc服务名要大写！！！不然找不到

```protobuf
service UserService{
  rpc GetUserList(getUserListReq) returns (getUserListResp);  //RPC服务名要大写！
}
```