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

g
 In order to get published product, do this:

```
curl -i -X GET \
  http://localhost:8080/api/v1/get/product/1
  ```

In order to add product, do this:

```
curl -i -X POST \
 --data-binary '{"name_product": "potato","price_product": 30.0,"status_product": "published"}' \
  http://localhost:8080/api/v1/add/product
  ```

 In order to change price of product, do this:

```
curl -i -X PUT \
 --data-binary '{"id_product":1,"new_price": 35.0}' \
  http://localhost:8080/api/v1/change/price/product"
  ```

 In order to change name of product, do this:

```
curl -i -X PUT \
 --data-binary '{"id_product":1,"new_name": "salat"}' \
  http://localhost:8080/api/v1/change/name/product"
  ```

   In order to change status of product, do this:

```
curl -i -X PUT \
 --data-binary '{"id_product":1,"new_status": "unpublished"}' \
  http://localhost:8080/api/v1/change/status/product"
  ```


 In order to delete product, do this:

```
curl -i -X DELETE \
 --data-binary '{"id_product":1}' \
  http://localhost:8080/api/v1/delete/product"
  ```



In order to register buyer, do this:

```
curl -i -X POST \
 --data-binary '{"name_buyer": "Josh","email_buyer": "josh@mail.ru","balance_buyer": 500.50, "status_buyer": "active"}' \
  http://localhost:8080/api/v1/register/buyer
  ```

 In order to change name of buyer, do this:

```
curl -i -X PUT \
 --data-binary '{"id_buyer":1,"new_name": "Kenny"}' \
  http://localhost:8080/api/v1/change/name/buyer"
  ```

 In order to change email of buyer, do this:

```
curl -i -X PUT \
 --data-binary '{"id_buyer":1,"new_email": "kenny@ked.ru"}' \
  http://localhost:8080/api/v1/change/email/buyer"
  ```


 In order to change status of buyer, do this:

```
curl -i -X PUT \
 --data-binary '{"id_buyer":1,"new_status": "deleted"}' \
  http://localhost:8080/api/v1/change/statusg/buyer"
  ```

 In order to delete account of buyer, do this:

```
curl -i -X DELETE \
 --data-binary '{"id_buyer":1}' \
  http://localhost:8080/api/v1/delete/account"
  ```

 In order to replenish balance of buyer, do this:

```
curl -i -X PUT \
 --data-binary '{"id_buyer":1,"price_product": 35.0}' \
  http://localhost:8080/api/v1/replenish/balance/buyer"
  ```
