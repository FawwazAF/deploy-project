# Register

Register is to get user_id.

## How to do

Data constraint :

```json
{
    "name": "[name in plain text]",
    "email": "[valid email]",
    "password": "[password in plain text]",
    "address": "[address in plain text]",
}
```
* You can use body form-data to input the data

Example :

```json
{
    "name": "fawwaz",
    "email": "faw@123",
    "password": "12345678",
    "address": "bandung",
}
```

## Error 

**Condition** : if email is duplicate


