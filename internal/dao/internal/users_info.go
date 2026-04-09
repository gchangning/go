// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UsersInfoDao is the data access object for the table users_info.
type UsersInfoDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  UsersInfoColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// UsersInfoColumns defines and stores column names for the table users_info.
type UsersInfoColumns struct {
	Id        string //
	Uid       string //
	Nickname  string //
	Avatar    string //
	Phone     string //
	Address   string //
	CreatedAt string //
	UpdatedAt string //
}

// usersInfoColumns holds the columns for the table users_info.
var usersInfoColumns = UsersInfoColumns{
	Id:        "id",
	Uid:       "uid",
	Nickname:  "nickname",
	Avatar:    "avatar",
	Phone:     "phone",
	Address:   "address",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewUsersInfoDao creates and returns a new DAO object for table data access.
func NewUsersInfoDao(handlers ...gdb.ModelHandler) *UsersInfoDao {
	return &UsersInfoDao{
		group:    "default",
		table:    "users_info",
		columns:  usersInfoColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UsersInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UsersInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UsersInfoDao) Columns() UsersInfoColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UsersInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UsersInfoDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *UsersInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
