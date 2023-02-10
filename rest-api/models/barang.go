package models

type Barang struct {
	Kodebarang string `gorm:"primaryKey" json:"id"`
	Nama       string `gorm:"column:namabarang" json:"nama"`
	Satuan     int    `json:"satuan"`
	Stock      int    `json:"stock"`
	Harga      string `gorm:"column:hargasatuan" json:"harga"`
}
