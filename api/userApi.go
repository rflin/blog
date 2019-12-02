package api

import "blog/models"
import "github.com/gin-gonic/gin"
import "strconv"


func PostUserApi(context *gin.Context){
	var _u models.User
	context.Bind(&_u)

	if _u.Name != ""{
		err := models.PostUser(&_u)
		if err != nil{
			context.JSON(404, gin.H{"msg": err.Error()})
		}else{
			context.JSON(201, _u)
		}
	}else{
		content := gin.H{"msg": "empty field"}
		context.JSON(406, content)
	}
}

func GetUsersApi(context *gin.Context){
	var u_list []models.User
	err := models.GetUsers(&u_list)
	if err != nil{
		context.JSON(404, gin.H{"msg": err.Error()})
	}else{
		context.JSON(200, u_list)
	}
}

func GetUserApi(context *gin.Context){
	id, _ := strconv.Atoi(context.Param("id"))

	var _u models.User
	err := models.GetUser(id, &_u)

	if err != nil{
		context.JSON(404, gin.H{"msg": err.Error()})
	}else{
		context.JSON(200, _u)
	}
}

func UpdateUserApi(context *gin.Context){
	id, _ := strconv.Atoi(context.Param("id"))

	var _u models.User
	context.Bind(&_u)

	err := models.UpdateUser(id, &_u)

	if err != nil{
		context.JSON(404, gin.H{"msg": err.Error()})
	}else{
		context.JSON(200, _u)
	}
}

func DeleteUserApi(context *gin.Context){
	id, _ := strconv.Atoi(context.Param("id"))

	err := models.DeleteUser(id)
	if err != nil{
		context.JSON(404, gin.H{"msg": err.Error()})
	}else{
		content := gin.H{"msg": "delete OK"}
		context.JSON(200, content)
	}
}
