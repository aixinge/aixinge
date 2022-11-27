package message

import (
	v1 "aixinge/api/v1"
	"github.com/gofiber/fiber/v2"
)

type MailTemplateRouter struct {
}

func (s *MailTemplateRouter) InitMailTemplateRouter(router fiber.Router) (R fiber.Router) {
	mtRouter := router.Group("mail-template")
	var mtApi = v1.AppApi.MessageApi.MailTemplate
	{
		mtRouter.Post("create", mtApi.Create) // 创建
		mtRouter.Post("get", mtApi.Get)       // 根据id获取邮件模板
	}
	return mtRouter
}