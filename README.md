# Student API in Go 

REST API built with Go and Gin framework.
Containerized with Docker. CI/CD with GitHub Actions.

## Tech Stack
- Go + Gin
- Docker (Multi-stage build)
- GitHub Actions

## Run Without Docker
```bash
go run main.go
```

## Run With Docker
```bash
docker build -t student-api-go .
docker run -p 8080:8080 student-api-go
```

## API Endpoints
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/students` | Get all students |
| GET | `/api/students/:id` | Get by ID |
| POST | `/api/students` | Add student |
| DELETE | `/api/students/:id` | Delete student |

## Author
Prachi | 2nd Year CS | Learning DevOps 
