package routers

import (
	_ "Pro/blog-service/docs"
	"Pro/blog-service/internal/middleware"
	v1 "Pro/blog-service/internal/routers/api/v1"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {

	r := gin.Default()
	r.Use(middleware.Translations())
	//http://127.0.0.1:8000/swagger/index.html
	docUrl := ginSwagger.URL("http://127.0.0.1:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, docUrl))
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/tags", v1.Tag{}.Create)
		apiv1.DELETE("/tags/:id", v1.Tag{}.Delete)
		apiv1.PUT("/tags/:id", v1.Tag{}.Update)
		apiv1.PATCH("/tags/:id/state", v1.Tag{}.Update)
		apiv1.GET("tags", v1.Tag{}.List)

		apiv1.POST("/articles", v1.Article{}.Create)
		apiv1.DELETE("/articles/:id", v1.Article{}.Delete)
		apiv1.PUT("/articles/:id", v1.Article{}.Update)
		apiv1.PATCH("/articles/:id/state", v1.Article{}.Update)
		apiv1.GET("/articles/:id", v1.Article{}.Get)
		apiv1.GET("/articles", v1.Article{}.List)
	}
	return r
}
