# Go-Noting

Go-Noting is a simple, easy-to-use, noting web application.

## Stacks

The server side is built with Go to handle operations including CRUD and migraiton. The primary packages that are used in the server side are:

- Echo for the web framework.
- Cobra to build the CLI.
- Gorm ORM library to work with database.

The client side is built with Vue3 with many libraries to make the noting application more interactive.

## Deploy Go Binary (Server)

Compile the project:

```bash
go build -o go-noting main.go # linux
go build -o go-noting.exe main.go # Windows
```

Available Commands on the binary:

```bash
./go-noting serve # serve the applicaiton

./go-noting migrate     # Automatically migrate the database
./go-noting migrate -f  # Reset database, then run migration
```

## Contribute

Feel free to contribute and make a pull request :)

## License

This project is under [MIT LICENSE](LICENSE)
