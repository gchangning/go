package users

import (
	"context"
	"star/internal/dao"
)

func (u Users) UserDel(ctx context.Context, id int) error {
	_, err := dao.Users.Ctx(ctx).Where("id", id).Delete()
	if err != nil {
		return err
	}
	return nil
}
