package main

import(
	"blog/api"
	"github.com/gin-gonic/gin"
	// "fmt"
)

func main(){
	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		v1.POST("/user", api.PostUserApi)
		v1.GET("/users", api.GetUsersApi)
		v1.GET("/user/:id", api.GetUserApi)
		v1.PUT("/user/:id", api.UpdateUserApi)
		v1.DELETE("/user/:id", api.DeleteUserApi)

		v1.POST("/article", api.PostArticleApi)
		v1.GET("/articles", api.GetArticlesApi)
		v1.GET("/article/:id", api.GetArticleApi)
		v1.PUT("/article/:id", api.UpdateArticleApi)
		v1.DELETE("/article/:id", api.DeleteArticleApi)


		v1.POST("/category", api.PostCategoryApi)
		v1.GET("/categories", api.GetCategoriesApi)
		v1.GET("/category/:id", api.GetCategoryApi)
		v1.PUT("/category/:id", api.UpdateCategoryApi)
		v1.DELETE("/category/:id", api.DeleteCategoryApi)
	}
	r.Run(":8080")
}