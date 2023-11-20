package main

import (
	"fmt"
	"html/template"
	"path/filepath"
)

func TempParse() (map[string]*template.Template,error) {

	cache:=map[string]*template.Template{}

	pages, err := filepath.Glob("../photo_share/temp/*.html")

	if err != nil {
		return nil, err
	}

    include_pages,err:=filepath.Glob("../photo_share/temp/include/*.html")
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

	return cache, nil
}
