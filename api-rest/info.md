# API REST de Clientes

> Quarto desafio do curso de Go developer da Dio.

Esta API REST permite **cadastrar, listar, editar e excluir** clientes. Cada cliente tem um **ID gerado automaticamente** e um **endereço** estruturado.

### Exemplo de Cliente (Darth Vader)
```json
{
  "nome": "Darth",
  "sobrenome": "Vader",
  "telefone": "(99)9999999",
  "endereco": {
    "rua": "Death Star",
    "numero": 666,
    "complemento": "Sala do Imperador",
    "cidade": "Galáxia Distante",
    "estado": "Império"
  }
}
```

### Endpoints da API (Rodando na porta 8080)
Cadastrar um cliente:
```sh
curl -X POST http://localhost:8080/clientes -H "Content-Type: application/json" -d '{"nome":"Darth","sobrenome":"Vader","telefone":"999999999","endereco":{"rua":"Death Star","numero":666,"complemento":"Sala do Imperador","cidade":"Galáxia Distante","estado":"Império"}}'
```

Listar todos os clientes:
```sh
curl -X GET http://localhost:8080/clientes
```

Buscar um cliente pelo ID:
```sh
curl -X GET http://localhost:8080/clientes/ID_DO_CLIENTE
```
Editar um cliente pelo ID:
```sh
curl -X PUT http://localhost:8080/clientes/ID_DO_CLIENTE -H "Content-Type: application/json" -d '{"nome":"Darth","sobrenome":"Vader","telefone":"999999999","endereco":{"rua":"Death Star","numero":666,"complemento":"Sala do Imperador","cidade":"Galáxia Distante","estado":"Império"}}'
```
Excluir um cliente pelo ID:
```sh
curl -X DELETE http://localhost:8080/clientes/ID_DO_CLIENTE
```

### Como rodar a API
Basta somente iniciar o servidor com o comando:
```sh
go run main.go
```

Dica: para testar a API sem o `curl`, você pode usar o [Postman](https://www.postman.com/) ou o [Insomnia](https://insomnia.rest/).