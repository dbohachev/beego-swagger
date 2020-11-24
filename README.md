# beego-swagger
[![Go Report Card](https://goreportcard.com/badge/github.com/dbohachev/beego-swagger)](https://goreportcard.com/report/github.com/dbohachev/beego-swagger)

Beego wrapper to automatically serve RESTful API documentation, created with Swagger 2.0.
This is simple implementation based on original http swagger: https://github.com/swaggo/http-swagger, 
changed a bit for Beego web framework.
## Usage

### Start using it
1. Add comments to your API source code, [See Declarative Comments Format](https://github.com/swaggo/swag#declarative-comments-format).
2. Download [Swag](https://github.com/swaggo/swag) for Go by using:

```sh
$ go get github.com/swaggo/swag/cmd/swag
```

3. Run the [Swag](https://github.com/swaggo/swag) in your Go project root folder which contains `main.go` file, [Swag](https://github.com/swaggo/swag) will parse comments and generate required files(`docs` folder and `docs/doc.go`).
```sh
$ swag init
```
4.Download [beego-swagger](https://github.com/dbohachev/beego-swagger) by using:
```sh
$ go get -u github.com/dbohachev/beego-swagger
```
And import following in your code:

```go
import "github.com/dbohachev/beego-swagger" // beego-swagger middleware
```

See "example" folder for example "classic" Beego API application.

![swagger_index.html](https://user-images.githubusercontent.com/51129427/73311427-0ef42980-422f-11ea-8c85-32dd301c39f5.png)
