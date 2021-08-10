# Checkout 

Checkout is a data that show total quantity and price of product that in cart

## How to do

Data constraint :

* :user_id = id of an user ([login](login.md) first)
* jwt token = jwt token after login

Output :

* total_price = total price is needed to do [transaction](transaction.md)

```json
{
    "total_price": 4000000
}
```

## Error 

**Condition** : if jwt token is not valid / expired or user_id not exist