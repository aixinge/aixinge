package initialize

import (
	"aixinge/api/router"
	_ "aixinge/docs"
	"aixinge/global"
	"aixinge/middleware"
	"aixinge/web"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"net/http"
)

func Routers() *fiber.App {
	var app = fiber.New(fiber.Config{
		DisableStartupMessage: true,

		// https://github.com/goccy/go-json
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	global.LOG.Debug("register swagger handler")
	app.Get("/swagger/*", swagger.HandlerDefault)

	global.LOG.Debug("register filesystem handler")
	app.Use("/", filesystem.New(filesystem.Config{
		Root:         http.FS(web.Dist),
		Browse:       true,
		Index:        "index.html",
		NotFoundFile: "404.html",
		PathPrefix:   "/dist",
		MaxAge:       3600,
	}))

	global.LOG.Debug("register upload file handler")

	global.LOG.Debug("use middleware logger")
	app.Use(middleware.Logger())

	global.LOG.Debug("use middleware recover")
	app.Use(middleware.Recover())
	// 跨域
	global.LOG.Debug("use middleware cors")
	app.Use(middleware.Cors())

	// 获取路由组实例
	systemRouter := router.AppRouter.System

	// 注入免鉴权路由
	publicGroup := app.Group("/v1/")
	{
		systemRouter.InitBaseRouter(publicGroup) // 注册
	}

	// 注入鉴权路由
	privateGroup := app.Group("/v1/")
	privateGroup.Use(middleware.JWTAuth()).Use(middleware.RbacHandler())
	{
		/**  系统管理  */
		systemRouter.InitUserRouter(privateGroup) // 用户
		systemRouter.InitRoleRouter(privateGroup) // 角色
		systemRouter.InitMenuRouter(privateGroup) // 菜单

	}
	global.LOG.Debug("router register success")
	return app
}