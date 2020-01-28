package routers

import (
	"github.com/astaxie/beego"
	beeSwagger "github.com/dbohachev/beego-swagger"
	"github.com/dbohachev/beego-swagger/example/beego-api/controllers"
)

func init() {
	beego.Get("/swagger/*", beeSwagger.Handler(beeSwagger.URL("http://localhost:8080/swagger/doc.json")))

	// To show how can we write custom "GET" methods:
	beego.Router("/api/v1/users/:id", &controllers.UsersController{}, "get:GetByID")
	// Maps to standard beego RESTful controller. The "id" part is required for DELETE call
	beego.Router("/api/v1/users/?:id", &controllers.UsersController{})
}
