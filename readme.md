# Golang CEP

Welcome to the Golang CEP Study Project! This project is designed to help you learn and practice your Golang (Go) programming skills by implementing a simple web application that retrieves and displays details for a given Brazilian CEP (postal code).

## Table of Contents

- [Introduction](#introduction)
- [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
- [Usage](#usage)
- [Routes](#routes)
- [Contributing](#contributing)
- [License](#license)

## Introduction

This project focuses on creating a web application using the Go programming language to retrieve and display details about a Brazilian CEP (postal code). The application provides a basic understanding of handling HTTP requests and serving HTML pages using the Go standard library.

## Getting Started

### Prerequisites

Before you begin, ensure you have the following prerequisites installed on your system:

- Go (1.13 or higher)
- Git

### Installation

1. Clone the repository:

```bash
git clone https://github.com/pmenta/golang-cep
```

2. Change to the project directory:

```bash
cd golang-cep
```

3. Build the project:

```bash
go build
```

## Usage

To run the application, execute the compiled binary:

```bash
./golang-cep
```

By default, the application will start and listen on port 8080. You can access the application in your web browser by navigating to `http://localhost:8080`.

## Routes

The application exposes the following route:

- `GET /cep/:cep` - Retrieves details for the specified Brazilian CEP and displays them as an HTML page.

## Contributing

Contributions are welcome! If you'd like to contribute to this project, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix: `git checkout -b feature-name`.
3. Make your changes and commit them: `git commit -m "Add feature"`.
4. Push your changes to your fork: `git push origin feature-name`.
5. Create a pull request detailing your changes.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---
