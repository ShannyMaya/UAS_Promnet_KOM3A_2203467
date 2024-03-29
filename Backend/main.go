package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	Routers()
}

func Routers() {
	InitDB()
	defer db.Close()
	log.Println("Starting the HTTP server on port 9080")
	router := mux.NewRouter()
	router.HandleFunc("/users",
		GetUsers).Methods("GET")
	router.HandleFunc("/users",
		CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}",
		GetUser).Methods("GET")
	router.HandleFunc("/users/{id}",
		UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}",
		DeleteUser).Methods("DELETE")
	http.ListenAndServe(":9080",
		&CORSRouterDecorator{router})
}

/***************************************************/

//Get all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User

	result, err := db.Query("SELECT id," +
		"nama_barang, jumlah, harga_satuan, lokasi, deskripsi from item")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var user User
		err := result.Scan(&user.ID, &user.NamaBarang,
			&user.Jumlah, &user.HargaSatuan, &user.Lokasi, &user.Deskripsi)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users)
}

//Create user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	stmt, err := db.Prepare("INSERT INTO item (nama_barang,jumlah,harga_satuan,lokasi,deskripsi) VALUES(?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	nama_barang := keyVal["nama_barang"]
	jumlah := keyVal["jumlah"]
	harga_satuan := keyVal["harga_satuan"]
	lokasi := keyVal["lokasi"]
	deskripsi := keyVal["deskripsi"]
	// print jenis_kelamin
	//fmt.Println(jenis_kelamin)
	_, err = stmt.Exec(nama_barang,jumlah,harga_satuan,lokasi,deskripsi)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New item was created")
}

//Get user by ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT id,"+
		"nama_barang,jumlah,harga_satuan,lokasi,deskripsi from item WHERE id = ?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var user User
	for result.Next() {
		err := result.Scan(&user.ID, &user.NamaBarang,
			&user.Jumlah, &user.HargaSatuan, &user.Lokasi, &user.Deskripsi)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(user)
}

//Update user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("UPDATE item SET nama_barang= ?, jumlah=?, harga_satuan=?, lokasi=?, deskripsi=? WHERE id = ?")

	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	nama_barang := keyVal["nama_barang"]
	jumlah := keyVal["jumlah"]
	harga_satuan := keyVal["harga_satuan"]
	lokasi := keyVal["lokasi"]
	deskripsi := keyVal["deskripsi"]
	_, err = stmt.Exec(nama_barang,jumlah,harga_satuan,lokasi,deskripsi, params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Item with ID = %s was updated",
		params["id"])
}


func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// stmt, err := db.Prepare("DELETE FROM users WHERE id = ?")
	stmt, err := db.Prepare("DELETE FROM item WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Item with ID = %s was deleted",
		params["id"])
}



type User struct {
	ID        string `json:"id"`
	NamaBarang		   string `json:"nama_barang"`
	Jumlah     string `json:"jumlah"`
	HargaSatuan  string `json:"harga_satuan"`
	Lokasi   string `json:"lokasi"`
	Deskripsi	string `json:"deskripsi"`

}


var db *sql.DB
var err error

func InitDB() {
	db, err = sql.Open("mysql",
		"root:@tcp(127.0.0.1:3306)/db_2203467_shanny_uas")
	if err != nil {
		panic(err.Error())
	}
}

/***************************************************/

// CORSRouterDecorator applies CORS headers to a mux.Router
type CORSRouterDecorator struct {
	R *mux.Router
}

func (c *CORSRouterDecorator) ServeHTTP(rw http.ResponseWriter,
	req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods",
			"POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers",
			"Accept, Accept-Language,"+
				" Content-Type, YourOwnHeader")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}

	c.R.ServeHTTP(rw, req)
}
