package users

import (
	"context"
	"errors"
	"star/internal/dao"
	"star/internal/model/entity"
)

func (u *Users) Login(ctx context.Context, username, password string) (*entity.Users, string, error) {
	var user *entity.Users
	err := dao.Users.Ctx(ctx).Where("username", username).Scan(&user)
	if err != nil {
		return nil, "", err
	}

	if user == nil {
		return nil, "", errors.New("用户不存在")
	}

	if user.Password != password {
		return nil, "", errors.New("密码错误")
	}

	token, err := u.tokenManager.GenerateToken(ctx, uint(user.Id), user.Username, user.Email)
	if err != nil {
		return nil, "", errors.New("生成 Token 失败")
	}

	if err := u.tokenManager.SaveTokenToFile(ctx, username, token); err != nil {
		return nil, "", errors.New("保存 Token 失败")
	}

	return user, token, nil
}
