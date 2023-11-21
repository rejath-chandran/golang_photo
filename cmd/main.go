package main

import (
	"database/sql"
	
	"html/template"
	"log"
	"net/http"
	"github.com/go-chi/chi/v5"
	_"github.com/go-sql-driver/mysql"
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

	app:=Applcation{
		Temp: template,
	}

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
   


	 r.Get("/",func(w http.ResponseWriter, r *http.Request) {
		data:=PageData{
			Title: "HOME darta",
		}
		err=app.Render(w,"home.html",data)
		if err!=nil {
			log.Println(err)
		}

	 })

	 err = http.ListenAndServe(":9000", r)
	 if err!= nil {
		 log.Println(err)
	 }

}
