# GoCalcAPI - A Calculator Web API Written in Go
A simple yet robust HTTP service that provides basic integer arithmetic operations through a RESTful API. Built with Go's standard library, featuring structured logging and rate limiting.

## Features
- Basic arithmetic operations (addition, subtraction, multiplication, division)
- JSON request/response format
- Input validation
- Structured JSON logging
- Rate limiting for all endpoints

## Endpoints
| Endpoint | Method | Description |
| --- | --- | --- |
| `/add` | POST | Add two integers |
| `/subtract` | POST | Subtract two integers |
| `/multiply` | POST | Multiply two integers |
| `/divide` | POST | Divide two integers |

## Usage
1. Clone the repository
```bash
git clone https://github.com/yingtu35/GoCalcAPI.git
cd GoCalcAPI
```

2. Run the server directly or build the binary
```bash
go run cmd/server/main.go
```
or
```bash
go build -o GoCalcAPI cmd/server/main.go
./GoCalcAPI
```

3. Send a request
```bash
curl -X POST http://localhost:8080/add -d '{"a": 1, "b": 2}'
```

## Roadmap
- [ ] Add suuport for floating-point numbers
- [ ] Add support for user authentication

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.