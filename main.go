package main

import "fmt"
import "database/sql"
import _"modernc.org/sqlite"

type Structure struct{
	Number string
	Tag string
}

func main(){

	db,err:=sql.Open("sqlite","database.db")
	if err!=nil{
		panic(err)
	}
	//var st Structure
	//db.QueryRow(`SELECT number,tagword FROM tag`).Scan(&st.Number,&st.Tag)
	//fmt.Println(st.Number,st.Tag)
	rows,err:=db.Query(`SELECT tagword FROM tag`)
	if err!=nil{
		panic(err)
	}
	var people []Structure
	for rows.Next(){
		var u Structure
		rows.Scan(&u.Tag)
		people=append(people,u)
		fmt.Println(people)
	}

}
