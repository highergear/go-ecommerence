# go-ecommerence

## Requirements
- MYSQL
- Golang
- Run the following command to get the project dependencies: 
```
// mysql driver
go get -u github.com/go-sql-driver/mysql

// http framework
github.com/gin-gonic/gin

// JWT
go get -u github.com/golang-jwt/jwt

// environment variables
go get github.com/spf13/viper

// ORM
go get -u github.com/jinzhu/gorm

// Password hashing
go get -u golang.org/x/crypto

// unit test
go get github.com/stretchr/testify

// mock sql for unit test
go get github.com/DATA-DOG/go-sqlmock
```


## API Endpoints

There are two groups of API endpoints: 
1. Endpoints starting with `/t1` which indicates that the endpoints are public and do not need JWT bearer token
2. Endpoints starting with `/t2` which indicates that the endpoints are private and each request need appropriate JWT bearer token

- Public API `t1`:
  1.  Buyer - Register new **POST** `/t1/buyer/register`
      **Request Body**
      ```
      {
        "email": "xxx",
        "name": "xxx",
        "password": "xxx",
        "alamat_pengiriman": "xxx"
      }
      ```
      Example curl:
      ```
      curl --location --request POST 'localhost:8080/t1/buyer/register' \
      --header 'Content-Type: application/json' \
      --data-raw '{
        "email": "xxx",
        "name": "xxx",
        "password": "xxx",
         "alamat_pengiriman": "xxx"
      }'
      ```
      
  2.  Buyer - Login **POST** `/t1/buyer/login` 
      **Request Body**
      ```
      {
        "email": "xxx",
        "password": "xxx"
      }
      ```
      The response will be a JWT Token of that specific buyer user that can be use for requesting to private API `t2`
      Sample response: 
      ```
      {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6Im1hbWFAYWJzeHp2LmNvbSIsImV4cCI6MTY3MTM0MTc4OSwicm9sZSI6ImJ1eWVyIiwidXNlcl9pZCI6NX0.hebseMMGEeO5iGGJixJpQwwOx1gDe1EwYhvAGZ3KOj8"
      }
      ```
      Example curl:
      ```
       curl --location --request POST 'localhost:8080/t1/buyer/login' \
      --header 'Content-Type: application/json' \
      --data-raw '{
          "email": "xxx",
         "password": "xxx"
       }'
       ```
       
   3. Seller - Register new **POST** `/t1/seller/register`
      **Request Body**
      ```
      {
        "email": "xxx",
        "name": "xxx",
        "password": "xxx",
        "alamat_pickup": "xxx"
      }
      ```
      **Example curl**
      ```
      curl --location --request POST 'localhost:8080/t1/seller/register' \
      --header 'Content-Type: application/json' \
      --data-raw '{
          "email": "xxx",
          "name": "xxx",
          "password": "xxx",
          "alamat_pickup": "xxx"
      }'
      ```
     
   4. Seller - Login **POST** `/t1/seller/login`
      **Request Body**
      ```
      {
        "email": "xxx",
        "password": "xxx"
      }
      ```
      The response will be a JWT Token of that specific seller user that can be use for requesting to private API `t2`
      Sample response:
      ```
      {
          "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6InRoaXNzZWxsZXIiLCJleHAiOjE2NzEzNDM4NjIsInJvbGUiOiJzZWxsZXIiLCJ1c2VyX2lkIjoxfQ.qAcpFI7XvzVKEt04wWl_vZa8cF2AAa5NKcFDUUxD1-0"
      }
      ```
      **Example Curl**
      ```
      curl --location --request POST 'localhost:8080/t1/seller/login' \
      --header 'Content-Type: application/json' \
      --data-raw '{
        "email": "xxx",
        "password": "xxx"
      }'
      ```
      
   5. Get product list **GET** `/t1/product`
      Parameters:
      - `offset`: number of items to be skipped
      - `limit` : number of items to be shown in the response
      
      Sample response:
      ```
      [
        {
            "ID": 1,
            "CreatedAt": "2022-12-17T11:58:04+07:00",
            "UpdatedAt": "2022-12-17T11:58:04+07:00",
            "DeletedAt": null,
            "product_name": "",
            "description": "",
            "price": 5000.5,
            "SellerID": 0
        },
        {
            "ID": 2,
            "CreatedAt": "2022-12-17T11:59:27+07:00",
            "UpdatedAt": "2022-12-17T11:59:27+07:00",
            "DeletedAt": null,
            "product_name": "indomie",
            "description": "",
            "price": 5000.5,
            "SellerID": 1
        }
        ```
        **Example Curl**
        ```
        curl --location --request GET 'localhost:8080/t1/product?offset=0&limit=10'
        ```
 - Private API `t2`:
    For the endpoints below, you need to provide the **JWT token** as authorization in the request header. You can get the token from the login API above (depending on whether you login as seller or buyer account)
    
   1. Buyer - Get Order List **GET** `/t2/buyer/order`
      Get the current ongoing orders for the specific buyer account (buyer information is obtained from JWT token).
      Sample response:
      ```
      [
          {
              "ID": 1,
              "CreatedAt": "2022-12-17T14:49:30+07:00",
              "UpdatedAt": "2022-12-18T10:24:20+07:00",
              "DeletedAt": null,
              "BuyerId": 5,
              "SellerId": 1,
              "DeliverySourceAddress": "jalan jalan",
              "DeliveryDestAddress": "akhir jalan",
              "Items": 5,
              "Quantity": 3,
              "Price": 5000.5,
              "TotalPrice": 15001.5,
              "Status": "Accepted"
          }
      ]
      ```
      **Example Curl**
      ```
        curl --location --request GET 'localhost:8080/t2/buyer/order' \
        --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6Im1hbWFAYWJzeHp2LmNvbSIsImV4cCI6MTY3MTM0NDI5MCwicm9sZSI6ImJ1eWVyIiwidXNlcl9pZCI6NX0.iXTMqup1nt F_102djksIBcQT-LriR7lTD6Id7JazHWY'
      ```
      
   2. Buyer - Create Order **POST** `/t2/buyer/order`
      **Request Body**
      ```
      {
          "source_address": "xxx",
          "destination_address": "xxx",
          "item_id": 0,
          "quantity": 0
      }
      ```
      Sample response:
      ```
        {
            "ID": 2,
            "CreatedAt": "2022-12-18T12:22:00.335228+07:00",
            "UpdatedAt": "2022-12-18T12:22:00.335228+07:00",
            "DeletedAt": null,
            "BuyerId": 5,
            "SellerId": 0,
            "DeliverySourceAddress": "jalan jalan",
            "DeliveryDestAddress": "akhir jalan",
            "Items": 1,
            "Quantity": 3,
            "Price": 5000.5,
            "TotalPrice": 15001.5,
            "Status": "Pending"
       }
       ```
       **Example Curl**
       ```
         curl --location --request POST 'localhost:8080/t2/buyer/order' \
          --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6Im1hbWFAYWJzeHp2LmNvbSIsImV4cCI6MTY3MTM0NDI5MCwicm9sZSI6ImJ1eWVyIiwidXNlcl9pZCI6NX0.iXTMqup1ntF_102djksIBcQT-LriR7lTD6Id7JazHWY' \
          --header 'Content-Type: application/json' \
          --data-raw '{
              "source_address": "xxx",
              "destination_address": "xxx",
              "item_id": 0,
              "quantity": 0
          }'
        ```
      
   3. Seller - Create Product **POST** `/t2/product`
      **Request Body**
      ```
      {
          "product_name": "xxx",
          "product_description": "xxx",
          "price": 0.00
      }
      ```
      Sample response:
      ```
      {
          "ID": 13,
          "CreatedAt": "2022-12-18T12:25:50.6449716+07:00",
          "UpdatedAt": "2022-12-18T12:25:50.6449716+07:00",
          "DeletedAt": null,
          "product_name": "indomie",
          "description": "mie instant",
          "price": 5000.5,
          "SellerID": 1
      }
      ```
      **Example Curl**
      ```
        curl --location --request POST 'localhost:8080/t2/product' \
        --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6InRoaXNzZWxsZXIiLCJleHAiOjE2NzEzNDQ2NzcsInJvbGUiOiJzZWxsZXIiLCJ1c2VyX2lkIjoxfQ.ZGX_Q-_KlOFiC0y10HYXxYUVimoGnY7TJ5o2YcSqbzM' \
        --header 'Content-Type: application/json' \
        --data-raw '{
            "product_name": "xxx",
            "product_description": "xxx",
            "price": 0.00
        }'
      ```
      
   4. Seller - Get Product of the seller **GET** `/t2/seller/product` 
      Parameters:
      - `offset`: number of items to be skipped
      - `limit` : number of items to be shown in the response
      
      Sample response:
      ```
       [
          {
              "ID": 7,
              "CreatedAt": "2022-12-17T12:26:21+07:00",
              "UpdatedAt": "2022-12-17T12:26:21+07:00",
              "DeletedAt": null,
              "product_name": "indomie",
              "description": "this is a test",
              "price": 5000.5,
              "SellerID": 1
          },
          {
              "ID": 8,
              "CreatedAt": "2022-12-17T12:26:22+07:00",
              "UpdatedAt": "2022-12-17T12:26:22+07:00",
              "DeletedAt": null,
              "product_name": "indomie",
              "description": "this is a test",
              "price": 5000.5,
              "SellerID": 1
          }
      ]
      ```
      **Example Curl**
      ```
      curl --location --request GET 'localhost:8080/t2/seller/product?offset=5&limit=5' \
      --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6InRoaXNzZWxsZXIiLCJleHAiOjE2NzEzNDQ2NzcsInJvbGUiOiJzZWxsZXIiLCJ1c2VyX2lkIjoxfQ.ZGX_Q-_KlOFiC0y10HYXxYUVimoGnY7TJ5o2YcSqbzM'
      ```
     
      
   5. Seller - Get order list **GET** `/t2/seller/order` 
      Sample response:
      ```
      [
          {
              "ID": 1,
              "CreatedAt": "2022-12-17T14:49:30+07:00",
              "UpdatedAt": "2022-12-18T10:24:20+07:00",
              "DeletedAt": null,
              "BuyerId": 5,
              "SellerId": 1,
              "DeliverySourceAddress": "jalan jalan",
              "DeliveryDestAddress": "akhir jalan",
              "Items": 5,
              "Quantity": 3,
              "Price": 5000.5,
              "TotalPrice": 15001.5,
              "Status": "Accepted"
          }
      ]
      ```
       **Example Curl**
      ```
      curl --location --request GET 'localhost:8080/t2/seller/order' \
      --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6InRoaXNzZWxsZXIiLCJleHAiOjE2NzEzNDQ2NzcsInJvbGUiOiJzZWxsZXIiLCJ1c2VyX2lkIjoxfQ.ZGX_Q-_KlOFiC0y10HYXxYUVimoGnY7TJ5o2YcSqbzM'
      ```

   6. Seller - Accept pending order **PUT** `/t2/seller/order`
      Parameters:
      - `order_id`: id of the order to be accepted
      
      Sample response:
      ```
      {
          "ID": 3,
          "CreatedAt": "2022-12-18T12:36:36+07:00",
          "UpdatedAt": "2022-12-18T12:36:56.6888182+07:00",
          "DeletedAt": null,
          "BuyerId": 5,
          "SellerId": 1,
          "DeliverySourceAddress": "xxx",
          "DeliveryDestAddress": "xxx",
          "Items": 5,
          "Quantity": 2,
          "Price": 5000.5,
          "TotalPrice": 10001,
          "Status": "Accepted"
      }
      ```
      **Example Curl**
      ```
      curl --location --request PUT 'localhost:8080/t2/seller/order?order_id=3' \
      --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6InRoaXNzZWxsZXIiLCJleHAiOjE2NzEzNDU0MDEsInJvbGUiOiJzZWxsZXIiLCJ1c2VyX2lkIjoxfQ.FHqVgJNxKwtoFUbYITpRcH1cGG_KTXpkGQSw9RWmx2o'
      ```
      
      
   
