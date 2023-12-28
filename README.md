# Trafilea Http Rest Number Server

We have 4 endpoints :

`GET /v1/numbers ` 
Retrieve in unsorted way all the numbers stored 

Response:
```json
[
  {
    "number":21,
    "value":"Type 1"
  },
  {
    "number":1,
    "value":1
  },
  {
    "number":3,
    "value":"Type 1"
  },
  {
    "number":5,
    "value":"Type 2"
  },
  {
    "number":30,
    "value":"Type 3"
  }
]
```

`GET /v1/numbers/{number}`
Retrieve the number if it exist in the memory 
Response:
```json
{
"number": 30,
"value": "Type 3"
}
```
If number doesn't not exist it return 404 not found

POST /v1/numbers 
application/json body
{
"number" : 30
}
DELETE /v1/numbers/{number}
