
# Teacher Services API

This is a backend REST api that allows teachers to perform serveral requests.

## Setup
Ensure you have [Go](https://go.dev/) and [MySQL](https://www.mysql.com/) installed.

To set up the project locally:
1. Clone the repository using `git clone https://github.com/Gnoznaug/Govtech-take-home-assignment.git`
2. Run `go get` to install all dependencies
3. Create a .env file with `MYSQLPASSWORD={password for user 'test'}`
4. Set up your MySQL server with the user being `test` and a database called `takehome`
5. Set up the database in this format

![Table format](https://i.imgur.com/L2ClBMe.png)

## Usage
To start the server run:
```
cd Govtech-take-home-assignment/src
go run main.go
```

To run tests, in the `src` directory, run:
```
go test ./...
```
