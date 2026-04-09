package users

import (
	"context"
	"fmt"

	"star/api/users/v1"
)

func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {

	fmt.Println(*req)
	err = c.users.Register(ctx, req.Username, req.Password, req.Email)
	return nil, err
}
