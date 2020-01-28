package main

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/dbohachev/beego-swagger/example/beego-api/docs" // Docs, generated by the tool
	_ "github.com/dbohachev/beego-swagger/example/beego-api/routers"

	"github.com/astaxie/beego"
)

// @title Swagger Example API
// @version 1.0
// @description This is simple beego API application server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth


func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
	}

	// In some cases, the below is needed. Enables CORS for your application.
	// But fix it accordingly, if you'll ever use it in production.
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	// Example of basic auth verification. To see it, use Authorize button for your UI.
	// If you don't see auth. being sent, make sure you enabled CORS (lines 39-45),
	// or if you properly set authorization requirements for your API (@Security BasicAuth in example).
	// Finally, Basic Auth is not reliable, specifically for cases when http is on,
	// not https.
	beego.InsertFilter("*", beego.BeforeRouter, func(context *context.Context){
		uname, pwd, _ := context.Request.BasicAuth()

		logs.Debug("Basic auth parameters are: %s %s", uname, pwd)
		// TODO: handle redirect if required
	})

	// See routers folder, routers.go for app routes + swagger routes registration.

	beego.Run()
}
