# Company Service

[![Go Report Card](https://goreportcard.com/badge/github.com/patricksferraz/company-service)](https://goreportcard.com/report/github.com/patricksferraz/company-service)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![GoDoc](https://godoc.org/github.com/patricksferraz/company-service?status.svg)](https://godoc.org/github.com/patricksferraz/company-service)

A modern, scalable microservice for managing company data, built with Go and following clean architecture principles.

## 🚀 Features

- REST and gRPC APIs for company and employee management
- Event-driven architecture using Apache Kafka
- PostgreSQL database integration
- Docker and Docker Compose support
- Swagger API documentation
- Elastic APM monitoring
- Clean Architecture implementation
- Environment-based configuration
- Comprehensive test coverage

## 🏗 Architecture

This service follows Clean Architecture principles, organized into the following layers:

- **Domain**: Core business logic and entities
- **Application**: Use cases and business rules
- **Infrastructure**: External services, databases, and frameworks
- **Interface**: API handlers and controllers

## 🛠 Tech Stack

- **Language**: Go 1.16+
- **Framework**: Gin (REST) and gRPC
- **Database**: PostgreSQL
- **Message Broker**: Apache Kafka
- **Container**: Docker
- **Monitoring**: Elastic APM
- **Documentation**: Swagger
- **Testing**: Go testing framework

## 📋 Prerequisites

- Go 1.16 or higher
- Docker and Docker Compose
- Make (optional, for using Makefile commands)

## 🚀 Getting Started

1. Clone the repository:
```bash
git clone https://github.com/patricksferraz/company-service.git
cd company-service
```

2. Copy the environment file and configure it:
```bash
cp .env.example .env
```

3. Start the services using Docker Compose:
```bash
docker-compose up -d
```

4. The service will be available at:
   - REST API: http://localhost:8080
   - gRPC: localhost:50051
   - Swagger UI: http://localhost:8080/swagger/index.html
   - Adminer (Database UI): http://localhost:9000
   - Kafka Control Center: http://localhost:9021

## 🧪 Running Tests

```bash
go test ./...
```

## 📚 API Documentation

The API documentation is available through Swagger UI at `http://localhost:8080/swagger/index.html` when the service is running.

## 🔧 Configuration

The service can be configured through environment variables. See `.env.example` for all available options.

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## 👥 Authors

- **Patrick Ferraz** - *Initial work* - [patricksferraz](https://github.com/patricksferraz)

## 🙏 Acknowledgments

- Thanks to all contributors who have helped shape this project
- Inspired by clean architecture principles and best practices in Go development
