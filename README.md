# Synapsis Backend Challenge

This is an API for Synapsis Backend Challenge, and also an assignment for Backend Engineer Position at Synapsis.

## Version: 1.0

### Terms of service

http://swagger.io/terms/

**Contact information:**  
Alif Dewantara  
http://github.com/alifdwt  
aputradewantara@gmail.com

**License:** Apache 2.0

### Security

**BearerAuth**

| apiKey | _API Key_     |
| ------ | ------------- |
| In     | header        |
| Name   | Authorization |

### /cart

#### DELETE

##### Summary:

Delete cart

##### Description:

Delete cart from logged in user

##### Responses

| Code | Description | Schema                                                          |
| ---- | ----------- | --------------------------------------------------------------- |
| 200  | OK          | [api.cartWithCartItemsResponse](#api.cartWithCartItemsResponse) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| BearerAuth      |        |

#### GET

##### Summary:

Get cart

##### Description:

Get cart from logged in user

##### Responses

| Code | Description | Schema                                                          |
| ---- | ----------- | --------------------------------------------------------------- |
| 200  | OK          | [api.cartWithCartItemsResponse](#api.cartWithCartItemsResponse) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| BearerAuth      |        |

#### POST

##### Summary:

Create cart

##### Description:

Create cart to logged in user

##### Parameters

| Name    | Located in | Description         | Required | Schema                                          |
| ------- | ---------- | ------------------- | -------- | ----------------------------------------------- |
| request | body       | Create cart request | Yes      | [api.createCartRequest](#api.createCartRequest) |

##### Responses

| Code | Description | Schema                                                          |
| ---- | ----------- | --------------------------------------------------------------- |
| 200  | OK          | [api.cartWithCartItemsResponse](#api.cartWithCartItemsResponse) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| BearerAuth      |        |

### /cart-items/{productId}

#### DELETE

##### Summary:

Delete cart item

##### Description:

Delete cart item

##### Parameters

| Name      | Located in | Description | Required | Schema |
| --------- | ---------- | ----------- | -------- | ------ |
| productId | path       | Product ID  | Yes      | string |

##### Responses

| Code | Description | Schema                                                          |
| ---- | ----------- | --------------------------------------------------------------- |
| 200  | OK          | [api.cartWithCartItemsResponse](#api.cartWithCartItemsResponse) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| BearerAuth      |        |

### /categories

#### GET

##### Summary:

List categories

##### Description:

List categories

##### Parameters

| Name      | Located in | Description | Required | Schema  |
| --------- | ---------- | ----------- | -------- | ------- |
| page_id   | query      |             | Yes      | integer |
| page_size | query      |             | Yes      | integer |

##### Responses

| Code | Description | Schema                                            |
| ---- | ----------- | ------------------------------------------------- |
| 200  | OK          | [ [api.categoryResponse](#api.categoryResponse) ] |

#### POST

##### Summary:

Create new category

##### Description:

Add a new category

##### Parameters

| Name     | Located in | Description | Required | Schema                                                  |
| -------- | ---------- | ----------- | -------- | ------------------------------------------------------- |
| category | body       | Category    | Yes      | [api.createCategoryRequest](#api.createCategoryRequest) |

##### Responses

| Code | Description | Schema                                        |
| ---- | ----------- | --------------------------------------------- |
| 200  | OK          | [api.categoryResponse](#api.categoryResponse) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| BearerAuth      |        |

### /categories/{id}

#### GET

##### Summary:

Get category

##### Description:

Get category

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| id   | path       | Category ID | Yes      | string |

##### Responses

| Code | Description | Schema                                        |
| ---- | ----------- | --------------------------------------------- |
| 200  | OK          | [api.categoryResponse](#api.categoryResponse) |

### /orders

#### GET

##### Summary:

List orders

##### Description:

List orders from logged in user

##### Responses

| Code | Description | Schema                                      |
| ---- | ----------- | ------------------------------------------- |
| 200  | OK          | [ [api.orderResponse](#api.orderResponse) ] |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| BearerAuth      |        |

#### POST

##### Summary:

Create order

##### Description:

Create order from cart (Payment method must be [COD, BANK_TRANSFER, E_WALLET])

##### Parameters

| Name  | Located in | Description  | Required | Schema                                            |
| ----- | ---------- | ------------ | -------- | ------------------------------------------------- |
| order | body       | Create order | Yes      | [api.createOrderRequest](#api.createOrderRequest) |

##### Responses

| Code | Description | Schema                                  |
| ---- | ----------- | --------------------------------------- |
| 200  | OK          | [api.orderResponse](#api.orderResponse) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| BearerAuth      |        |

### /products

#### GET

##### Summary:

List products

##### Description:

List products

##### Parameters

| Name      | Located in | Description | Required | Schema  |
| --------- | ---------- | ----------- | -------- | ------- |
| page_id   | query      | Page ID     | No       | integer |
| page_size | query      | Page Size   | No       | integer |

##### Responses

| Code | Description | Schema                                          |
| ---- | ----------- | ----------------------------------------------- |
| 200  | OK          | [ [api.productResponse](#api.productResponse) ] |

#### POST

##### Summary:

Create new product

##### Description:

Add a new product

##### Parameters

| Name    | Located in | Description | Required | Schema                                                |
| ------- | ---------- | ----------- | -------- | ----------------------------------------------------- |
| product | body       | Product     | Yes      | [api.createProductRequest](#api.createProductRequest) |

##### Responses

| Code | Description | Schema                                      |
| ---- | ----------- | ------------------------------------------- |
| 200  | OK          | [api.productResponse](#api.productResponse) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| BearerAuth      |        |

### /products/{id}

#### DELETE

##### Summary:

Delete product

##### Description:

Delete product

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| id   | path       | Product ID  | Yes      | string |

##### Responses

| Code | Description | Schema                                      |
| ---- | ----------- | ------------------------------------------- |
| 200  | OK          | [api.productResponse](#api.productResponse) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| BearerAuth      |        |

#### GET

##### Summary:

Get product

##### Description:

Get product

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| id   | path       | Product ID  | Yes      | string |

##### Responses

| Code | Description | Schema                                      |
| ---- | ----------- | ------------------------------------------- |
| 200  | OK          | [api.productResponse](#api.productResponse) |

#### PUT

##### Summary:

Update product

##### Description:

Update product

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| id   | path       | Product ID  | Yes      | string |

##### Responses

| Code | Description | Schema                                      |
| ---- | ----------- | ------------------------------------------- |
| 200  | OK          | [api.productResponse](#api.productResponse) |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| BearerAuth      |        |

### /users

#### GET

##### Summary:

List users

##### Description:

List users

##### Parameters

| Name      | Located in | Description | Required | Schema  |
| --------- | ---------- | ----------- | -------- | ------- |
| page_id   | query      |             | No       | integer |
| page_size | query      |             | No       | integer |

##### Responses

| Code | Description | Schema                                                            |
| ---- | ----------- | ----------------------------------------------------------------- |
| 200  | OK          | [ [api.userWithProductsResponse](#api.userWithProductsResponse) ] |

#### POST

##### Summary:

Create new user

##### Description:

Add a new user

##### Parameters

| Name | Located in | Description | Required | Schema                                          |
| ---- | ---------- | ----------- | -------- | ----------------------------------------------- |
| user | body       | User        | Yes      | [api.createUserRequest](#api.createUserRequest) |

##### Responses

| Code | Description | Schema                                |
| ---- | ----------- | ------------------------------------- |
| 200  | OK          | [api.userResponse](#api.userResponse) |

### /users/{id}

#### GET

##### Summary:

Get user

##### Description:

Get user

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| user | path       | User        | Yes      | string |

##### Responses

| Code | Description | Schema                                                        |
| ---- | ----------- | ------------------------------------------------------------- |
| 200  | OK          | [api.userWithProductsResponse](#api.userWithProductsResponse) |

### /users/login

#### POST

##### Summary:

Login user

##### Description:

Login user

##### Parameters

| Name | Located in | Description | Required | Schema                                        |
| ---- | ---------- | ----------- | -------- | --------------------------------------------- |
| user | body       | User        | Yes      | [api.loginUserRequest](#api.loginUserRequest) |

##### Responses

| Code | Description | Schema                                          |
| ---- | ----------- | ----------------------------------------------- |
| 200  | OK          | [api.loginUserResponse](#api.loginUserResponse) |

### Models

#### api.cartWithCartItemsResponse

| Name       | Type                            | Description | Required |
| ---------- | ------------------------------- | ----------- | -------- |
| cart_items | [ [db.CartItem](#db.CartItem) ] |             | No       |
| created_at | string                          |             | No       |
| id         | string                          |             | No       |
| user_id    | string                          |             | No       |

#### api.categoryResponse

| Name     | Type                          | Description | Required |
| -------- | ----------------------------- | ----------- | -------- |
| id       | string                        |             | No       |
| name     | string                        |             | No       |
| products | [ [db.Product](#db.Product) ] |             | No       |

#### api.createCartRequest

| Name       | Type   | Description | Required |
| ---------- | ------ | ----------- | -------- |
| product_id | string |             | No       |

#### api.createCategoryRequest

| Name | Type   | Description | Required |
| ---- | ------ | ----------- | -------- |
| name | string |             | No       |

#### api.createOrderRequest

| Name           | Type   | Description | Required |
| -------------- | ------ | ----------- | -------- |
| payment_method | string |             | Yes      |

#### api.createProductRequest

| Name        | Type    | Description | Required |
| ----------- | ------- | ----------- | -------- |
| category_id | string  |             | Yes      |
| description | string  |             | Yes      |
| price       | integer |             | Yes      |
| title       | string  |             | Yes      |

#### api.createUserRequest

| Name      | Type   | Description | Required |
| --------- | ------ | ----------- | -------- |
| email     | string |             | Yes      |
| full_name | string |             | Yes      |
| password  | string |             | Yes      |
| username  | string |             | Yes      |

#### api.loginUserRequest

| Name     | Type   | Description | Required |
| -------- | ------ | ----------- | -------- |
| password | string |             | Yes      |
| username | string |             | Yes      |

#### api.loginUserResponse

| Name         | Type                                                          | Description | Required |
| ------------ | ------------------------------------------------------------- | ----------- | -------- |
| access_token | string                                                        |             | No       |
| user         | [api.userWithProductsResponse](#api.userWithProductsResponse) |             | No       |

#### api.orderResponse

| Name           | Type                              | Description | Required |
| -------------- | --------------------------------- | ----------- | -------- |
| id             | string                            |             | No       |
| order_date     | string                            |             | No       |
| order_items    | [ [db.OrderItem](#db.OrderItem) ] |             | No       |
| payment_method | string                            |             | No       |
| total_cost     | integer                           |             | No       |
| user_id        | string                            |             | No       |

#### api.productResponse

| Name        | Type                        | Description | Required |
| ----------- | --------------------------- | ----------- | -------- |
| category    | [db.Category](#db.Category) |             | No       |
| created_at  | string                      |             | No       |
| description | string                      |             | No       |
| id          | string                      |             | No       |
| price       | integer                     |             | No       |
| title       | string                      |             | No       |
| updated_at  | string                      |             | No       |
| user_id     | string                      |             | No       |

#### api.userResponse

| Name       | Type   | Description | Required |
| ---------- | ------ | ----------- | -------- |
| created_at | string |             | No       |
| email      | string |             | No       |
| full_name  | string |             | No       |
| username   | string |             | No       |

#### api.userWithProductsResponse

| Name       | Type                          | Description | Required |
| ---------- | ----------------------------- | ----------- | -------- |
| created_at | string                        |             | No       |
| email      | string                        |             | No       |
| full_name  | string                        |             | No       |
| products   | [ [db.Product](#db.Product) ] |             | No       |
| username   | string                        |             | No       |

#### db.CartItem

| Name       | Type    | Description | Required |
| ---------- | ------- | ----------- | -------- |
| cart_id    | string  |             | No       |
| id         | string  |             | No       |
| product_id | string  |             | No       |
| quantity   | integer |             | No       |

#### db.Category

| Name | Type   | Description | Required |
| ---- | ------ | ----------- | -------- |
| id   | string |             | No       |
| name | string |             | No       |

#### db.OrderItem

| Name              | Type    | Description | Required |
| ----------------- | ------- | ----------- | -------- |
| id                | string  |             | No       |
| order_id          | string  |             | No       |
| price_at_purchase | integer |             | No       |
| product_id        | string  |             | No       |
| quantity          | integer |             | No       |

#### db.Product

| Name        | Type    | Description | Required |
| ----------- | ------- | ----------- | -------- |
| category_id | string  |             | No       |
| created_at  | string  |             | No       |
| description | string  |             | No       |
| id          | string  |             | No       |
| price       | integer |             | No       |
| title       | string  |             | No       |
| updated_at  | string  |             | No       |
| user_id     | string  |             | No       |
