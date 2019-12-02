package models

import "blog/database"

func init(){
	db := database.InitDb()
	defer db.Close()

	if !db.HasTable(&User{}) {
		db.CreateTable(&User{})
	}
}

func PostUser(_u *User)error{
	db := database.InitDb()
	defer db.Close()

	err := db.Create(_u).Error
	return err
}

func GetUsers(u_list *[]User)error{
	db := database.InitDb()
	defer db.Close()

	// err := db.Preload("Articles").Find(u_list).Error
	err := db.Find(u_list).Error
	return err
}

func GetUser(id int, _u *User)error{
	db := database.InitDb()
	defer db.Close()

	err := db.First(_u, id).Error
	return err
}

func UpdateUser(id int, new__u *User)error{
	db := database.InitDb()
	defer db.Close()

	var _u User
	err := db.First(&_u, id).Error

	if err == nil{
		_u.Name = new__u.Name
		new__u.Id = _u.Id
		err = db.Save(_u).Error
		return err
	}
	return err
}

func DeleteUser(id int)error{
	db := database.InitDb()
	defer db.Close()

	var _u User
	err := db.First(&_u, id).Error

	if err == nil{
		err = db.Delete(&_u).Error
		return err
	}
	return err
}