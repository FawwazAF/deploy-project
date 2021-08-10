# Product in cart

To get all products that available in cart

## How to do

Data constraint :
- :user_id = id from user (need to [register](register.md) first)
- :product_id = id from product 

To get list all of the users products :
- GET /products
- GET /users

Example :

Input product : POST /carts/1/1

## Error 

**Condition** : if 'user_id' and 'product_id' not exist in database


