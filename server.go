package main

import (
	"database/sql"
	"fmt"
	"log"
	"runtime"

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

/* estrutura de dados que será
inserida no banco de dados */
type Row struct {
	Cpf           string
	Private       string
	Incompleto    string
	UltCompra     string
	TicketMedio   string
	TicketUltimo  string
	LojaMaisFreq  string
	LojaUltCompra string
}

var docInvalido []string

func indexHandler(c *fiber.Ctx, db *sql.DB) error {
	var rowData Row
	var data []Row
	rows, err := db.Query("Select * from data")
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
		c.JSON("An error ocurred.")
	}
	for rows.Next() {
		rows.Scan(&rowData.Cpf, &rowData.Private, &rowData.Incompleto, &rowData.UltCompra, &rowData.TicketMedio, &rowData.TicketUltimo, &rowData.LojaMaisFreq, &rowData.LojaUltCompra)
		data = append(data, rowData)
	}
	return c.Render("index", fiber.Map{
		"Data": data,
	})
}

func cpfHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.Render("cpf", fiber.Map{
		"CPF": docInvalido,
	})
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%d sslmode=%s", user, dbname, pass, host, dbport, sslmode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// comando para criar tabela do banco de dados
	sqlCreateTable := `CREATE TABLE IF NOT EXISTS data (
					  cpf varchar(20) NOT NULL,
					  private bool NOT NULL,
					  incompleto bool NOT NULL,
					  ultima_compra DATE,
					  ticket_medio decimal(12,2),
					  ticket_ultimo decimal(12,2),
					  loja_mais_frequente varchar(20),
					  loja_ultima_compra varchar(20)
					  );`

	// comando de inserção de dados no DB
	sqlInsertData := `INSERT INTO data (cpf, private, incompleto, ultima_compra,
	                  ticket_medio, ticket_ultimo, loja_mais_frequente, loja_ultima_compra)
	                  VALUES ($1, $2, $3, NULLIF($4,'NULL')::date, $5, $6, $7, $8);`

	fileScanned, err := readLines("./base_teste.txt")
	if err != nil {
		log.Fatal(err)
	}

	finalData := splitData(fileScanned)

	// criação da tabela
	_, err = db.Exec(sqlCreateTable)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nAguarde. Inserindo dados no banco de dados...\n")
	// inserção de dados no DB
	for i := range finalData {
		_, err := db.Exec(sqlInsertData, finalData[i].Cpf, finalData[i].Private, finalData[i].Incompleto, finalData[i].UltCompra, finalData[i].TicketMedio, finalData[i].TicketUltimo, finalData[i].LojaMaisFreq, finalData[i].LojaUltCompra)
		if err != nil {
			log.Println(err)
		}
	}
	fmt.Printf("Pronto! Abra o navegador e digite localhost:8080\n\n")

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return indexHandler(c, db)
	})

	app.Get("/cpf", func(c *fiber.Ctx) error {
		return cpfHandler(c, db)
	})

	port := "8080"

	app.Static("/", "./views")
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
