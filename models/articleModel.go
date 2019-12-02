package models

import "blog/database"

func init(){
	db := database.InitDb()
	defer db.Close()

	if !db.HasTable(&Article{}) {
		db.CreateTable(&Article{})
	}
}

func PostArticle(_a *Article)error{
	db := database.InitDb()
	defer db.Close()

	err := db.Create(_a).Error
	return err
}

func GetArticles(a_list *[]Article)error{
	db := database.InitDb()
	defer db.Close()

	err := db.Find(a_list).Error
	return err
}

func GetArticle(id int, _a *Article)error{
	db := database.InitDb()
	defer db.Close()

	err := db.First(_a, id).Error
	return err
}

func UpdateArticle(id int, new__a *Article)error{
	db := database.InitDb()
	defer db.Close()

	var _a Article
	err := db.First(&_a, id).Error

	if err == nil{
		new__a.Id = _a.Id
		err = db.Save(new__a).Error
		return err
	}
	return err
}

func DeleteArticle(id int)error{
	db := database.InitDb()
	defer db.Close()

	var _a Article
	err := db.First(&_a, id).Error

	if err == nil{
		err = db.Delete(&_a).Error
		return err
	}
	return err
}