# CONSERVICE
## Projeto de teste e aprendizado

Aplicativo que lê o arquivo base_teste.txt, trata os dados, insere em um banco de dados Postgresql e sobe um servidor para expor os dados em uma página de internet.

## Instalação

- Clone o repositório
- Certifique-se de que não há nenhum serviço do Postgresql rodando na sua máquina.
- `docker-compose up --build`
- Aguarde o banco de dados ser populado
- Abra o navegador e digite localhost:8080
