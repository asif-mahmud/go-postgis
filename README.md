# go-postgis
![go workflow](https://github.com/asif-mahmud/go-postgis/actions/workflows/go.yml/badge.svg)

Golang support for PostGIS datatypes

## Features
1. Fully tested
2. Supports both little endian and big endian byte orders
3. Simple usage of PostGIS datatypes in `struct`s or standalone variables
4. Supports any postgresql driver that utilizes `sql.Scanner` and `driver.Valuer` interfaces
5. Out of the box support for json marshal/unmarshal

## Installation
To add the package to your project run -

```
go get -u github.com/asif-mahmud/go-postgis
```

### Documentation

godoc: [https://pkg.go.dev/github.com/asif-mahmud/go-postgis](https://pkg.go.dev/github.com/asif-mahmud/go-postgis)
well known binary format - [https://github.com/postgis/postgis/blob/master/doc/bnf-wkb.txt](https://github.com/postgis/postgis/blob/master/doc/bnf-wkb.txt) 
well known text format - [https://github.com/postgis/postgis/blob/master/doc/bnf-wkt.txt](https://github.com/postgis/postgis/blob/master/doc/bnf-wkt.txt)

### Supported datatypes
1. Point (X, Y) - added in v0.1.1
2. PointS (SRID X, Y) - added in v0.1.1
3. PointZ (X, Y, Z) - added in v0.1.2
4. PointZS (SRID X, Y, Z) - added in v0.1.2
5. PointM (X, Y, M) - added in v0.1.2
6. PointMS (SRID X, Y, M) - added in v0.1.2
7. PointZM (X, Y, Z, M) - added in v0.1.2
8. PointZMS (SRID X, Y, Z, M) - added in v0.1.2

### Example usage

```
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
```

## Version history

### Version 0.1.2

- Added `PointZ`, `PointZS`, `PointM`, `PointMS`, `PointZM` and `PointZMS` types
- Updated readme

### Version 0.1.1

- Added `Point` and `PointS` types
- Updated readme
- First release in go pkg

### Version 0.1.0

Initial version with tests for hex ewkb decoder.
