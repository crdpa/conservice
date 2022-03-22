package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
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
func splitData(data []string) []row {
	parsedData := make([]row, len(data)-1)
	for i1, v1 := range data {
		if i1 == 0 {
			continue
		}
		lines := strings.Fields(v1)
		for i2, v2 := range lines {
			newIndex := i1 - 1
			switch i2 {
			case 0:
				if !contains(cpfInvalido, v2) {
					if !brdoc.IsCPF(v2) {
						cpfInvalido = append(cpfInvalido, v2)
					}
				}
				parsedData[newIndex].cpf = cleanStrings(v2)
			case 1:
				parsedData[newIndex].private = v2
			case 2:
				parsedData[newIndex].incompleto = v2
			case 3:
				parsedData[newIndex].ultCompra = v2
			case 4:
				parsedData[newIndex].ticketMedio = strToFloat(v2)
			case 5:
				parsedData[newIndex].ticketUltimo = strToFloat(v2)
			case 6:
				if v2 != "NULL" && !contains(cpfInvalido, v2) {
					if !brdoc.IsCNPJ(v2) {
						cpfInvalido = append(cpfInvalido, v2)
					}
				}
				parsedData[newIndex].lojaMaisFreq = cleanStrings(v2)
			case 7:
				if v2 != "NULL" && !contains(cpfInvalido, v2) {
					if !brdoc.IsCNPJ(v2) {
						cpfInvalido = append(cpfInvalido, v2)
					}
				}
				parsedData[newIndex].lojaUltCompra = cleanStrings(v2)
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

/* converte valor para float pois o campo
do banco de dados utiliza duas casas decimais */
func strToFloat(value string) float64 {
	if value == "NULL" {
		value = "0"
	}
	value = strings.Replace(value, ",", ".", 1)
	toFloat, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Fatal(err)
	}

	return toFloat
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
