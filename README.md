# Learning GO

![Test](https://github.com/bigmassa/gohttp-starter/workflows/Test/badge.svg)

A basic site with user authentication and email.

## Docker

Build and run the docker container

```bash
docker-compose build
docker-compose up
```

## Testing

run all tests

```bash
docker-compose exec app go test -v ./...
```

## Commands

Serve manually:
```bash
docker-compose exec app serve -listen :80
```

Create a new user:
```bash
docker-compose exec app createuser -email=admin@example.com -firstname=Admin -lastname=User -password=password
```

Migrate the database:
```bash
docker-compose exec app migrate
```

## Links

Docs

* [https://golang.org/doc](https://golang.org/doc/)
* [https://gorm.io/docs](https://gorm.io/docs/)
* [http://www.gorillatoolkit.org/pkg/mux](http://www.gorillatoolkit.org/pkg/mux)
* [http://www.gorillatoolkit.org/pkg/securecookie](http://www.gorillatoolkit.org/pkg/securecookie)
* [http://www.gorillatoolkit.org/pkg/sessions](http://www.gorillatoolkit.org/pkg/sessions)
* [https://github.com/dchest/passwordreset](https://github.com/dchest/passwordreset)
* [https://github.com/dchest/authcookie](https://github.com/dchest/authcookie)
* [https://godoc.org/github.com/go-playground/validator](https://godoc.org/github.com/go-playground/validator)

Examples

* [https://gobyexample.com](https://gobyexample.com/)
* [https://github.com/avelino/awesome-go](https://github.com/avelino/awesome-go)
