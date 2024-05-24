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
- The limits default is 5 request per second to IP
- I configured Redis, because I want persist the access data


> [!Note]
> When the limit is exceeded, the server return a 429 HTTP code with the message "You have rached the Maximum number of requests or actions allowed Within a certain time frame"

