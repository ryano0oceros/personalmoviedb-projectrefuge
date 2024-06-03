package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "modernc.org/sqlite"
)

func main() {
	if err := main1(); err != nil {
		log.Fatal(err)
	}
}

func main1() error {
	dir, err := os.MkdirTemp("", "test-")
	if err != nil {
		return err
	}

	defer os.RemoveAll(dir)

	fn := filepath.Join(dir, "movie_database.db")

	db, err := sql.Open("sqlite", fn)
	if err != nil {
		return err
	}
	defer db.Close()

	// // Create tables
	// if err := createTables(db); err != nil {
	// 	return err
	// }

	// Populate database with movie data
	if err := populateDatabase(db); err != nil {
		return err
	}

	// Query to get a specific row
	row := db.QueryRow("SELECT * FROM movies WHERE id = 191243")
	var id int
	var name string
	var year int
	var rank float64
	err = row.Scan(&id, &name, &year, &rank)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID: %d, Name: %s, Year: %d, Rank: %f\n", id, name, year, rank)

	// Remove quotation marks from name
	name = strings.ReplaceAll(name, "\"", "")

	// Query to get the total number of movies from 2022
	row = db.QueryRow("SELECT COUNT(*) FROM movies WHERE year = 1991")
	var count int
	err = row.Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total movies from 1991: %d\n", count)

	return nil
}

func populateDatabase(db *sql.DB) error {
	// Create table
	query := `
	CREATE TABLE movies (
		id INTEGER,
		name TEXT,
		year INTEGER,
		rank REAL
	);
	`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	// Define file names
	files := []string{
		//		"../database/IMDB-directors_genres.csv",
		//		"../database/IMDB-directors.csv",
		//		"../database/IMDB-movies_genres.csv",
		"../database/IMDB-movies.csv",
		//		"../database/IMDB-roles.csv",
	}

	// Define queries for each file
	queries := []string{
		//		"INSERT INTO directors_genres (director_id, genre, prob) VALUES (?, ?, ?)",
		//		"INSERT INTO directors (id, first_name, last_name) VALUES (?, ?, ?)",
		//		"INSERT INTO movies_genres (movie_id, genre) VALUES (?, ?)",
		"INSERT INTO movies (id, name, year, rank) VALUES (?, ?, ?, ?)",
		//		"INSERT INTO roles (actor_id, movie_id, role) VALUES (?, ?, ?)",
	}

	// Loop over files
	for i, file := range files {
		// Open file
		f, err := os.Open(file)
		if err != nil {
			return err
		}
		defer f.Close()

		// Create CSV reader
		r := csv.NewReader(f)

		// Loop over rows
		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("CSV parsing error on line: %v", err)
				continue
			}

			// Convert []string to []interface{}
			recordInterface := make([]interface{}, len(record))
			for i, v := range record {
				// Remove quotation marks from name field
				if i == 1 {
					v = strings.ReplaceAll(v, "\"", "")
				}
				// Convert "NULL" to nil
				if v == "NULL" {
					recordInterface[i] = 0
				} else {
					recordInterface[i] = v
				}
			}

			// Insert data into database
			_, err = db.Exec(queries[i], recordInterface...)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
