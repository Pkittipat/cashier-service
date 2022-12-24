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
## Features
- Accept payment method `cash`.
- Displaying the current inventory of bank notes and coins on the screen for customers to see.
- (idea) Accepting different forms of payment, such as credit cards, debit cards, and mobile payments.
- (idea) Providing options for customers to donate their change to a charity or save it for later use.
## Functions
- A function to calculate the change money when customer pays with the cash value larger than the product price.
- A function to update the inventory of bank notes and coins when a purchase is made or when new bank notes and coins are added to the system.
- (idea) A function to process and validate different forms of payment, such as checking the validity of a credit card or verifying a mobile payment.
- (idea) A function to handle donations or saving of change for later use, if these options are available. This could involve updating a database or other storage system to track the donated or saved change.

## System Flow
- The customer selects the products and proceeds to checkout.
- The system calculates the total price of the products and displays it to the customer.
- The customer selects the payment method and enters the payment information (cash).
- The system processes the payment and updates the inventory accordingly.
- If the payment is successful, the system prints the receipt and gives the change to the customer, if applicable.
- If the payment is unsuccessful, the system displays an error message and asks the customer to try again.
