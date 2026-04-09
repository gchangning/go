package users

import (
	"context"
)

func (u *Users) Logout(ctx context.Context, username string) error {
	return u.tokenManager.RemoveTokenFile(ctx, username)
}
