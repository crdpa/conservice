package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/paemuri/brdoc"
)

// lê o arquivo de texto e coloca em um slice de strings
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

/* separa os dados e converte para os tipos
 * compatíveis com o banco de dados */
func splitData(data []string) []Row {
	parsedData := make([]Row, len(data)-1)
	for i1, v1 := range data {
		if i1 == 0 {
			continue
		}
		lines := strings.Fields(v1)
		for i2, v2 := range lines {
			newIndex := i1 - 1
			switch i2 {
			case 0:
				if !contains(docInvalido, v2) {
					if !brdoc.IsCPF(v2) {
						docInvalido = append(docInvalido, v2)
					}
				}
				parsedData[newIndex].Cpf = v2
			case 1:
				parsedData[newIndex].Private = v2
			case 2:
				parsedData[newIndex].Incompleto = v2
			case 3:
				parsedData[newIndex].UltCompra = v2
			case 4:
				parsedData[newIndex].TicketMedio = commaToPeriod(v2)
			case 5:
				parsedData[newIndex].TicketUltimo = commaToPeriod(v2)
			case 6:
				if v2 != "NULL" && !contains(docInvalido, v2) {
					if !brdoc.IsCNPJ(v2) {
						docInvalido = append(docInvalido, v2)
					}
				}
				parsedData[newIndex].LojaMaisFreq = v2
			case 7:
				if v2 != "NULL" && !contains(docInvalido, v2) {
					if !brdoc.IsCNPJ(v2) {
						docInvalido = append(docInvalido, v2)
					}
				}
				parsedData[newIndex].LojaUltCompra = v2
			}
		}
	}

	return parsedData
}

// retira pontos, barras e traços dos CPF/CPNJs
func cleanStrings(value string) string {
	if value != "NULL" {
		re, err := regexp.Compile(`[\\\-\.//]`)
		if err != nil {
			log.Fatal(err)
		}

		cleanString := re.ReplaceAllString(value, "")

		return cleanString
	}

	return value
}

// substitui vírgula por ponto
func commaToPeriod(value string) string {
	if value == "NULL" {
		value = "0"
	}
	value = strings.Replace(value, ",", ".", 1)

	return value
}

/* checa se o elemento ja está contido no slice
 * para evitar CPFs/CNPJs inválidos duplicados */
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
