package models

type Barang struct {
	Id int    `gorm:"primaryKey" json:"id"`
	Nama string `gorm:"type:varchar(50)" json:"nama"`
	Satuan     int    `json:"satuan"`
	Stock      int    `json:"stock"`
	Harga      int    `gorm:"type:numeric" json:"harga"`
}
