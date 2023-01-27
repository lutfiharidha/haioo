
<h1 align="center">Haioo Cart</h1>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![GitHub Issues](https://img.shields.io/github/issues-pr/kylelobo/The-Documentation-Compendium.svg)](https://github.com/lutfiharidha/haioo/issues)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/kylelobo/The-Documentation-Compendium.svg)](https://github.com/lutfiharidha/haioo/pulls)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE)

</div>

## 📝 Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Usage](#usage)
- [API Documentation](#doc)
- [Built Using](#built_using)
- [Authors](#authors)

## 🧐 About <a name = "about"></a>

Haioo Cart Service

## 🏁 Getting Started <a name = "getting_started"></a>

### Installing

A step by step series of examples that tell you how to get a development env running.

- Rename .env.example file to .env
- Configuration database in .env file
```
DB_HOST= //Database Host
DB_PORT= //Database Port
DB_NAME= //Database Name
DB_USERNAME= //Database Username
DB_PASSWORD= //Database Password
```
Don't worry, you don't need to create a table for this application because this application will automatically create a table by itself. Just make sure your database is connected correctly.

## 🔧 Running the test <a name = "tests"></a>

```
make testing
```

### Scenario Unit Test

The scenarios that have been created in the unit test include positive scenarios and negative scenarios, such as successfully requesting API endpoints and validating errors such as incorrect request data input, if the data product already in cart it will just update the quantity, and if an incoming request is not properly completed.

## 🎈 Usage <a name="usage"></a>

### Run without build
to run the application use the following command
```
make run
```
### Run with build

to build the application use the following command
```
make build
```

and run it with command
```
./bin/app
```
default servering HTTP on port 8081


## 📖 API Documentation <a name="doc"></a>

See the API documentation for this application

- [https://petstore.swagger.io/?url=https://raw.githubusercontent.com/lutfiharidha/haioo/master/docs/swagger.yaml](https://petstore.swagger.io/?url=https://raw.githubusercontent.com/lutfiharidha/haioo/master/docs/swagger.yaml)

## ⛏️ Built Using <a name = "built_using"></a>

- [Golang](https://go.dev/) - Server Environment
- [Gin](https://gin-gonic.com/) - Server Framework
- [MySQL](https://www.mysql.com/) - Database
- [Swagger](https://swagger.io/) - API Documentation

## ✍️ Authors <a name = "authors"></a>

- [@lutfiharidha](https://github.com/lutfiharidha) 

