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
  http://localhost:8080/api/v1/list/products/admin
  ```

In order to get lists of products which have published status(public), do this:

```
curl -i -X GET \
  http://localhost:8080/api/v1/list/products/public
```

In order to get product, do this:

```
curl -i -X GET \
  http://localhost:8080/api/v1/get/product/{1}/admin
  ```

  In order to get published product, do this:

```
curl -i -X GET \
  http://localhost:8080/api/v1/get/product/{id_product}
  ```

In order to add product, do this:

```
curl -i -X POST \
 --data-binary '{"name_product": "potato","price_product": 30.0,"status": "published"}' \
  http://localhost:8080/api/v1/add/product
  ```

 In order to add product, do this:

```
curl -i -X PUT \
 --data-binary '{"id_product":1,"new_price": 35.0}' \
  http://localhost:8080/api/v1/change/price/product/1"
  ```
