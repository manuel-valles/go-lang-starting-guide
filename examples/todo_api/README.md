# REST API with Go

## Setup

- Set up dependency tracking inside the Go app running: `$ go mod init example/todo_api`

- Install [Git Web Framework](https://github.com/gin-gonic/gin): `$ go get github.com/gin-gonic/gin`

## API Concepts

Since the communication between `Client` and `Server` is through JSON, we need to convert the `struct` into a `json` format:

```go
type todo struct {
	ID 			string 	`json:"id"`
	Item 		string 	`json:"item"`
	Completed 	bool	`json:"completed"`
}
```

- To get the response from the server: `context.IndentedJSON(http.StatusOK, todos)`
- To get data from the request: `context.BindJSON(&todo)`
