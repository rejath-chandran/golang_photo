package main

import (
	"database/sql"
	"fmt"

	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
)

type PageData struct {
	Title   string
	Content string
}
type Applcation struct{
	Temp map[string]*template.Template
}
func main() {

	r := chi.NewRouter()
	dataSourceName := `root:rooPasswrd@tcp(165.232.188.131:7000)/ricky`
	
	
	template, err := TempParse()
	if err != nil {
		log.Fatal(err)
	}

	// app:=Applcation{
	// 	Temp: template,
	// }

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
   

	 r.Get("/home",func(w http.ResponseWriter, r *http.Request) {

		data:=PageData{
			Title: "hello",
		}

		t:=template["index.html"]
		fmt.Println("con",template)
		err:=t.ExecuteTemplate(w,"index.html",data)
		if err!=nil{
		log.Fatal(err)
		}

	 })

	http.ListenAndServe(":3000", r)

}
