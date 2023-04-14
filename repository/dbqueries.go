package repository

import (
	types "github.com/chandanaavadhani/BusService/models"
)

func InserUser(user types.Signup) {
	db, err := Connectdb()
	if err != nil {
		panic(err)
	}
	stmt, err := db.Prepare("INSERT INTO `poojadb`.`userdetails`(fullname, username, email,password) VALUES (?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(person.Firstname+" "+person.Lastname, person.Username, person.Email, person.Password)
	if err != nil {
		return err
	}
	return nil

}
