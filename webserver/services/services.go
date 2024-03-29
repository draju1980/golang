package services

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"webserver/models"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

var dbconn *sqlx.DB

func SetDb(db *sqlx.DB){
	dbconn = db
}

func GetAllPosts(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var posts = models.GetPosts()
	sqlStmt:= 'SELECT * FROM posts'
	rows, err := dbconn.Queryx(sqlStmt)
	if err==nil{
		var tempPost = models.GetPost()
		for rows.Next(){
			err = rows.StructScan(&tempPost)
			posts=append(posts,tempPost)
		}
		switch err {
		case spl.ErrNoRows:
			{
				log.Println("No rows Returned")
				http.Error(w,err.Error(),204)
			}
		case nil:
			json.NewEncoder(w).Encode(&posts)
		default:
			http.Error(w,err.Error(), 400)
			return
		}

	}else{
		http.Error(w,err.Error(),400)
		return
	}
}