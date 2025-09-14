# Library-API

Uma API RESTful para gerenciamento de livros, autores e empréstimos em uma livraria, escrita em Go com Gin e GORM.

## ✨ Características

* 📚 CRUD completo para livros, autores e empréstimos
* 🔗 Relacionamentos entre livros e autores (many2many)
* 🏦 Controle de disponibilidade de livros
* ⚡ Rápida e leve, utilizando SQLite
* 🌐 Rotas documentadas com Swagger
* 🛠️ Middleware de logging e CORS configurado

## 🚀 Instalação

### Opção 1: Build local

```bash
git clone https://github.com/seu-usuario/library-api
cd library-api
go build -o library-api ./cmd/server
```

### Opção 2: Executar diretamente

```bash
git clone https://github.com/seu-usuario/library-api
cd library-api
go run ./cmd/server/main.go
```

### Opção 3: Instalar globalmente

```bash
go install github.com/seu-usuario/library-api/cmd/server@latest
```

## 📖 Uso

### Rodar a API localmente

```bash
./library-api
```

* API disponível em: `http://localhost:8080`
* Swagger UI em: `http://localhost:8080/swagger/index.html`

## ⚙️ Endpoints

| Método | Rota               | Descrição                       |
| ------ | ------------------ | ------------------------------- |
| GET    | /books             | Lista todos os livros           |
| POST   | /books             | Cria um novo livro              |
| GET    | /books/{id}        | Busca livro pelo ID             |
| PUT    | /books/{id}        | Atualiza um livro               |
| DELETE | /books/{id}        | Remove um livro                 |
| GET    | /authors           | Lista todos os autores          |
| POST   | /authors           | Cria um novo autor              |
| GET    | /authors/{id}      | Busca autor pelo ID             |
| PUT    | /authors/{id}      | Atualiza um autor               |
| DELETE | /authors/{id}      | Remove um autor                 |
| GET    | /loans             | Lista todos os empréstimos      |
| POST   | /loans             | Cria um novo empréstimo         |
| GET    | /loans/{id}        | Busca empréstimo pelo ID        |
| PUT    | /loans/{id}/return | Marca empréstimo como devolvido |
| DELETE | /loans/{id}        | Remove um empréstimo            |

## 💡 Exemplos

### Criar um livro

```bash
curl -X POST http://localhost:8080/books \
-H "Content-Type: application/json" \
-d '{
  "title": "Livro Exemplo",
  "isbn": "123-456",
  "author_ids": [1,2]
}'
```

### Listar livros

```bash
curl http://localhost:8080/books
```

### Registrar empréstimo

```bash
curl -X POST http://localhost:8080/loans \
-H "Content-Type: application/json" \
-d '{
  "book_id": 1,
  "user_name": "João Silva"
}'
```

## 🔧 Build

### Build simples

```bash
go build -o library-api ./cmd/server
```

### Build otimizado (tamanho reduzido)

```bash
go build -ldflags="-s -w" -o library-api ./cmd/server
```

## 🧪 Teste

```bash
# Teste básico
curl http://localhost:8080/books

# Teste avançado com criação
curl -X POST http://localhost:8080/books \
-H "Content-Type: application/json" \
-d '{"title":"Novo Livro","isbn":"999-888"}'
```

## 📋 Requisitos

* Go 1.16 ou superior
* SQLite (já incluso via GORM)

## 🛠️ Desenvolvimento

```bash
# Clonar repositório
git clone https://github.com/seu-usuario/library-api
cd library-api

# Executar em modo desenvolvimento
go run ./cmd/server/main.go

# Gerar documentação Swagger
swag init -g cmd/server/main.go
```

## 📄 Licença

MIT License - veja o arquivo [LICENSE](LICENSE) para detalhes.

## 🤝 Contribuição

Contribuições são bem-vindas! Por favor, abra uma issue ou envie um pull request.

## 🚧 Próximas Funcionalidades

* [ ] Autenticação e autorização de usuários
* [ ] Filtros e paginação nos endpoints
* [ ] Testes unitários e de integração
* [ ] Validação de dados mais robusta
* [ ] Suporte a banco de dados PostgreSQL/MySQL
