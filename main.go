package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Struktur data Anggota disesuaikan dengan database
type Anggota struct {
	IDAnggota     int    `json:"id_anggota"`
	Nama          string `json:"nama"`
	Alamat        string `json:"alamat"`
	TotalSimpanan int    `json:"total_simpanan"`
}

func main() {
	// Koneksi ke database MySQL (XAMPP)
	dsn := "root:@tcp(127.0.0.1:3306)/db_koperasi"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}
	defer db.Close()

	r := gin.Default()

	// Endpoint 1: Mengambil semua data anggota
	r.GET("/anggota", func(c *gin.Context) {
		// Menyesuaikan nama kolom dengan struktur asli
		rows, err := db.Query("SELECT id_anggota, nama, alamat, total_simpanan FROM anggota")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var listAnggota []Anggota
		for rows.Next() {
			var a Anggota
			if err := rows.Scan(&a.IDAnggota, &a.Nama, &a.Alamat, &a.TotalSimpanan); err != nil {
				continue
			}
			listAnggota = append(listAnggota, a)
		}

		c.JSON(http.StatusOK, gin.H{"data": listAnggota})
	})

	// Endpoint 2: Menambah anggota baru
	r.POST("/anggota", func(c *gin.Context) {
		var a Anggota
		if err := c.ShouldBindJSON(&a); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Menyesuaikan query INSERT
		result, err := db.Exec("INSERT INTO anggota (nama, alamat, total_simpanan) VALUES (?, ?, ?)", a.Nama, a.Alamat, a.TotalSimpanan)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		id, _ := result.LastInsertId()
		a.IDAnggota = int(id)
		c.JSON(http.StatusCreated, gin.H{"message": "Anggota berhasil ditambahkan", "data": a})
	})

	r.Run(":8080")
}
