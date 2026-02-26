// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AgentActivityInfoDao is the data access object for the table t_agent_activity_info.
type AgentActivityInfoDao struct {
	table    string                   // table is the underlying table name of the DAO.
	group    string                   // group is the database configuration group name of the current DAO.
	columns  AgentActivityInfoColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler       // handlers for customized model modification.
}

// AgentActivityInfoColumns defines and stores column names for the table t_agent_activity_info.
type AgentActivityInfoColumns struct {
	Id         string // id
	Type       string // 1 充值返利 2挖矿返利
	Amount     string // 返利金额
	CoinType   string // 币种
	FromId     string // 返利用户
	UserId     string // 用户id
	CreateBy   string //
	CreateTime string // 创建时间
	UpdateBy   string //
	UpdateTime string // 更新时间
	Status     string // 1  待返  2  已返
	LoginName  string //
	SerialId   string //
}

// agentActivityInfoColumns holds the columns for the table t_agent_activity_info.
var agentActivityInfoColumns = AgentActivityInfoColumns{
	Id:         "id",
	Type:       "type",
	Amount:     "amount",
	CoinType:   "coin_type",
	FromId:     "from_id",
	UserId:     "user_id",
	CreateBy:   "create_by",
	CreateTime: "create_time",
	UpdateBy:   "update_by",
	UpdateTime: "update_time",
	Status:     "status",
	LoginName:  "login_name",
	SerialId:   "serial_id",
}

// NewAgentActivityInfoDao creates and returns a new DAO object for table data access.
func NewAgentActivityInfoDao(handlers ...gdb.ModelHandler) *AgentActivityInfoDao {
	return &AgentActivityInfoDao{
		group:    "default",
		table:    "t_agent_activity_info",
		columns:  agentActivityInfoColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AgentActivityInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AgentActivityInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AgentActivityInfoDao) Columns() AgentActivityInfoColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AgentActivityInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AgentActivityInfoDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AgentActivityInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
