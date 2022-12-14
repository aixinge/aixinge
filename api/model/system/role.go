package system

import (
	"aixinge/global"
	"aixinge/utils/snowflake"
)

type Role struct {
	global.MODEL
	Name   string `json:"name"`   // 名称
	Alias  string `json:"alias"`  // 别名
	Remark string `json:"remark"` // 备注
	Status int    `json:"status"` // 状态，1、正常 2、禁用
	Sort   int    `json:"sort"`   // 排序
}

type RoleMenu struct {
	RoleId snowflake.ID `json:"roleId" swaggertype:"string"` // 角色ID
	MenuId snowflake.ID `json:"menuId" swaggertype:"string"` // 菜单ID
}
