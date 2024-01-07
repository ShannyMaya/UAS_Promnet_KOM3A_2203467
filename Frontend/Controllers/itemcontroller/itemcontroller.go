package itemcontroller

import (
	"github.com/jeypc/go-crud/entities"
	"github.com/jeypc/go-crud/models"
	"github.com/jeypc/go-crud/libraries"
	"html/template"
	"net/http"
	"strconv"
)

var validation = libraries.NewValidation()
var itemModel = models.NewItemModel()

func Index(response http.ResponseWriter, request *http.Request) {

	item, _ := itemModel.FindAll()

	data := map[string]interface{}{
		"item": item,
	}
	
	temp, err := template.ParseFiles("views/item/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)

}


func Add(response http.ResponseWriter, request *http.Request) {
	
	if  request.Method == http.MethodGet{
		temp, err := template.ParseFiles("views/item/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost{

		request.ParseForm()

		var item entities.Item
		item.NamaBarang = request.Form.Get("nama_barang")
		item.Jumlah = request.Form.Get("jumlah")
		item.HargaSatuan = request.Form.Get("harga_satuan")
		item.Lokasi = request.Form.Get("lokasi")
		item.Deskripsi = request.Form.Get("deskripsi")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(item)

		if vErrors != nil{
			data["item"] = item
			data["validation"] = vErrors
		}else{
			data["pesan"] = "Item added succesfully"
			itemModel.Create(item)
		}

		temp, _ := template.ParseFiles("views/item/add.html")
		temp.Execute(response, data)

	}
	
}

func Edit(response http.ResponseWriter, request *http.Request) {

	if  request.Method == http.MethodGet{

		queryString := request.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)
		
		var item entities.Item
		itemModel.Find(id, &item)

		data := map[string]interface{}{
			"item": item,
		}

		temp, err := template.ParseFiles("views/item/edit.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, data)
	} else if request.Method == http.MethodPost{

		request.ParseForm()

		var item entities.Item
		item.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
		item.NamaBarang = request.Form.Get("nama_barang")
		item.Jumlah = request.Form.Get("jumlah")
		item.HargaSatuan = request.Form.Get("harga_satuan")
		item.Lokasi = request.Form.Get("lokasi")
		item.Deskripsi = request.Form.Get("deskripsi")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(item)

		if vErrors != nil{
			data["item"] = item
			data["validation"] = vErrors
		}else{
			data["pesan"] = "Item updated succesfully"
			itemModel.Update(item)
		}

		temp, _ := template.ParseFiles("views/item/edit.html")
		temp.Execute(response, data)

	}
	
}

func Delete(response http.ResponseWriter, request *http.Request) {
	
	queryString := request.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"),10, 64)

	itemModel.Delete(id)

	http.Redirect(response, request, "/item", http.StatusSeeOther)
}
