# gestaoClientes

Comandos:
Iniciar o banco de dados: docker-compose up
Rodar testes: go test
Iniciar a aplicação: go run main.go

Instalar Swagger: go install github.com/swaggo/swag/cmd/swag@latest
Iniciar Swagger: swag init --parseDependency --parseInternal --parseDepth 1

Verificar o que entrou nos testes: go test -coverpkg=./... -coverprofile=coverage.out && go tool cover -html=coverage.out

![Swagger](https://user-images.githubusercontent.com/53271581/208182134-7a153f0f-c3fb-44e1-8a20-4d13cbada19a.png)
