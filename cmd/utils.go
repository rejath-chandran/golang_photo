package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func TempParse() (map[string]*template.Template,error) {

	cache:=map[string]*template.Template{}

	pages, err := filepath.Glob("/home/ioss/Desktop/project/golang_photo/temp/*.html")

	if err != nil {
		return nil, err
	}
	fmt.Println(pages)
    include_pages,err:=filepath.Glob("/home/ioss/Desktop/project/golang_photo/temp/include/*.html")
	if err != nil {
		return nil, err
	}
	base_files:=[]string{}
	for _,inc_page:=range include_pages{
		base_files=append(base_files,inc_page)
	}
	
	for _, page := range pages {
		name := filepath.Base(page)
		files:=[]string{}
		files=append(base_files,page)
		fmt.Println(files)
		ts, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}
		fmt.Println("file",files)
		cache[name]=ts
		fmt.Println("cache",cache)

	}
	fmt.Println(cache)
	return cache, nil
}


func (a *Applcation)Render(w http.ResponseWriter,page string,data interface{}) error{
	t:=a.Temp[page]
	err:=t.ExecuteTemplate(w,"base",data)
	if err!=nil{
	return err
	}
	return nil
}