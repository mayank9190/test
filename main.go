package main

import (
	"database/sql"
	"fmt"

	_"github.com/go-sql-driver/mysql"
)
type User struct{
	rollno string `json:"rollno"`
	name string `json:"name"`
	class string `json:"class"`
	Name string `json:"rollno"`
}

func main(){
	fmt.Println("Go mysql sever")
	db, err := sql.Open("mysql","root:mayank99@tcp(127.0.0.1:3306)/test")
	if err!=nil{
		panic(err.Error())
	}
	defer db.Close()

   insert, err:= db.Query("INSERT INTO teststudent VALUES('102','amittt','10','f')")
   if err!=nil{
	   panic(err.Error())
   }
   defer insert.Close()
   fmt.Println("Success")

   result, err := db.Query("SELECT name FROM teststudent")
   if err!=nil{
  panic(err.Error())
   }
   for result.Next(){
	   var user User
	   err= result.Scan(&user.name)
	   if err!=nil{
		   panic(err.Error())
	   }
	   fmt.Println(user.name)
   }




}