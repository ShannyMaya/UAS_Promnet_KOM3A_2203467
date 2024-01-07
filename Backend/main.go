package main

import (
	"net/http"
	"github.com/jeypc/go-crud/controllers/itemcontroller"
)
func main(){

	http.HandleFunc("/", itemcontroller.Index)
	http.HandleFunc("/item", itemcontroller.Index)
	http.HandleFunc("/item/index", itemcontroller.Index)
	http.HandleFunc("/item/add", itemcontroller.Add)
	http.HandleFunc("/item/edit", itemcontroller.Edit)
	http.HandleFunc("/item/delete", itemcontroller.Delete)

	http.ListenAndServe(":3000", nil)
	
}
