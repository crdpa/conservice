package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

const (
	user    = "conservice"
	pass    = "conservice"
	dbname  = "data"
	host    = "db"
	dbport  = 5432
	sslmode = "disable"
)

type row struct {
	cpf           string
	private       string
	incompleto    string
	ultCompra     string
	ticketMedio   float64
	ticketUltimo  float64
	lojaMaisFreq  string
	lojaUltCompra string
}

func indexHandler(c *fiber.Ctx, db *sql.DB) error {
	var rowData row
	var data []row
	rows, err := db.Query("Select * from data")
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
		c.JSON("An error ocurred.")
	}
	for rows.Next() {
		rows.Scan(&rowData.cpf, &rowData.private, &rowData.incompleto, &rowData.ultCompra, &rowData.ticketMedio, &rowData.ticketUltimo, &rowData.lojaMaisFreq, &rowData.lojaUltCompra)
		data = append(data, rowData)
	}
	return c.Render("index", fiber.Map{
		"Data": data,
	})
}

func postHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello")
}

func main() {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%d sslmode=%s", user, dbname, pass, host, dbport, sslmode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	sqlCreateTable := `CREATE TABLE IF NOT EXISTS data (
					  cpf varchar(14) NOT NULL,
					  private bool NOT NULL,
					  incompleto bool NOT NULL,
					  ultima_compra DATE,
					  ticket_medio decimal(12,2),
					  ticket_ultimo decimal(12,2),
					  loja_mais_frequente varchar(14),
					  loja_ultima_compra varchar(14)
					  );`

	sqlInsertData := `INSERT INTO data (cpf, private, incompleto, ultima_compra,
	                  ticket_medio, ticket_ultimo, loja_mais_frequente, loja_ultima_compra)
	                  VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`

	fileScanned, err := readLines("./base_teste.txt")
	if err != nil {
		log.Fatal(err)
	}

	finalData := splitData(fileScanned)

	_, err = db.Exec(sqlCreateTable)
	if err != nil {
		log.Fatal(err)
	}

	for i := range finalData {
		_, err := db.Exec(sqlInsertData, finalData[i].cpf, finalData[i].private, finalData[i].incompleto, finalData[i].ultCompra, finalData[i].ticketMedio, finalData[i].ticketUltimo, finalData[i].lojaMaisFreq, finalData[i].lojaUltCompra)
		if len(finalData[i].cpf) > 11 {
			fmt.Println(finalData[i].cpf)
		}
		if err != nil {
			log.Fatal(err)
		}
	}

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return indexHandler(c, db)
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return postHandler(c, db)
	})

	port := "8080"

	app.Static("/", "./public")
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
