// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UsersInfo is the golang structure for table users_info.
type UsersInfo struct {
	Id        int         `json:"id"        orm:"id"         description:""` //
	Uid       int         `json:"uid"       orm:"uid"        description:""` //
	Nickname  string      `json:"nickname"  orm:"nickname"   description:""` //
	Avatar    string      `json:"avatar"    orm:"avatar"     description:""` //
	Phone     string      `json:"phone"     orm:"phone"      description:""` //
	Address   string      `json:"address"   orm:"address"    description:""` //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""` //
}
