# Transaction 

Transaction is a data that tell if the transaction is finished.It will soft delete all data in checkout and cart that available.

## How to do

Data constraint :

* paid = Total price that you transferred

```json
{
    "paid": "[sufficient money]"
}
```

Example : 

* total_price that obtained while checkout input into paid

```json
{
    "paid": 4000000
}
```

## Error 

**Condition** : if paid is not sufficient