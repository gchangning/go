package users

import (
	"context"
	"star/api/users/v1"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	user, token, err := c.users.Login(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	return &v1.LoginRes{
		Id:       uint(user.Id),
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}, nil
}

func (c *ControllerV1) DelUser(ctx context.Context, req *v1.DelUserReq) (res *v1.DelUserRes, err error) {
	claims, err := c.users.GetTokenManager().ParseToken(ctx, req.Token)
	if err != nil {
		return &v1.DelUserRes{
			Success: false,
			Message: "Token 验证失败: " + err.Error(),
		}, nil
	}

	if claims.Username == "" {
		return &v1.DelUserRes{
			Success: false,
			Message: "无效的 Token",
		}, nil
	}

	err = c.users.UserDel(ctx, req.Id)
	if err != nil {
		return &v1.DelUserRes{
			Success: false,
			Message: "删除用户失败: " + err.Error(),
		}, nil
	}

	return &v1.DelUserRes{
		Success: true,
		Message: "用户删除成功",
	}, nil
}
func (c *ControllerV1) GetUserWithInfo(ctx context.Context, req *v1.GetUserWithInfoReq) (res *v1.GetUserWithInfoRes, err error) {
	result, err := c.users.GetUserWithInfo(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	return &v1.GetUserWithInfoRes{
		User:     result.User,
		UserInfo: result.UserInfo,
	}, nil
}
