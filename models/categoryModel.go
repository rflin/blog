package models

import "blog/database"

func init(){
	db := database.InitDb()
	defer db.Close()

	if !db.HasTable(&Category{}) {
		db.CreateTable(&Category{})
	}
}

func PostCategory(_c *Category)error{
	db := database.InitDb()
	defer db.Close()

	err := db.Create(_c).Error
	return err
}

func GetCategories(c_list *[]Category)error{
	db := database.InitDb()
	defer db.Close()

	err := db.Find(c_list).Error
	return err
}

func GetCategory(id int, _c *Category)error{
	db := database.InitDb()
	defer db.Close()

	_c.Id = id
	var articles []Article
	err := db.Model(_c).Association("Articles").Find(&articles).Error
	_c.Articles = articles
	return err
}

func UpdateCategory(id int, new__c *Category)error{
	db := database.InitDb()
	defer db.Close()

	var _c Category
	err := db.First(&_c, id).Error

	if err == nil{
		_c.Name = new__c.Name
		new__c.Id = _c.Id
		err = db.Save(_c).Error
		return err
	}
	return err
}

func DeleteCategory(id int)error{
	db := database.InitDb()
	defer db.Close()

	var _c Category
	err := db.First(&_c, id).Error

	if err == nil{
		err = db.Delete(&_c).Error
		return err
	}
	return err
}