Tech Stack

Go 1.22+

Chi – HTTP router and middleware

pgx/v5 – PostgreSQL driver

SQLC – Type-safe database queries

Goose (optional) – Database migrations

PostgreSQL – Main database

Project structure
.
├── cmd/
│   └── api/            # Main entrypoint
├── internal/
│   ├── adapters/
│   │   └── postgresql/
│   │       └── sqlc/   # SQLC generated code
│   ├── products/       # Product service + handler
│   ├── orders/         # Order service + handler
│   └── env/            # Env helper
├── db/
│   ├── migrations/     # SQL migrations
│   └── queries/        # SQLC queries
├── api.go              # HTTP server + routes
└── README.md
