# CTI Group - Skill Test Back End Developer

- Name: Rahman Insani
- Email: rahmaninsani54@gmail.com
- Position: Back End Developer

## Project Structure

```
.
├── app
│   ├── database_app.go
│   └── http_router_app.go
├── config
│   ├── subcommand
│   │   ├── serve.go
│   │   └── subcommand.go
│   └── config.go
├── exception
│   ├── error_exception_handler.go
│   └── not_found_exception.go
├── handler
│   ├── security_case_handler.go
│   └── security_case_handler_impl.go
├── helper
│   ├── error_helper.go
│   ├── http_helper.go
│   ├── response_helper.go
│   └── string_helper.go
├── model
│   ├── domain
│   │   └── security_case.go
│   └── web
│       ├── response.go
│       ├── security_case_request.go
│       └── security_case_response.go
├── repository
│   ├── security_case_repository.go
│   └── security_case_repository_impl.go
├── route
│   └── security_case_route.go
├── test
│   └── security_case_handler_test.go
├── usecase
│   ├── security_case_usecase.go
│   └── security_case_usecase_impl.go
├── README.md
├── data.json
├── data_test.json
├── go.mod
└── main.go
```

## Technology Stack
- Go v1.21.6

## How to Run Application
1. Clone the repository
2. Build the application using the following command:
```sh
go build -o server
```
3. Show the help command using the following command:
```sh
./server --help
```
4. Run the application using the following command:
```sh
./server serve
```
5. Open Postman and import the collection from the following link:
https://documenter.getpostman.com/view/9933041/2s9Yyzbd3F


## API Documentation

| Method | Endpoint                | Description               |
|--------|-------------------------|---------------------------|
| POST   | /api/security-cases     | Create security case      |
| PUT    | /api/security-cases/:id | Update security case      |
| DELETE | /api/security-cases/:id | Delete security case      |
| GET    | /api/security-cases     | Get all security cases    |
| GET    | /api/security-cases/:id | Get security case by id   |

## How to Run Test
Run the test using the following command:
```sh
go test -v ./test
```
Result:
![Test Result](./result/test-result.png?raw=true "Test Result")

## Result
1. Create Security Case
   ![Create Security Case](./result/create-result.png?raw=true "Create Security Case")
2. Find All Security Cases
   ![Find All Security Cases](./result/find-all-result.png?raw=true "Find All Security Cases")
3. Update Security Case
   - Params
   ![Update Security Case - Params](./result/update-params.png?raw=true "Update Security Case - Params")
   - Result
   ![Update Security Case](./result/update-result.png?raw=true "Update Security Case")
4. Find By ID Security Case
   ![Find By Id Security Case](./result/find-by-id-result.png?raw=true "Find By Id Security Case")
5. Delete Security Case
   - Result
   ![Delete Security Case](./result/delete-result.png?raw=true "Delete Security Case")
   - Find All Security Cases After Delete
   ![Find All Security Cases After Delete](./result/find-all-after-delete.png?raw=true "Find All Security Cases After Delete")
   - Find By ID Security Case After Delete
   ![Find By Id Security Case After Delete](./result/find-by-id-after-delete.png?raw=true "Find By Id Security Case After Delete")
   - Update Security Case After Delete
   ![Update Security Case After Delete](./result/update-after-delete.png?raw=true "Update Security Case After Delete")
   - Delete Security Case After Delete
   ![Delete Security Case After Delete](./result/delete-after-delete.png?raw=true "Delete Security Case After Delete")