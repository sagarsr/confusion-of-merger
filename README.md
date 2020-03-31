# confusion-of-merger
This is a small service which has Restful API endpoints for getting bank and branch details.
###
 Following are the features of this
 use PostgreSQL as a backend database
1. GET API to fetch a bank details, given branch IFSC code
2. GET API to fetch all details of branches, given bank name and a city. 
3. This API should also support limit and offset parameters
4. APIs should be authenticated using a JWT key, with validity = 5 days

Dependencies to run in local

1. postgres should be running locally
2. Specify ```$DATABASE_URL``` and ```$PORT``` as environment variables 


To run application in local 
```
go mod tidy
go run main.go
```
To build binary and execute 
```
go build .
```

## Testing application deployed on Heroku 

To test the API END POINTS and to obtain JWT bearer token 
```
TOKEN=$(curl GET -H "Accept:application/json" confusionofmerger.herokuapp.com/get-token)
```
  **_Note:_** JWT token is always attached in the auth header of each request. Refer to short.sql file to verify the data obtained is proper.  

 ```
curl -H 'Accept: application/json' -H "Authorization: Bearer ${TOKEN}" 'confusionofmerger.herokuapp.com/banklist?ifsc=ZSBL0000341'
 ```
Response                                        302 Found
```
{
    "ifsc": "ZSBL0000341",
    "BankDetails": {
        "bank_name": "ZILA SAHAKRI BANK LIMITED GHAZIABAD",
        "BankID": 0
    },
    "branch": "LONI",
    "address": "HANSH AUTO MOBILS LONI GHAZIABAD",
    "city": "LONI",
    "district": "GHAZIABAD",
    "state": "UTTAR PRADESH"
}
```
 ```
curl -H 'Accept: application/json' -H "Authorization: Bearer ${TOKEN}" 'confusionofmerger.herokuapp.com/branchlist?bank_name=ABHYUDAYA COOPERATIVE BANK LIMITED&city=MUMBAI&limit=3&offset=6'
 ```
Response 302 Found
 ```
 {
    "total_count": 55,
    "data": [
        {
            "ifsc": "ABHY0065007",
            "BankDetails": {
                "bank_name": "ABHYUDAYA COOPERATIVE BANK LIMITED",
                "BankID": 0
            },
            "branch": "GHATKOPAR",
            "address": "UNIT NO 2 & 3, SILVER HARMONY BLDG,NEW MANIKLAL ESTATE, GHATKOPAR (WEST), MUMBAI-400086",
            "city": "MUMBAI",
            "district": "GREATER MUMBAI",
            "state": "MAHARASHTRA"
        },
        {
            "ifsc": "ABHY0065008",
            "BankDetails": {
                "bank_name": "ABHYUDAYA COOPERATIVE BANK LIMITED",
                "BankID": 0
            },
            "branch": "KANJUR",
            "address": "BHANDARI CO-OP. HSG. SOCIETY, KANJUR VILLAGE, KANJUR (EAST), MUMBAI-400078",
            "city": "MUMBAI",
            "district": "GREATER MUMBAI",
            "state": "MAHARASHTRA"
        },
        {
            "ifsc": "ABHY0065009",
            "BankDetails": {
                "bank_name": "ABHYUDAYA COOPERATIVE BANK LIMITED",
                "BankID": 0
            },
            "branch": "NEHRU NAGAR",
            "address": "ABHYUDAYA BANK BLDG., B.NO.71, NEHRU NAGAR, KURLA (E), MUMBAI-400024",
            "city": "MUMBAI",
            "district": "GREATER MUMBAI",
            "state": "MAHARASHTRA"
        }
    ]
}
 ```