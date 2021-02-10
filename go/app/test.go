package app

import (
	"database/sql"
	"html/template"
	// "fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	// "./app"
	// "./common"
)

func Test(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:@/test")
	if err != nil {
	  panic(err.Error())
	}
	defer db.Close()
  
	rows, err := db.Query("SELECT test_id, test_txt FROM test1")
	if err != nil {
	  panic(err.Error())
	}
	testId := ""
	testTxt := ""
	for rows.Next() {
		if err := rows.Scan(&testId,&testTxt); err != nil {
			log.Print(err)
		}
	}
	
	tpl := template.Must(template.ParseFiles("test.html"))
	type View struct {
		Test              string
		Txt              string
	}
	var view View
	view.Test = testId
	view.Txt = testTxt
	tpl.Execute(w, view)
}