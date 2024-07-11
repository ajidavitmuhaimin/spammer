package main

import "fmt"
import "database/sql"
import _"modernc.org/sqlite"
import "encoding/json"
import "net/http"

type SearchResponse struct{
	Number string `json:"number"`
	Tags string `json:"tags"`
}

func main(){

	http.HandleFunc("/search",func(w http.ResponseWriter, r *http.Request){
		searchQuery:=r.URL.Query().Get("number")
		searchQuery2:=r.URL.Query().Get("tag")
		fmt.Println("get number",searchQuery,searchQuery2)
		//get query from number and tag

		db,err:=sql.Open("sqlite","database.db")
		if err!=nil{
			panic(err)
		}
		//opening database
		rows,err:=db.Query(`SELECT DISTINCT number,tag FROM tags WHERE number=?`,searchQuery)
		if err!=nil{
			panic(err)
		}
		//select number and tags
		var resp []SearchResponse
		//prepare variable resp with array SearchResponse data structure/type

		for rows.Next(){
			var r SearchResponse
			//preparing variable r for save data with SearchResponse data structure/type
			rows.Scan(&r.Number,&r.Tags)
			//scan data per row and save it to r
			resp=append(resp,r)
			//define resp with appended resp and r
			//print result
			//fmt.Println(resp)
		}
		fmt.Println(resp)
		//set http header to disable CORS and set it to json response type
		w.Header().Set("Access-Control-Allow-Origin","*")
		w.Header().Set("Content-Type","application/json")
		//w.Write([]byte(&resp))
		result,err:=json.Marshal(resp)
		if err!=nil{
			panic(err)
		}
		//convert resp array to json
		w.Write([]byte(result))
		//send to frontend
		db.Exec(`INSERT INTO tags values(?,?)`,searchQuery,searchQuery2)
		//add new tag to database every user search
	})
	http.ListenAndServe(":1395",nil)

}
//final.go
