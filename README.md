# go-rest-api
Basic REST API example using golang with code splitting

We just need to install go. Follow the link for installation.

[Install GoLang](https://golang.org/doc/install)

Once the installation is done execute following command to start Go Rest API server.
As this is an example with the code splitting we need to execute the following command to execute the code.

```shell
go run *.go
```

The command will start server on port 7777.

Hit one of the API end points to get API response

1. GET - http://localhost:7777/api/v1/books (Get all Books)
2. GET - http://localhost:7777/api/v1/books/:id (Get Book by id)
3. POST - http://localhost:7777/api/v1/books (Add new Book)
4. PUT - http://localhost:7777/api/v1/books/:id (Update Book by id)
5. DELETE - http://localhost:7777/api/v1/books/:id (Delete Book by id)

