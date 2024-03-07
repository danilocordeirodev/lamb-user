package db

import (
	"fmt"

	"github.com/danilocordeirodev/lamb-user/models"
	"github.com/danilocordeirodev/lamb-user/tools"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Initialize signup")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	insertSQL := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('"+sig.UserEmail+"','"+sig.UserUUID+"', '"+tools.DateMySQL()+"')"
	fmt.Println(insertSQL)

	_, err = Db.Exec(insertSQL)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Successfull signup")

	return nil
}