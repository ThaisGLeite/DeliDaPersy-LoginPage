package routes

import (
	"log"

	"github.com/gin-gonic/gin"

	controllers "loginpage/controllers"
)

func PublicRoutes(g *gin.RouterGroup) {
	log.Println("entrou em public routes")
	g.GET("/login", controllers.LoginGetHandler())
	g.POST("/login", controllers.LoginPostHandler())
	g.GET("/", controllers.IndexGetHandler())

}

// Rotas q so pode usar depois do logon
func PrivateRoutes(g *gin.RouterGroup) {
	log.Println("entrou em private routes")
	g.GET("/dashboard", controllers.DashboardGetHandler())
	g.GET("/logout", controllers.LogoutGetHandler())
	g.GET("/cadastro", controllers.CadastroGetHandler())
	g.GET("/signin", controllers.SigninGetHandler())

}
