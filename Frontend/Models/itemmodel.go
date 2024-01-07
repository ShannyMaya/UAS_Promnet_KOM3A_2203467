package models

import (
	"fmt"
	"database/sql"
	"github.com/jeypc/go-crud/config"
	"github.com/jeypc/go-crud/entities"
)

type ItemModel struct {
	conn *sql.DB
}

func NewItemModel() *ItemModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &ItemModel{
		conn: conn,
	}
}

func (p *ItemModel) FindAll() ([]entities.Item, error) {
	
	rows, err := p.conn.Query("select * from item")
	if err != nil {
		return []entities.Item{}, err
	}
	defer rows.Close()

	var dataItem []entities.Item
	for rows.Next(){
		var item entities.Item
		rows.Scan(&item.Id, &item.NamaBarang, &item.Jumlah, &item.HargaSatuan, &item.Lokasi, &item.Deskripsi)

		if item.Lokasi == "1"{
			item.Lokasi = "Bandung"
		} else if item.Lokasi == "2"{
			item.Lokasi = "Jakarta"
		}else if item.Lokasi == "3"{
			item.Lokasi = "Denpasar"
		}else if item.Lokasi == "4"{
			item.Lokasi = "Manokwari"
		}

		dataItem = append(dataItem, item)
	}
	return dataItem, nil
}

func (p *ItemModel) Detail(id int64) (*entities.Item, error) {
    row := p.conn.QueryRow("SELECT * FROM item WHERE id = ?", id)

    var item entities.Item

    err := row.Scan(
        &item.Id,
        &item.NamaBarang,
        &item.Jumlah,
        &item.HargaSatuan,
        &item.Lokasi,
        &item.Deskripsi,
    )

    if err != nil {
        return nil, err
    }

    return &item, nil
}

func (p *ItemModel) Create(item entities.Item) bool {
	result, err := p.conn.Exec("insert into item (nama_barang, jumlah, harga_satuan, lokasi, deskripsi) values(?,?,?,?,?)", 
	item.NamaBarang, item.Jumlah, item.HargaSatuan, item.Lokasi, item.Deskripsi)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}

func (p *ItemModel) Find(id int64, item *entities.Item) error {

	return p.conn.QueryRow("select * from item where id = ?", id).Scan(
		&item.Id, 
		&item.NamaBarang, 
		&item.Jumlah, 
		&item.HargaSatuan, 
		&item.Lokasi, 
		&item.Deskripsi)
}

func (p *ItemModel) Update(item entities.Item) error {

	_, err := p.conn.Exec("update item set nama_barang = ?, jumlah = ?, harga_satuan = ?, lokasi = ?, deskripsi = ? where id = ?",
	item.NamaBarang, item.Jumlah, item.HargaSatuan, item.Lokasi, item.Deskripsi, item.Id)

	if err != nil {
		return err
	}

	return nil
}

func (p * ItemModel) Delete(id int64) {
	p.conn.Exec("delete from item where id = ?", id)
}
