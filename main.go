package main

import (
	// digunakan untuk encoding (marshaling) dan decoding (unmarshaling) data dalam format JSON. 
	"encoding/json"
	//Paket fmt digunakan untuk format input dan output di Golang. 
	"fmt"
	//Paket io/ioutil di Golang digunakan untuk membantu dalam operasi pembacaan dan penulisan berkas (file Input/Output).
	"io/ioutil"
	//Paket os di Golang menyediakan fungsionalitas untuk berinteraksi dengan sistem operasi
	"os"
	//Paket time di Golang menyediakan fungsi-fungsi untuk pengolahan waktu dan penanganan waktu. 
	"time"
)

// Struktur data untuk menyimpan informasi tentang event
type Event struct {
	Nama    string `json:"nama"`
	Tanggal string `json:"tanggal"`
	Lokasi  string `json:"lokasi"`
}

// Struktur data untuk menyimpan informasi tentang cosplayer
type Cosplayer struct {
	Nama     string `json:"nama"`
	Karakter string `json:"karakter"`
	Anime    string `json:"anime"`
}

// Struktur data untuk menyimpan informasi tentang partisipasi cosplayer dalam event
type CosplayerEvent struct {
	Cosplayer
	EventNama string `json:"eventNama"`
}

// Struktur data untuk menyimpan keseluruhan database aplikasi menggunakan slice
type Database struct {
	Events     []Event          `json:"events"`
	Cosplayers []Cosplayer      `json:"cosplayers"`
	CosEvents  []CosplayerEvent `json:"cosplayerEvents"`
}

const databaseFile = "database.json"
const textDatabaseFile = "database.json"

func main() {
	var db Database

	loadDatabase(&db)

	for {
		fmt.Println(`
======================================================
=        Program Registrasi Event Jejepangan         =
======================================================
|1. Tambah Event                                     |
|2. Cari Event                                       |
|3. Hapus Event                                      |
|4. Tambah Cosplayer                                 | 
|5. Cari Cosplayer                                   |
|6. Hapus Cosplayer                                  |
|7. Tambah Cosplayer ke Event                        |
|8. Cari Cosplayer dalam Event                       |
|9. Hapus Cosplayer dalam Event                      |
|10. Event yang akan diadakan dalam 7 hari mendatang |
|11. Keluar dan Simpan data                          |
======================================================
`)
		fmt.Print("Pilih Menu [1-11]: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			tambahEvent(&db)
		case 2:
			cariEvent(&db)
		case 3:
			hapusEvent(&db)
		case 4:
			tambahCosplayer(&db)
		case 5:
			cariCosplayer(&db)
		case 6:
			hapusCosplayer(&db)
		case 7:
			tambahCosplayerKeEvent(&db)
		case 8:
			cariCosplayerDalamEvent(&db)
		case 9:
			hapusCosplayerDalamEvent(&db)
		case 10:
			tampilkanEventMendatang(&db)
		case 11:
			simpanDatabase(&db)
			simpanTextDatabase(&db)
			fmt.Println("Program selesai. Sampai jumpa!")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

// Fungsi json.Unmarshal untuk memuat data dari file JSON menjadi struktur data Database
func loadDatabase(db *Database) {
	fileData, err := ioutil.ReadFile(databaseFile)
	if err == nil {
		err = json.Unmarshal(fileData, db)
		if err != nil {
			fmt.Println("Gagal membaca data dari file.")
		}
	}
}

// Fungsi json.MarshalIndent untuk menyimpan data ke dalam file JSON
func simpanDatabase(db *Database) {
	data, err := json.MarshalIndent(db, "", "  ")
	if err != nil {
		fmt.Println("Gagal menyimpan data.")
		return
	}

	err = ioutil.WriteFile(databaseFile, data, 0644)
	if err != nil {
		fmt.Println("Gagal menyimpan data ke file.")
	}
}

// Fungsi untuk menambahkan event baru ke dalam database
func tambahEvent(db *Database) {
	var event Event
	fmt.Print("Nama Event: ")
	fmt.Scanln(&event.Nama)
	fmt.Print("Tanggal Event (YYYY-MM-DD): ")
	fmt.Scanln(&event.Tanggal)
	fmt.Print("Lokasi Event: ")
	fmt.Scanln(&event.Lokasi)

	db.Events = append(db.Events, event)
	fmt.Println("Event berhasil ditambahkan!")
}

// Fungsi untuk mencari event berdasarkan nama
func cariEvent(db *Database) {
	var eventName string
	fmt.Print("Masukkan nama event: ")
	fmt.Scanln(&eventName)

	for _, event := range db.Events {
		if event.Nama == eventName {
			fmt.Println("Event ditemukan!")
			fmt.Printf("Nama Event: %s\nTanggal Event: %s\nLokasi Event: %s\n", event.Nama, event.Tanggal, event.Lokasi)
			return
		}
	}

	fmt.Println("Event tidak ditemukan.")
}

// Fungsi untuk menghapus event dari database
func hapusEvent(db *Database) {
	fmt.Println("Daftar Event:")
	for i, event := range db.Events {
		fmt.Printf("%d. %s\n", i+1, event.Nama)
	}

	var choice int
	fmt.Print("Pilih Event yang akan dihapus [1-", len(db.Events), "]: ")
	fmt.Scanln(&choice)

	if choice < 1 || choice > len(db.Events) {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	eventName := db.Events[choice-1].Nama
	db.Events = append(db.Events[:choice-1], db.Events[choice:]...)
	fmt.Printf("Event %s berhasil dihapus!\n", eventName)
}

// Fungsi untuk menambahkan cosplayer baru ke dalam database
func tambahCosplayer(db *Database) {
	var cosplayer Cosplayer

	fmt.Print("Nama Cosplayer: ")
	fmt.Scanln(&cosplayer.Nama)
	fmt.Print("Nama Karakter: ")
	fmt.Scanln(&cosplayer.Karakter)
	fmt.Print("Nama Anime: ")
	fmt.Scanln(&cosplayer.Anime)

	db.Cosplayers = append(db.Cosplayers, cosplayer)
	fmt.Println("Cosplayer berhasil ditambahkan!")
}

// Fungsi untuk mencari cosplayer berdasarkan nama
func cariCosplayer(db *Database) {
	var cosplayerName string
	fmt.Print("Masukkan nama cosplayer: ")
	fmt.Scanln(&cosplayerName)

	for _, cosplayer := range db.Cosplayers {
		if cosplayer.Nama == cosplayerName {
			fmt.Println("Cosplayer ditemukan!")
			fmt.Printf("Nama Cosplayer: %s\nNama Karakter: %s\nNama Anime: %s\n", cosplayer.Nama, cosplayer.Karakter, cosplayer.Anime)
			return
		}
	}

	fmt.Println("Cosplayer tidak ditemukan.")
}

// Fungsi untuk menghapus cosplayer dari database
func hapusCosplayer(db *Database) {
	fmt.Println("Daftar Cosplayer:")
	for i, cosplayer := range db.Cosplayers {
		fmt.Printf("%d. %s\n", i+1, cosplayer.Nama)
	}

	var choice int
	fmt.Print("Pilih Cosplayer yang akan dihapus [1-", len(db.Cosplayers), "]: ")
	fmt.Scanln(&choice)

	if choice < 1 || choice > len(db.Cosplayers) {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	cosplayerName := db.Cosplayers[choice-1].Nama
	db.Cosplayers = append(db.Cosplayers[:choice-1], db.Cosplayers[choice:]...)
	fmt.Printf("Cosplayer %s berhasil dihapus!\n", cosplayerName)
}

// Fungsi untuk menambahkan cosplayer ke dalam suatu event
func tambahCosplayerKeEvent(db *Database) {
	fmt.Println("Daftar Cosplayer:")
	for i, cosplayer := range db.Cosplayers {
		fmt.Printf("%d. %s\n", i+1, cosplayer.Nama)
	}

	var cosplayerChoice int
	fmt.Print("Pilih Cosplayer yang akan ditambahkan ke event [1-", len(db.Cosplayers), "]: ")
	fmt.Scanln(&cosplayerChoice)

	if cosplayerChoice < 1 || cosplayerChoice > len(db.Cosplayers) {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	fmt.Println("Daftar Event:")
	for i, event := range db.Events {
		fmt.Printf("%d. %s\n", i+1, event.Nama)
	}

	var eventChoice int
	fmt.Print("Pilih Event yang akan ditambahkan cosplayer [1-", len(db.Events), "]: ")
	fmt.Scanln(&eventChoice)

	if eventChoice < 1 || eventChoice > len(db.Events) {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	cosplayerEvent := CosplayerEvent{
		Cosplayer: db.Cosplayers[cosplayerChoice-1],
		EventNama: db.Events[eventChoice-1].Nama,
	}

	db.CosEvents = append(db.CosEvents, cosplayerEvent)
	fmt.Printf("Cosplayer %s berhasil ditambahkan ke event %s!\n", cosplayerEvent.Nama, cosplayerEvent.EventNama)
}

// Fungsi untuk mencari cosplayer dalam suatu event
func cariCosplayerDalamEvent(db *Database) {
	fmt.Println("Daftar Cosplayer:")
	for i, cosplayer := range db.Cosplayers {
		fmt.Printf("%d. %s\n", i+1, cosplayer.Nama)
	}

	var cosplayerChoice int
	fmt.Print("Pilih Cosplayer yang akan dicari dalam event [1-", len(db.Cosplayers), "]: ")
	fmt.Scanln(&cosplayerChoice)

	if cosplayerChoice < 1 || cosplayerChoice > len(db.Cosplayers) {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	cosplayerName := db.Cosplayers[cosplayerChoice-1].Nama

	fmt.Println("Daftar Event:")
	for i, event := range db.CosEvents {
		if event.Cosplayer.Nama == cosplayerName {
			fmt.Printf("%d. %s\n", i+1, event.EventNama)
		}
	}

	fmt.Print("Pilih Event yang ingin dilihat cosplayernya [1-", len(db.CosEvents), "]: ")
	var eventChoice int
	fmt.Scanln(&eventChoice)

	if eventChoice < 1 || eventChoice > len(db.CosEvents) {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	cosplayerEvent := db.CosEvents[eventChoice-1]
	fmt.Printf("Cosplayer %s ditemukan dalam event %s!\n", cosplayerEvent.Cosplayer.Nama, cosplayerEvent.EventNama)
	fmt.Printf("Nama Cosplayer: %s\nNama Karakter: %s\nNama Anime: %s\n", cosplayerEvent.Cosplayer.Nama, cosplayerEvent.Cosplayer.Karakter, cosplayerEvent.Cosplayer.Anime)
}

// Fungsi untuk menghapus cosplayer dari suatu event
func hapusCosplayerDalamEvent(db *Database) {
	fmt.Println("Daftar Cosplayer:")
	for i, cosplayer := range db.Cosplayers {
		fmt.Printf("%d. %s\n", i+1, cosplayer.Nama)
	}

	var cosplayerChoice int
	fmt.Print("Pilih Cosplayer yang akan dihapus dalam event [1-", len(db.Cosplayers), "]: ")
	fmt.Scanln(&cosplayerChoice)

	if cosplayerChoice < 1 || cosplayerChoice > len(db.Cosplayers) {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	cosplayerName := db.Cosplayers[cosplayerChoice-1].Nama

	fmt.Println("Daftar Event:")
	for i, event := range db.CosEvents {
		if event.Cosplayer.Nama == cosplayerName {
			fmt.Printf("%d. %s\n", i+1, event.EventNama)
		}
	}

	fmt.Print("Pilih Event yang ingin dihapus cosplayernya [1-", len(db.CosEvents), "]: ")
	var eventChoice int
	fmt.Scanln(&eventChoice)

	if eventChoice < 1 || eventChoice > len(db.CosEvents) {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	cosplayerEvent := db.CosEvents[eventChoice-1]
	db.CosEvents = append(db.CosEvents[:eventChoice-1], db.CosEvents[eventChoice:]...)
	fmt.Printf("Cosplayer %s berhasil dihapus dari event %s!\n", cosplayerEvent.Cosplayer.Nama, cosplayerEvent.EventNama)
}

// Fungsi untuk menampilkan event yang akan diadakan dalam 7 hari mendatang
func tampilkanEventMendatang(db *Database) {
	fmt.Println("Event yang akan diadakan dalam 7 hari mendatang adalah:")

	for _, event := range db.Events {
		eventDate, err := time.Parse("2006-01-02", event.Tanggal)
		if err == nil && time.Until(eventDate).Hours() <= 7*24 {
			fmt.Printf("%s - %s\n", event.Nama, eventDate.Format("2006-01-02"))
		}
	}
}

// Fungsi untuk menyimpan data ke dalam file teks
func simpanTextDatabase(db *Database) {
	file, err := os.Create(textDatabaseFile)
	if err != nil {
		fmt.Println("Gagal membuat file json.")
		return
	}
	defer file.Close()

	// Menyimpan data event ke dalam file teks
	for _, event := range db.Events {
		_, _ = fmt.Fprintf(file, "Event: %s\nTanggal: %s\nLokasi: %s\n\n", event.Nama, event.Tanggal, event.Lokasi)
	}

	// Menyimpan data cosplayer ke dalam file teks
	for _, cosplayer := range db.Cosplayers {
		_, _ = fmt.Fprintf(file, "Cosplayer: %s\nKarakter: %s\nAnime: %s\n\n", cosplayer.Nama, cosplayer.Karakter, cosplayer.Anime)
	}

	// Menyimpan data partisipasi cosplayer dalam event ke dalam file teks
	for _, cosEvent := range db.CosEvents {
		_, _ = fmt.Fprintf(file, "Cosplayer: %s\nEvent: %s\n\n", cosEvent.Nama, cosEvent.EventNama)
	}

	fmt.Println("Data berhasil disimpan ke dalam file json.")
}
