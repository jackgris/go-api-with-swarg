# go-api-with-swarg is:
Simple repo for writing an example of documenting an API with Swagger in Go

For running the server in the root path of the repository:

- First run: `swag init`
- Second run: `go run ./main.go`


## Calling the API using Curl.

Thanks to swagger you don't need to think in all these commands that you'll see below. Because you can see them in the docs endpoint: `http://localhost:3000/swagger`

Get all element from Todo:
```bash
curl -X 'GET' \
  'http://localhost:3000/todo' \
  -H 'accept: application/json'
```

Get one element from Todo using Curl:
```bash
curl -X 'GET' \
  'http://localhost:3000/todo/2' \
  -H 'accept: application/json'
```

Create Todo example using Curl:
```bash
curl -X 'POST' \
  'http://localhost:3000/todo' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "id": "5",
  "task": "new item"
}'
```

Delete Todo example using Curl:
```bash
curl -X 'DELETE' \
  'http://localhost:3000/todo/5' \
  -H 'accept: application/json'
```

And remember every time you update the documentation you need to run `swag init` command. For more details, I'll write a post on my blog.
