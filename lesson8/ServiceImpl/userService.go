package ServiceImpl

import (
	"context"
	userService "micro/helloworld2/lesson8/Services"
	"strconv"
)

type UserServiceImpl struct {
}

func (impl *UserServiceImpl) GetUserList(ctx context.Context, req *userService.GetUserListReq, resp *userService.GetUserListResp) error {
	var i int64
	for i = 0; i < req.Size; i++ {
		resp.Data = append(resp.Data, &userService.UserModel{Id: i, Name: "userName_" + strconv.Itoa(int(i))})
	}
	return nil
}
