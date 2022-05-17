package customer

import (
	"fmt"
	"log"
	"phonenumbers_backend/database"
)

func DbGetCustomers() (users []Customer) {
	db := database.Db
	row, err := db.Query("SELECT * FROM user ORDER BY ID")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var user Customer
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&user.Id, &user.Name, &user.Phone)
		users = append(users, user)
	}
	return users
}
func DbGetCustomer(id int) (user Customer) {
	db := database.Db
	stmt, err := db.Prepare("SELECT id, name, description FROM user WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRow(id)
	if err != nil {
		log.Fatal(err)
	}
	row.Scan(&user.Id, &user.Name, &user.Phone)
	return user
}
func DbAddCustomer(user Customer) Customer {
	db := database.Db
	insert := "INSERT INTO user (name, description) values ( ?, ?) RETURNING id"
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatalln(err.Error())
	}
	id := 0
	err = stmt.QueryRow(user.Name, user.Phone).Scan(&id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	user.Id = id
	return user
}
func DbUpdateCustomer(user Customer) Customer {
	db := database.Db
	update := "UPDATE user SET name = ?, description = ? WHERE id = ?"
	stmt, err := db.Prepare(update)
	if err != nil {
		log.Fatalln(err.Error())
	}
	res, err := stmt.Exec(user.Name, user.Phone, user.Id)
	affect, err := res.RowsAffected()
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(affect)
	return DbGetCustomer(user.Id)
}
func DbDeleteCustomer(id int) Customer {
	db := database.Db
	delete := "DELETE FROM user WHERE id = ?"
	stmt, err := db.Prepare(delete)
	if err != nil {
		log.Fatalln(err.Error())
	}
	res, err := stmt.Exec(id)
	affect, err := res.RowsAffected()
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(affect)
	return DbGetCustomer(id)
}
