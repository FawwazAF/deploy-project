# API Project alta_store

API Project alta_store is API project of fake e-commerce which 
you can test to get response as a customer

## Open Endpoints

Endpoints which not required any authentication :

* [Browsing product](md/browsing.md)           : `GET /products/:category/:type`
* [Input product into cart](md/cartPost.md)    : `POST /carts/:user_id/:product_id`
* Product in cart                              : `GET /carts`
* [Register](md/register.md)                   : `POST /users`
* [Login](md/login.md)                         : `POST/jwt/user`


## Authenticated Endpoints

Endpoints which need authentication to get 200 response :


* [Checkout](md/checkout.md)        : `POST /jwt/checkout/:user_id`
* [Transaction](md/transaction.md)  : `POST /jwt/transaction/:user_id`

