package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "modernc.org/sqlite"
)

type Row struct {
	Id       int         `json:"id"`
	Ts       time.Time   `json:"ts"`
	NodeName string      `json:"nodeName"`
	NodeId   string      `json:"nodeId"`
	DataType string      `json:"dataType"`
	Value    interface{} `json:"value"`
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

		var selection string

		if err := rows.Scan(&selection); err != nil {
			return results, err
		}

		results = append(results, selection)
	}

	return results, nil
}

func GetTimeSeries(n string, nodeId string, nodeName string, start string, end string) ([]Row, error) {

	results := make([]Row, 0)

	db, err := connectToDB(n)

	if err != nil {
		return results, err
	}

	defer db.Close()

	s, err := time.Parse(time.RFC3339, start)

	if err != nil {
		return results, err
	}

	e, err := time.Parse(time.RFC3339, end)

	if err != nil {
		return results, err
	}

	query := `SELECT * FROM GUANACO WHERE`

	if nodeName != "" {
		query = query + fmt.Sprintf(" nodeName = '%s' AND", nodeName)
	}
	if nodeId != "" {
		query = query + fmt.Sprintf(" nodeId = '%s' AND", nodeId)
	}

	query = query + " ts >= ? AND ts <= ?"

	rows, err := db.Query(query, s, e)

	if err != nil {
		return results, err
	}

	for rows.Next() {
		var entry Row
		rows.Scan(
			&entry.Id, &entry.Ts, &entry.NodeName, &entry.NodeId, &entry.DataType, &entry.Value,
		)
		results = append(results, entry)
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
