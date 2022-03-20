package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/paemuri/brdoc"
)

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
				if !brdoc.IsCPF(v2) {
					fmt.Println(v2, "invalido")
				}
				parsedData[newIndex].cpf = cleanStrings(v2)
			case 1:
				parsedData[newIndex].private = strToBool(v2)
			case 2:
				parsedData[newIndex].incompleto = strToBool(v2)
			case 3:
				parsedData[newIndex].ultCompra = strToDate(v2)
			case 4:
				parsedData[newIndex].ticketMedio = strToFloat(v2)
			case 5:
				parsedData[newIndex].ticketUltimo = strToFloat(v2)
			case 6:
				parsedData[newIndex].lojaMaisFreq = cleanStrings(v2)
			case 7:
				parsedData[newIndex].lojaUltCompra = cleanStrings(v2)
			}
		}
	}

	return parsedData
}

func strToDate(value string) string {
	if value == "NULL" {
		value = "1900-01-01"
	}
	return value
}

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

func strToBool(value string) string {
	if value == "1" {
		return "true"
	} else {
		return "false"
	}
}

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
