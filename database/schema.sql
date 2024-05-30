-- SQLite relational database schema for the movie data

CREATE TABLE Actors (
    id INTEGER PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    gender TEXT
);

CREATE TABLE Directors (
    id INTEGER PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL
);

CREATE TABLE Movies (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    year INTEGER,
    rank REAL
);

CREATE TABLE Genres (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    genre TEXT NOT NULL
);

CREATE TABLE MoviesGenres (
    movie_id INTEGER,
    genre_id INTEGER,
    FOREIGN KEY (movie_id) REFERENCES Movies(id),
    FOREIGN KEY (genre_id) REFERENCES Genres(id)
);

CREATE TABLE DirectorsGenres (
    director_id INTEGER,
    genre_id INTEGER,
    prob REAL,
    FOREIGN KEY (director_id) REFERENCES Directors(id),
    FOREIGN KEY (genre_id) REFERENCES Genres(id)
);

CREATE TABLE Roles (
    actor_id INTEGER,
    movie_id INTEGER,
    role TEXT,
    FOREIGN KEY (actor_id) REFERENCES Actors(id),
    FOREIGN KEY (movie_id) REFERENCES Movies(id)
);
