# Input product into cart

To input product into cart

## How to do

Data constraint :
* :user_id = id from user (need to [register](register.md) first)
* :product_id = id from product 

* qty = Quantity product :
```json
{
    "qty": "[more than 1]"
}
```

To get list all of the users products :
* GET /products
* GET /users

Example :

Input product : POST /carts/1/1
```json
{
    "qty": 5
}
```

## Error 

**Condition** : if 'user_id' and 'product_id' not exist in database


