# How to structure Golang applications
This is the repository for my article on how to structure Golang applications. It can be found [here]().

The purpose is to show newcomers to Go how structure their applications.

## How to run the application
You don't need much the run this application. It has the following installed on your computer:
- Golang 1.15+
- PostgreSQL

Clone the repository and change the connectionString variable on line 26 in main.go to the following:
"postgres://postgres:postgres@localhost/**NAME-OF-YOUR-DATABASE-HERE**?sslmode=disable"

After that, just run go run cmd/server/main.go from your terminal

## License
MIT
