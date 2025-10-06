package db

import (
	"database/sql"
	"log"

	"auraluvsu.com/nyx/gamestate"
)

func foo() {
	newCache := gamestate.GameCache
	db, err := sql.Open("sqlite3", "nyx.db")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	db.Exec(
		`create table if not exists games (
		game_id TEXT PRIMARY KEY,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		user_id INTEGER NOT NULL AUTOINCREMENT REFERENCES users(user_id) ON DELETE CASCADE
		);`)
	res, err := db.Exec("insert into games (game_id) values (?)", newCache.GameID)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
}
