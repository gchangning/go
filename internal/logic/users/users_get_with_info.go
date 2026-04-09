package users

import (
	"context"
	"star/internal/dao"
	"star/internal/model/entity"
)

type UserInfoWithDetail struct {
	User     *entity.Users     `json:"user"`
	UserInfo *entity.UsersInfo `json:"userInfo"`
}

func (u *Users) GetUserWithInfo(ctx context.Context, userId uint) (*UserInfoWithDetail, error) {
	var user *entity.Users
	err := dao.Users.Ctx(ctx).Where("id", userId).Scan(&user)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	var userInfo *entity.UsersInfo
	err = dao.UsersInfo.Ctx(ctx).Where("uid", userId).Scan(&userInfo)
	if err != nil {
		return nil, err
	}

	return &UserInfoWithDetail{
		User:     user,
		UserInfo: userInfo,
	}, nil
}
