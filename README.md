# cashier-service

## API Endpoints
#### GET /cashier/inventory
Returns the current inventory of bank notes and coins in the cashier desk, with the value and amount of each type.
#### POST /cashier/purchase
Takes a JSON request body with the following parameters:

- `price:` the price of the product being purchased (float)
- `payment:` the total payment made by the customer (float)

Returns a JSON response with the following parameters:

- `change:` the amount of change to be given to the customer (float)
- `breakdown:` a list of dictionaries representing the bank notes and coins to be given to the customer, with the following keys:
  - `value:` the value of the bank note or coin (float)
  - `amount:` the number of bank notes or coins of this value to be given (int)
#### PUT /admin/inventory/:value/:amount
Take path with the flowwing parameters:
- `value:` the value of the bank note or coin (float)
- `amount:` the number of bank note or coin (int)

Returns a JSON response with the following parameters:
- `value:` the value of the bank note or coin (float)
- `amount:` the number of bank notes or coins of this value to be given (int)
## Example
### Request
```Perl
POST /cashier/purchase
{
  "price": 12.50,
  "payment": 20.00
}
```
### Response
```Perl
{
  "change": 7.50,
  "breakdown": [
    {
      "value": 5.00,
      "amount": 1
    },
    {
      "value": 2.00,
      "amount": 1
    },
    {
      "value": 0.50,
      "amount": 1
    }
  ]
}
```
