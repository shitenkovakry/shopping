# Instructions of Shopping task

### How to use?

```sh
GOOS=linux GOARCH=amd64 go build -o ./build/app/myapp ./server/main.go
docker-compose up -d
```

then wait for the docker infra to launch.

### Tests

In order to get lists of products(admin), do this:

```
curl -i -X GET \
  http://localhost:8080/list/products/admin
  ```

In order to get lists of products which have published status(public), do this:

```
curl -i -X GET \
  http://localhost:8080/list/products/public
```
