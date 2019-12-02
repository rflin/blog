package api

import "blog/models"
import "github.com/gin-gonic/gin"
import "strconv"

func PostArticleApi(context *gin.Context){
	var _a models.Article
	context.Bind(&_a)

	if _a.Title != "" && _a.Body != ""{
		err := models.PostArticle(&_a)
		if err != nil{
			context.JSON(404, gin.H{"msg": err.Error()})
		}else{
			context.JSON(201, _a)
		}
	}else{
		content := gin.H{"msg": "empty field"}
		context.JSON(406, content)
	}
}

func GetArticlesApi(context *gin.Context){
	var a_list []models.Article
	err := models.GetArticles(&a_list)
	if err != nil{
		context.JSON(404, gin.H{"msg": err.Error()})
	}else{
		context.JSON(200, a_list)
	}
}

func GetArticleApi(context *gin.Context){
	id, _ := strconv.Atoi(context.Param("id"))

	var _a models.Article
	err := models.GetArticle(id, &_a)

	if err != nil{
		context.JSON(404, gin.H{"msg": err.Error()})
	}else{
		context.JSON(200, _a)
	}
}

func UpdateArticleApi(context *gin.Context){
	id, _ := strconv.Atoi(context.Param("id"))

	var _a models.Article
	context.Bind(&_a)

	err := models.UpdateArticle(id, &_a)

	if err != nil{
		context.JSON(404, gin.H{"msg": err.Error()})
	}else{
		context.JSON(200, _a)
	}
}

func DeleteArticleApi(context *gin.Context){
	id, _ := strconv.Atoi(context.Param("id"))

	err := models.DeleteArticle(id)
	if err != nil{
		context.JSON(404, gin.H{"msg": err.Error()})
	}else{
		content := gin.H{"msg": "delete OK"}
		context.JSON(200, content)
	}
}
