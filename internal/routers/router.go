package routers

import (
	_ "Pro/blog-service/docs"
	"Pro/blog-service/internal/middleware"
	"Pro/blog-service/internal/routers/api"
	v1 "Pro/blog-service/internal/routers/api/v1"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	article := v1.NewArticle()
	tag := v1.NewTag()
	r.Use(middleware.Translations())

	//http://127.0.0.1:8000/swagger/index.html
	docUrl := ginSwagger.URL("http://127.0.0.1:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, docUrl))

	//FIXME 应该写为api.NewUpload
	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("tags", tag.List)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}
	return r
}
