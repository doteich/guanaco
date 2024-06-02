package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "modernc.org/sqlite"
)

type Row struct {
	id       int
	ts       time.Time
	nodeName string
	nodeId   string
	dataType string
	value    interface{}
}

func GetUniqueNodes(n string, t string) ([]string, error) {

	results := make([]string, 0)

	db, err := connectToDB(n)

	if err != nil {
		return results, err
	}

	defer db.Close()

	rows, err := db.Query(fmt.Sprintf("SELECT DISTINCT %s FROM guanaco", t))

	if err != nil {
		return results, err
	}

	for rows.Next() {

		var nodeId string

		if err := rows.Scan(&nodeId); err != nil {
			return results, err
		}

		results = append(results, nodeId)
	}

	return results, nil
}

func connectToDB(n string) (*sql.DB, error) {

	wd, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("%s/services/%s/sqlite/data.db", wd, n)
	db, err := sql.Open("sqlite", path)

	if err != nil {
		return nil, err
	}

	return db, nil
}
