### Quick-start-gin

Este repo é um quick start criado por mim para inicializar projetos de API de forma mais rapida sem precisar perder tanto tempo com config, baseado em projetos criados nas lives do meu canal no youtube [devnych](https://youtube.com/@devnych)

(em breve - autenticação + authorization com casbin, ratelimit, pacotes úteis)

🧰 Tecnologias principais do quick-start

- Go (Golang) — API principal
- Gin — Framework HTTP
- SQLC — Geração de queries SQL tipadas
- PostgreSQL — Banco de dados principal
- Redis — Cache e filas
- Docker / Docker Compose — Ambientes locais
- MinIO — Armazenamento de arquivos local
- Mailhog - ferrementa de email local
- Makefile — Automação de comandos

⚙️ Pré-requisitos

Dependências e clis necessárias:

- Go ≥ 1.25.4
- Docker + Docker Compose
- Make
- SQLC
- atlas
- air
- swagger-go

🚀 Setup local

```
# 1. Clone o projeto
git clone https://github.com/nycholasmarques/quick-start-gin.git
# gh repo clone nycholasmarques/quick-start-gin.git
cd quick-start-gin

# 2. Subir os serviços locais
make up (docker-compose) ou docker compose up -d (são diferentes)

# 3. Rodar as migrations
make migrate

# 4. Iniciar o servidor
make dev
```

📦 Tabela de Serviços

| Serviço               | Porta Exposta | Interface / Endpoint                                                                               | Descrição                                                        |
| --------------------- | ------------- | -------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------- |
| **PostgreSQL (main)** | `5432`        | —                                                                                                  | Banco de dados principal da aplicação (`postgres_main`)          |
| **PostgreSQL (dev)**  | `5433`        | —                                                                                                  | Banco para dev e testes (`postgres_dev`  necessário pro atlas)   |
| **Redis**             | `6379`        | —                                                                                                  | Cache, rate-limit, locks e armazenamento rápido                  |
| **MinIO API**         | `9000`        | [http://localhost:9000](http://localhost:9000)                                                     | Armazenamento S3-like para uploads                               |
| **MinIO Console**     | `9001`        | [http://localhost:9001](http://localhost:9001)                                                     | Painel administrativo do MinIO                                   |
| **MailHog (SMTP)**    | `1025`        | smtp://localhost:1025                                                                              | Servidor SMTP fake para testes de e-mail                         |
| **MailHog Web UI**    | `8025`        | [http://localhost:8025](http://localhost:8025)                                                     | Interface para ver todos os e-mails enviados                     |
| **Adminer**           | `8081`        | [http://localhost:8081](http://localhost:8081)                                                     | UI Web para acessar PostgreSQL                                   |
| **API quick-start-gin (Go)** | `8080`        | [http://localhost:8080/api/v1/](http://localhost:8080/api/v1)                                      | API principal da aplicação                                       |
| **Swagger**           | —             | [http://localhost:8080/api/v1/swagger/index.html](http://localhost:8080/api/v1/swagger/index.html) | Documentação interativa da API                                   |

🧱 Migrations (Atlas)

| Comando               | Descrição                                |
| --------------------- | ---------------------------------------- |
| `make migrate.diff`   | cria uma nova migration com base no diff |
| `make migrate`        | aplica migrations pendentes              |
| `make migrate.status` | exibe status das migrations              |
| `make migrate.lint`   | lint das migrations                      |
| `make migrate.hash`   | recalcula/valida o atlas.sum             |

🧰 Utilidades
| Comando     | Descrição                  |
| ----------- | -------------------------- |
| `make fmt`  | formata código Go          |
| `make docs` | gera documentação Swagger  |
| `make test` | executa tests com coverage |

📌 Observação
- Verificar envs necessárias em env.example.
- As queries SQL devem ser adicionadas em internal/database/queries com base nos módulos, ex: `users.sql`, `gyms.sql`...
- As migrations são geradas automaticamente via `make migrate.diff` pós mudar algo no `schema.sql`, depois de gerar as migrations, você pode subir as alterações via `make migrate`, também é possível ver o status das migrações aplicadas e peendentes via `make migrate.status`
- Todo código gerado pelo `sqlc generate` fica em internal/database/sqlc. (confira o sqlc.yaml se necessário)
