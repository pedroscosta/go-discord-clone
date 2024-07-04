package repositories

import (
	"errors"
	db_ "go-discord-clone/configs"
	"go-discord-clone/models"
)

/*
Sensitive data isn't omitted from response because this is not accessible through the API, if the user has access to the data you should remove non-essential data (e.g. passwords) from the response.
*/

func GetUser(username string) (*models.User, error) {
	db := db_.DBConn

	var user models.User
	db.Where(&models.User{Username: username}).Find(&user)
	// db.Find(&user, "id = ?", "pedro")

	if user.Username == "" {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func GetUsers() ([]*models.User, error) {
	db := db_.DBConn

	var users []*models.User
	db.Find(&users)

	return users, nil
}

func CreateUser(user *models.User) error {
	db := db_.DBConn

	hash, err := user.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hash
	db.Create(user)

	return nil
}

func UpdateUser(user *models.User) error {
	db := db_.DBConn

	hash, err := user.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hash
	db.Save(user)

	return nil
}

func DeleteUser(username string) error {
	db := db_.DBConn

	db.Delete(&models.User{Username: username})

	return nil
}
