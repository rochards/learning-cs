package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Album struct {
	Id     int64
	Title  string
	Artist string
	Price  float32
}

func albumsByArtist(name string) ([]Album, error) {
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	for rows.Next() {
		var album Album
		if err := rows.Scan(&album.Id, &album.Title, &album.Artist, &album.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, album)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	return albums, nil
}

func albumById(id int64) (Album, error) {
	var album Album
	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&album.Id, &album.Title, &album.Artist, &album.Price); err != nil {
		if err == sql.ErrNoRows {
			return album, fmt.Errorf("albumsById %d: no such album", id)
		}
		return album, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return album, nil
}

func addAlbum(album Album) (int64, error) {
	result, err := db.Exec("INSERT INTO album(title, artist, price) VALUES (?,?,?)", album.Title, album.Artist, album.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	if dbUser == "" || dbPass == "" {
		log.Fatal("missing required env vars: DBUSER and/or DBPASS")
	}

	cfg := mysql.NewConfig()
	cfg.User = dbUser
	cfg.Passwd = dbPass
	cfg.Net = "tcp"
	cfg.Addr = "localhost:3306"
	cfg.DBName = "recordings"
	cfg.ParseTime = true

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalf("sql.Open failed (addr=%s db=%s user=%s): %v", cfg.Addr, cfg.DBName, cfg.User, err)
	}

	if pingErr := db.Ping(); pingErr != nil {
		log.Fatalf("db.Ping failed (addr=%s db=%s user=%s): %v", cfg.Addr, cfg.DBName, cfg.User, pingErr)
	}
	fmt.Printf("Connected to MySQL at %s (db=%s user=%s)\n", cfg.Addr, cfg.DBName, cfg.User)

	name := "John Coltrane"
	albums, err := albumsByArtist(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums for %s found: %v\n", name, albums)

	id := int64(2)
	album, err := albumById(id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album %d found: %v\n", id, album)

	albumId, err := addAlbum(Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albumId)
}
