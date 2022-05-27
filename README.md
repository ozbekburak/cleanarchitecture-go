# Clean Architecture - MongoDB & In-Memory

A simple RESTful Api application. 

An application with three endpoints, first one is (**/record**) that provide the ability to filter data with MongoDB connection, 
and other two (**/inmemory**, **/inmemory/set**) where you can perform set and get operations on key-value pairs by making an 
in-memory database connection.

The goal of the project is to develop resilient APIs that can easily be added to new endpoints, 
have less external dependencies, can be tested and at the same time have good error handling and a good log mechanism.


### Installation and Running the Application 

> PS: "Don't forget to set your .env variables. You can find sample .env in the root of the project"

### Docker

```bash
    docker build -t inmem-mongo
    docker run -p 8080:8080 inmem-mongo
```

### Local

```go
    go run main.go
```

```go
    go build -o server
    ./server
```


### Endpoints

#### Endpoint used to filter records in mongo database

```js
    POST /record
```

body:

```json
{
	"startDate": "2016-01-26",
	"endDate": "2018-02-02",
	"minCount": 6000,
	"maxCount": 6100
}
```

sample response: (we use 0 for success, 1 for others)

```json
{
	"code": 0,
	"msg": "success",
	"records": [
		{
			"key": "zqvnHclR",
			"createdAt": "2016-08-12T01:57:40.876Z",
			"totalCount": 6087
		},
		{
			"key": "zqvnHclR",
			"createdAt": "2016-08-12T01:57:40.876Z",
			"totalCount": 6087
		}
	]
}
```

#### Endpoint used to set key-value pairs in in-memory database

```js
    POST /inmemory/set
```

```json
{
	"key": "ping",
	"value": "pong"
}
```


sample response:

```json
{
	"key": "ping",
	"value": "pong"
}
```

#### Endpoint used to get key-value pairs in in-memory database

```js
    GET /inmemory?key=ping
```

sample response:

```json
{
	"key": "ping",
	"value": "pong"
}
```