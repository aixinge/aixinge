package request

import (
	"aixinge/api/model/common/request"
	"aixinge/utils/snowflake"
)

type UserCreate struct {
	Username string `json:"username"` // 用户登录名
	Password string `json:"password"` // 用户登录密码
	Nickname string `json:"nickname"` // 用户昵称
}

type Login struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}

type RefreshToken struct {
	RefreshToken string `json:"refreshToken"` // 刷新票据
}

type ChangePasswordStruct struct {
	Username    string `json:"username"`    // 用户名
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

// UserRoleParams 用户分配角色参数对象
type UserRoleParams struct {
	ID      snowflake.ID   `json:"id,omitempty" swaggertype:"string"`  // 用户ID
	RoleIds []snowflake.ID `json:"roleIds" swaggertype:"array,string"` // 角色ID集合
}

type UserPageParams struct {
	request.PageInfo
	Username string `json:"username"`
	Status   int    `json:"status,string,int"`
}
