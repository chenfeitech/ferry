package gorm

import (
	"ferry/models/process"
	"ferry/models/system"
	"ferry/models/wxquestion"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	// db.SingularTable(true)
	return db.AutoMigrate(
		// 系统管理
		&system.CasbinRule{},
		&system.Dept{},
		&system.Menu{},
		&system.LoginLog{},
		&system.RoleMenu{},
		&system.SysRoleDept{},
		&system.SysUser{},
		&system.SysRole{},
		&system.Post{},
		&system.Settings{},

		// 流程中心
		&process.Classify{},
		&process.TplInfo{},
		&process.TplData{},
		&process.WorkOrderInfo{},
		&process.TaskInfo{},
		&process.Info{},
		&process.History{},
		&process.CirculationHistory{},

		// 微信小程序
		&wxquestion.Rank{},
	)
}
