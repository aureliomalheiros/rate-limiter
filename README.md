# Rate Limiter

```markdown
rate-limiter-go/
│
├── .env
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
├── middleware/
│   └── rate_limiter.go
├── limiter/
│   └── limiter.go
└── tests/
    └── rate_limiter_test.go
```

## Running

- The project using golang version 1.21
- All ENVs it's possibly modify in file `.env`
- You can request the address in URL: `http://localhost:8080`
- 