# Entice
> Simple application deployment dummy

## Dependencies
Please use [dep](https://golang.github.io/dep/)

## Environment
```
DATABASE_ADDR
DATABASE_USER
DATABASE_PASSWORD
DATABASE_DATABASE
REDIS_ADDR
LISTEN
```

## Example request
**Request:**
```bash
curl http://localhost:8080
```

**Response:**
```json
{"db": "PostgreSQL 9.6.2 on x86_64-apple-darwin16.4.0, compiled by Apple LLVM version 8.0.0 (clang-800.0.42.1), 64-bit", "redis": "2018-03-20T20:58:43+03:00"}
```