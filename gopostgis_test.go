package gopostgis_test

import (
	"database/sql"

	gopostgis "github.com/asif-mahmud/go-postgis"
)

func Example() {
	// we know it will panic at testing time, so
	// recovering silently, ignore this part
	defer func() {
		recover()
	}()

	// this is written for purely example purpose
	// create db connection
	db, e := sql.Open("postgres", "database=test_db")
	if e != nil {
		panic(e)
	}

	// create a table with geometry column
	_, e = db.Exec(`CREATE TABLE IF NOT EXISTS test_table ( 
    id SERIAL PRIMARY KEY,
    location GEOMETRY
  )`)
	if e != nil {
		panic(e)
	}

	// construct a point
	point := gopostgis.Point{
		X:     10,
		Y:     20,
		Valid: true, // if you don't mark it as valid, null will be saved in db
	}

	// insert a point
	_, e = db.Exec(`
  INSERT INTO test_table (location) VALUES($1)`,
		point,
	)
	if e != nil {
		panic(e)
	}

	// read a point
	row := db.QueryRow(`SELECT location from test_table LIMIT 1`)
	if e := row.Scan(&point); e != nil {
		panic(e)
	}

	// Output:
}
