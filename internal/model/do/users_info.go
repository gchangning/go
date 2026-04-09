// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UsersInfo is the golang structure of table users_info for DAO operations like Where/Data.
type UsersInfo struct {
	g.Meta    `orm:"table:users_info, do:true"`
	Id        any         //
	Uid       any         //
	Nickname  any         //
	Avatar    any         //
	Phone     any         //
	Address   any         //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
