# go-itertools

Go package that contains methods that are helpful for working with iterators.

## Development

This project contains Docker Compose services that can be used for development.

- The `dev` service can be used to run the `go` executable:
  ```
  $ docker compose up -d dev
  $ docker compose exec dev bash
  go-user@...:/wd$ go test ./itertools
  PASS
  ok      github.com/dpbricher/go-itertools/itertools     0.006s
  ```
- The `doc` service can be used to preview the package docs
  - start the service with `docker compose up doc`
  - visit `localhost:8080` in a web browser
