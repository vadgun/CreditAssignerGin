# Credit Assigner API

## What it does

The Credit Allocator application implements an algorithm to know which is the best option to distribute the investment into small credits of 300, 500 and 700.
It provides and Gin Framework Endpoint that can be reached when executing the go build file on your computer.


[localhost:8080/credit-assignment/](localhost:8080/credit-assignment/)

### Input Parameters

This application has the following input parameters required when doing the request.

- *Investment* - Amount for investment on JSON format

## Request

POST → /credit-assignment

```javascript
{
    "investment": 400000,
}
```

## Response 200 OK

```javascript
{
    "credit_type_300": 267,
    "credit_type_500": 266,
    "credit_type_700": 267
}
```