package routes

import (
	"github.com/dally469/api/packages/controllers"
	"github.com/dally469/api/packages/utils"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/")
	v1.OPTIONS("create-books", utils.CorsReply)
	v1.OPTIONS("getAllBooks", utils.CorsReply)
	v1.OPTIONS("getBoodById/:id", utils.CorsReply)
	v1.OPTIONS("manipulate-author", utils.CorsReply)


	v1.POST("create-books", controllers.CreateBook)
	v1.GET("getAllBooks", controllers.GetAllBook)
	v1.GET("getBoodById/:id", controllers.GetBoodById)
	v1.POST("manipulate-author", controllers.ManipulateAuthor)


	return r

}
