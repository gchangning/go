package users

import (
	"context"
	"star/api/users/v1"
)

func (c *ControllerV1) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	err = c.users.Logout(ctx, req.Username)
	if err != nil {
		return &v1.LogoutRes{
			Success: false,
			Message: "登出失败",
		}, err
	}

	return &v1.LogoutRes{
		Success: true,
		Message: "登出成功",
	}, nil
}
