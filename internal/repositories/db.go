package repositories

import "github.com/jmoiron/sqlx"

func NewDb() (*sqlx.DB, error) {
	//с каких пор у нас не db а dba?
	db, err := sqlx.Open("sqlite3", "notes.dba")
	db.Exec("create table users" +
		"(" +
		"    id       INTEGER not null" +
		"	primary key autoincrement," +
		"    login    TEXT," +
		"    password TEXT" +
		");")
	db.Exec("create table notes" +
		"(" +
		"    id      TEXT" +
		"        primary key," +
		"    user_id INTEGER" +
		"        constraint notes_users_id_fk" +
		"            references users," +
		"    title   TEXT," +
		"    body    TEXT);")
	if err != nil {
		return nil, err
	}
	//зачем мы подключение закрываем? а как бдхой-то пользоваться
	db.Close()
	return db, nil
}
