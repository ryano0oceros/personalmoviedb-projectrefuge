# SQLite Movie Database Project

## Project Overview

This project involves setting up a SQLite relational database to manage a personal movie collection. The database is populated with movie data downloaded from Northwestern University's archive of IMDb data files. The project is implemented in Go, utilizing a cgo-free port of SQLite for database operations.

## Setting Up the SQLite Database

1. Download and extract the movie data files from [Northwestern University's archive](https://arch.library.northwestern.edu/concern/datasets/3484zh40n?locale=en).
2. Define the SQLite relational database schema to accommodate the extracted data.
3. Implement a Go program that serves as a local database server for SQLite.
4. Populate the SQLite database with the movie data.

## Performing SQL Queries

The SQLite database allows for various SQL queries to be performed, enabling users to search, filter, and manage their movie collection effectively. Examples of SQL queries include searching for movies by genre, year, or director.

The sample code provided below demonstrates how to perform a simple SQL query to grab a specific row and count the number of movies from a particular year:

```go
db.QueryRow("SELECT * FROM movies WHERE id = 191243")
...
db.QueryRow("SELECT COUNT(*) FROM movies WHERE year = 1991")
...
```

returns the following output:
```bash
ID: 191243, Name: Lion King 1, The, Year: 2004, Rank: 0.000000
Total movies from 1991: 6065
```

## Using the Database

The SQLite database allows for various SQL queries to be performed, enabling users to search, filter, and manage their movie collection. Beyond SQL queries, the Go application provides a user-friendly interface for interacting with the database. Users can easily add, update, or remove movies from their personal collection, as well as view detailed information about each movie.

## Potential Enhancements and Personal Movie Recommendation System

Some possible future enhancements, could include: 
- Personalized Recommendations: This system could leverage machine learning algorithms to provide this based on ratings and viewing history, the application can recommend movies from both your collection and external sources like IMDb.
- Inventory Management: Easily track where each movie is located (e.g., hard drives, streaming services, physical DVDs).
- Automated Web Scraping: could also be implemented to update the database with metadata (similarly to how Plex media server works) to include things like the movie poster, actors/actresses, etc..
- Detailed Reviews: Add personal reviews and notes to each movie, creating a comprehensive personal movie diary.
- Custom User Interface: can build in a web-based frontend for easy access and management of the movie collection with features like search/filter/sort, Add/Update/Delete, and visualizations.


## Advantages of the Go Movie Application

The Go movie application offers several advantages over IMDb alone. It provides
- Personalized experience: users can manage their own movie collection and ratings
- Offline access to movie data: which can be useful in areas with limited internet access
- Security and privacy: from advertisers and platforms looking to monetize your personal data
- Customization: solution can be expanded upon to meet individual user needs

## Conclusion

The SQLite movie database project demonstrates the feasibility of using Go and SQLite for managing a personal movie collection. The project lays the groundwork for further enhancements, listed above.
