package internal

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const file string = "mydb.db"

func Dbstart() *sql.DB {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		println("erro ao abrir o banco de dados")
	}

	createResourceTableSQL := "CREATE TABLE IF NOT EXISTS resources (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT UNIQUE, price REAL);"

	resourcesMap := map[string]float64{
		"money":           1,
		"wood":            5,
		"coal":            1.5,
		"iron":            7,
		"grain":           2,
		"protein":         4,
		"processedFood":   100,
		"mechanicalParts": 100,
		"copper":          3,
		"gold":            50_000,
		"cotton":          1.2,
		"electricity":     2,
		"paper":           50,
		"research":        -500,
		"silicon":         10,
		"sulfur":          100,
		"chemicals":       1_500,
		"components":      400,
		"basicCircuit":    1_000,
		"basicChip":       10_000,
	}

	_, err = db.Exec(createResourceTableSQL)
	if err != nil {
		log.Fatalf("erro ao criar a tabela resources: %v", err)
	}

	resourceExistsError := "UNIQUE constraint failed: resources.name"
	for name, price := range resourcesMap {
		result, err := db.Exec("INSERT INTO resources (name, price) VALUES (?, ?)", name, price)
		if err.Error() == resourceExistsError {
			_, err := db.Exec("UPDATE resources SET price = ? WHERE name = ?", price, name)
			if err != nil {
				log.Fatalf("erro ao atualizar a resource %v: %v", name, err)
			}
		} else if err != nil {
			log.Fatal("Erro não conhecido ao inserir a resource")
		} else {
			println(result)
		}
	}
	return db
}

func Dbclose(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatal("erro ao fechar o banco de dados")
	}
}

func DbQuerytest() []string {
	db := Dbstart()
	result, err := db.Query("SELECT * FROM test")
	if err != nil {
		println("erro ao fazer a query")
		log.Fatal(err)
	}
	var toReturn []string
	for result.Next() {
		var id int
		var name string
		err := result.Scan(&id, &name)
		if err != nil {
			println("erro ao fazer o parsing dos dados")
		}
		toReturn = append(toReturn, name)
	}
	Dbclose(db)
	return toReturn
}

func DbQueryResources() map[string]float64 {
	db := Dbstart()
	result, err := db.Query("SELECT * FROM resources")
	if err != nil {
		println("erro ao fazer a query")
		log.Fatal(err)
	}
	toReturn := make(map[string]float64)
	for result.Next() {
		var id int
		var name string
		var price float64
		err := result.Scan(&id, &name, &price)
		if err != nil {
			println("erro ao fazer o parsing dos dados")
		}
		toReturn[name] = price
	}
	Dbclose(db)
	return toReturn
}
