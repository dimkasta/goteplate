package goteplate

import (
	"database/sql"
	"github.com/dimkasta/gologger"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type SqliteTemplateRepository struct {
	Filename string
}

func (repository *SqliteTemplateRepository) Get(name string) string {
	db, err := sql.Open("sqlite3", repository.Filename)

	if nil != err {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("select html from templates where name = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var html string
	err = stmt.QueryRow(name).Scan(&html)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(html)
	return html
}

func NewSqliteTemplateRepository(filename string, logger *gologger.LoggerService) *SqliteTemplateRepository {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		logger.Error(err.Error())
	}
	defer db.Close()

	return &SqliteTemplateRepository{
		Filename: filename,
	}
}