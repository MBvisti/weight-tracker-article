# How to structure Golang applications

**2024 Update**:
Since I wrote this article lots have changed both in how I structure things but also the libraries available (like sqlc, templ) in the ecosystem. I still this has some value for newcomers as it discusses some code architectural concepts that are healthy to know. However, I will write a follow-up soon that discusses the layout and reasons, based on my project [Grafto](https://github.com/mbv-labs/grafto).

This is the repository for my article on how to structure Golang applications. It can be found [here](https://mortenvistisen.com/posts/practical-approach-to-structuring-go-apps).

The purpose is to show newcomers to Go how structure their applications.

If you want a comprehensive course on building fullstack web apps with go, check [Golang Blog Course](https://golangblogcourse.com?utm_source=github&utm_campaign=structure-golang-applications).

## How to run the application
You don't need much the run this application. It has the following installed on your computer:
- Golang 1.15+
- PostgreSQL

Clone the repository and change the connectionString variable on line 26 in main.go to the following:
"postgres://postgres:postgres@localhost/**NAME-OF-YOUR-DATABASE-HERE**?sslmode=disable"

After that, just run go run cmd/server/main.go from your terminal

## License
MIT
