# CONSERVICE
## Projeto de teste e aprendizado

Aplicativo que lê o arquivo base_teste.txt, trata os dados, insere em um banco de dados Postgresql e sobe um servidor para expor os dados em uma página de internet.

## Instalação

- Clone o repositório
- Certifique-se de que não há nenhum serviço do Postgresql rodando na sua máquina.
- `docker-compose up --build`
- Aguarde o banco de dados ser populado
- Abra o navegador e digite localhost:8080

---

## Learning and testing project

Application that reads base_teste.txt, split and insert the data in PostgreSQL. It runs a web server to expose the data in a html page.

## Install

- Clone
- Make sure there is no PostgreSQL instance running locally
- `docker-compose up --build`
- Wait for the database to be populated
- Open your browser and go to localhost:8080
