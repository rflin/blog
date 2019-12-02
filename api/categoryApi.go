package api

import "blog/models"
import "github.com/gin-gonic/gin"
import "strconv"

func PostCategoryApi(context *gin.Context){
	var _c models.Category
	context.Bind(&_c)

	if _c.Name != ""{
		err := models.PostCategory(&_c)
		if err != nil{
			context.JSON(404, gin.H{"msg": err.Error()})
		}else{
			context.JSON(201, _c)
		}
	}else{
		content := gin.H{"msg": "empty field"}
		context.JSON(406, content)
	}
}

func GetCategoriesApi(context *gin.Context){
	var c_list []models.Category
	err := models.GetCategories(&c_list)
	if err != nil{
		context.JSON(404, gin.H{"msg": err.Error()})
	}else{
		context.JSON(200, c_list)
	}
}

func GetCategoryApi(context *gin.Context){
	id, _ := strconv.Atoi(context.Param("id"))

	var _c models.Category
	err := models.GetCategory(id, &_c)

	if err != nil{
		context.JSON(404, gin.H{"msg": err.Error()})
	}else{
		context.JSON(200, _c)
	}
}

func UpdateCategoryApi(context *gin.Context){
	id, _ := strconv.Atoi(context.Param("id"))

	var _c models.Category
	context.Bind(&_c)

	err := models.UpdateCategory(id, &_c)

	if err != nil{
		context.JSON(404, gin.H{"msg": err.Error()})
	}else{
		context.JSON(200, _c)
	}
}

func DeleteCategoryApi(context *gin.Context){
	id, _ := strconv.Atoi(context.Param("id"))

	err := models.DeleteCategory(id)
	if err != nil{
		context.JSON(404, gin.H{"msg": err.Error()})
	}else{
		content := gin.H{"msg": "delete OK"}
		context.JSON(200, content)
	}
}
