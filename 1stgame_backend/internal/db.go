package internal

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const file string = "mydb.db?_busy_timeout=5000"

func DbInitializeIndustries(db *sql.DB) {
	if db == nil {
		log.Fatal("db não pode ser nulo")
	}

	createIndustriesTableSQL := "CREATE TABLE IF NOT EXISTS industries (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT UNIQUE);"
	_, err := db.Exec(createIndustriesTableSQL)
	if err != nil {
		log.Fatalf("erro desconhecido encontrado ao criar tabela \"industries\": %v", err)
	} else {
		println(`nenhum erro ao iniciar tabela "industries"`)
	}

	for _, industry := range industries {
		_, err := db.Exec("INSERT INTO industries (name) VALUES (?)", industry.Name)
		if err != nil {
			if err.Error() == "UNIQUE constraint failed: industries.name" {
			} else {
				log.Fatalf("erro desconhecido ao inserir a industry %v: %v", industry.Name, err)
			}
		}
	}
}

type industriesConsumption struct {
	id           int
	ResourceName string
	IndustryName string
	Quantity     float64
}
type industriesProduction struct {
	id           int
	ResourceName string
	IndustryName string
	Quantity     float64
}

func DbInitializeIndustriesConsumption(db *sql.DB) {
	if db == nil {
		log.Fatal("db não pode ser nulo")
	}

	createIndustriesConsumptionTableSQL := `CREATE TABLE IF NOT EXISTS industries_consumption (resource_name TEXT, industry_name TEXT, quantity REAL, FOREIGN KEY (industry_name) REFERENCES industries(name));`

	_, err := db.Exec(createIndustriesConsumptionTableSQL)
	if err != nil {
		log.Fatalf("erro desconhecido encontrado ao criar tabela \"industries_consumption\": %v", err)
	} else {
		println(`nenhum erro ao iniciar tabela "industries_consumption"`)
	}
	if err != nil {
		log.Fatalf("erro desconhecido encontrado ao fazer a query na tabela \"industries_consumption\": %v", err)
	}

	deleteAllIndustriesConsumptionSQL := "DELETE FROM industries_consumption"
	_, err = db.Exec(deleteAllIndustriesConsumptionSQL)
	if err != nil {
		log.Fatalf("erro desconhecido encontrado ao deletar os dados da tabela \"industries_consumption\": %v", err)
	} else {
		insertIndustriesConsumptionSQL := "INSERT INTO industries_consumption (resource_name, industry_name, quantity) VALUES (?, ?, ?);"
		for _, industry := range industries {
			for resourceName, quantity := range industry.Consumption {
				_, err := db.Exec(insertIndustriesConsumptionSQL, resourceName, industry.Name, quantity)
				if err != nil {
					log.Fatalf("erro desconhecido encontrado ao inserir os dados na tabela \"industries_consumption\": %v", err)
				}
			}
		}
	}
}

func DbInitializeEconomyMinisteryIndustries(db *sql.DB) {
	if db == nil {
		log.Fatal("db não pode ser nulo")
	}
	createEconomyMinisteryIndustriesTableSQL := `CREATE TABLE IF NOT EXISTS economy_ministery_industries (industry_name TEXT, price INT, cardTitle TEXT, cardImage TEXT, FOREIGN KEY (industry_name) REFERENCES industries(name));`

	_, err := db.Exec(createEconomyMinisteryIndustriesTableSQL)
	if err != nil {
		log.Fatalf("erro desconhecido encontrado ao criar tabela \"economy_ministery_industries\": %v", err)
	} else {
		println(`nenhum erro ao iniciar tabela "economy_ministery_industries"`)
	}

	deleteAllEconomyMinisteryIndustriesSQL := "DELETE FROM economy_ministery_industries"
	_, err = db.Exec(deleteAllEconomyMinisteryIndustriesSQL)
	if err != nil {
		log.Fatalf("erro desconhecido encontrado ao deletar os dados da tabela \"economy_ministery_industries\": %v", err)
	} else {
		insertEconomyMinisteryIndustriesSQL := "INSERT INTO economy_ministery_industries (industry_name, price, cardTitle, cardImage) VALUES (?, ?, ?, ?);"
		for _, industry := range industries {
			_, err := db.Exec(insertEconomyMinisteryIndustriesSQL, industry.Name, industry.Price, industry.CardTitle, industry.CardImage)
			if err != nil {
				log.Fatalf("erro desconhecido encontrado ao inserir os dados na tabela \"economy_ministery_industries\": %v", err)
			}
		}
	}
}

func DbInitializeIndustriesProduction(db *sql.DB) {
	if db == nil {
		log.Fatal("db não pode ser nulo")
	}

	createIndustriesProductionTableSQL := `CREATE TABLE IF NOT EXISTS industries_production (resource_name TEXT, industry_name TEXT, quantity REAL, FOREIGN KEY (industry_name) REFERENCES industries(name));`

	_, err := db.Exec(createIndustriesProductionTableSQL)
	if err != nil {
		log.Fatalf("erro desconhecido encontrado ao criar tabela \"industries_production\": %v", err)
	} else {
		println(`nenhum erro ao iniciar tabela "industries_production"`)
	}
	if err != nil {
		log.Fatalf("erro desconhecido encontrado ao fazer a query na tabela \"industries_production\": %v", err)
	}

	deleteAllIndustriesProductionSQL := "DELETE FROM industries_production"
	_, err = db.Exec(deleteAllIndustriesProductionSQL)
	if err != nil {
		log.Fatalf("erro desconhecido encontrado ao deletar os dados da tabela \"industries_production\": %v", err)
	} else {
		insertIndustriesProductionSQL := "INSERT INTO industries_production (resource_name, industry_name, quantity) VALUES (?, ?, ?);"
		for _, industry := range industries {
			for resourceName, quantity := range industry.Production {
				_, err := db.Exec(insertIndustriesProductionSQL, resourceName, industry.Name, quantity)
				if err != nil {
					log.Fatalf("erro desconhecido encontrado ao inserir os dados na tabela \"industries_production\": %v", err)
				}
			}
		}
	}
}

func findIndexByName(industries []Industry, name string) int {
	for i, industry := range industries {
		if industry.Name == name {
			return i
		}
	}
	return -1
}

func Dbstart() *sql.DB {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		println("erro ao abrir o banco de dados")
	}
	if err := db.Ping(); err != nil {
		log.Fatal("erro ao conectar-se ao banco de dados:", err)
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

	DbInitializeIndustries(db)
	DbInitializeIndustriesConsumption(db)
	DbInitializeIndustriesProduction(db)
	DbInitializeEconomyMinisteryIndustries(db)

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

func DbQueryIndustriesConsumption() []industriesConsumption {
	db := Dbstart()
	result, err := db.Query("SELECT * FROM industries_consumption")
	if err != nil {
		println("erro ao fazer a query")
		log.Fatal(err)
	}
	var toReturn []industriesConsumption
	for result.Next() {
		var id int
		var resourceName string
		var industryName string
		var quantity float64
		err := result.Scan(&resourceName, &industryName, &quantity)
		if err != nil {
			println("erro ao fazer o parsing dos dados")
		}
		toReturn = append(toReturn, industriesConsumption{id, resourceName, industryName, quantity})
	}
	Dbclose(db)
	return toReturn
}

func DbQueryIndustriesProduction() []industriesProduction {
	db := Dbstart()
	result, err := db.Query("SELECT * FROM industries_production")
	if err != nil {
		println("erro ao fazer a query")
		log.Fatal(err)
	}
	var toReturn []industriesProduction
	for result.Next() {
		var id int
		var resourceName string
		var industryName string
		var quantity float64
		err := result.Scan(&resourceName, &industryName, &quantity)
		if err != nil {
			println("erro ao fazer o parsing dos dados")
		}
		toReturn = append(toReturn, industriesProduction{id, resourceName, industryName, quantity})
	}
	Dbclose(db)
	return toReturn
}

type EconomyMinisteryIndustry struct {
	Id        string
	Price     int
	CardTitle string
	CardImage string
}

func DbQueryEconomyMinisteryIndustries() []EconomyMinisteryIndustry {
	db := Dbstart()
	result, err := db.Query("SELECT * FROM economy_ministery_industries")
	if err != nil {
		println("erro ao fazer a query")
		log.Fatal(err)
	}
	if result == nil {
		println("result é nulo")
	}
	var toReturn []EconomyMinisteryIndustry
	for result.Next() {
		var id string
		var price int
		var cardTitle string
		var cardImage string
		err := result.Scan(&id, &price, &cardTitle, &cardImage)
		if err != nil {
			println("erro ao fazer o parsing dos dados")
		}
		toReturn = append(toReturn, EconomyMinisteryIndustry{id, price, cardTitle, cardImage})
	}
	Dbclose(db)
	return toReturn
}
