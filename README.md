## Documentação da API

A especificação da nossa API está disponível no formato OpenAPI. Você pode visualizar e interagir com a documentação da API utilizando ferramentas como Swagger UI, Postman, ou qualquer outra que suporte OpenAPI.

### Arquivo OpenAPI

O arquivo de especificação OpenAPI pode ser encontrado em [`docs/openapi.yaml`](docs/openapi.yaml).

### Visualizando a Documentação

Para visualizar a documentação da API, você pode utilizar:

- [Swagger Editor](https://editor.swagger.io/): Carregue o arquivo `openapi.yaml` para visualizar e interagir com a API.
- [Postman](https://www.postman.com/): Importe o arquivo `openapi.yaml` para criar coleções de testes para a API.

### Endpoints

- `POST /process-payment`: Processa um pagamento.
- `GET /payment-status`: Obtém o status de um pagamento.
- `POST /convert-currency`: Converte moeda.

Veja a especificação completa no arquivo [openapi.yaml](docs/openapi.yaml).
