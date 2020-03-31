# confusion-of-merger
This is a small service which has Restful API endpoints for getting bank and branch details.
###
 Following are the features of this
 use PostgreSQL as a backend database
1. GET API to fetch a bank details, given branch IFSC code
2. GET API to fetch all details of branches, given bank name and a city. 
3. This API should also support limit and offset parameters
4. APIs should be authenticated using a JWT key, with validity = 5 days

To run application 
```
go mod tidy
go run main.go
```
To test the API END POINTS and to obtain JWT bearer token 
```
TOKEN=$(curl GET -H "Accept:application/json" localhost:8090/get-token)
```


 ```
curl -H 'Accept: application/json' -H "Authorization: Bearer ${TOKEN}" 'localhost:8090/banklist?ifsc=UTIB0001856'
 ```
Response                                        302 Found
```
{
    "ifsc": "UTIB0001856",
    "BankDetails": {
        "bank_name": "AXIS BANK",
        "BankID": 0
    },
    "branch": "SRINIVASNAGAR",
    "address": "PLOT NO.578, II PHASE, VI BLOCK,SRINIVASNAGAR, BANASHANKARI III STAGE, BANGALORE 560085",
    "city": "BANGALORE",
    "district": "BANGALORE URBAN",
    "state": "KARNATAKA"
}
```
 ```
curl -H 'Accept: application/json' -H "Authorization: Bearer ${TOKEN}" 'localhost:8090/branchlist?bank_name=HDFC%20BANK&city=BANGALORE'
 ```
Response 302 Found
 ```
 {"total_count":140,"data":[{"ifsc":"HDFC0000261","BankDetails":{"bank_name":"HDFC BANK","BankID":0},"branch":"BANGALORE - JAYANAGAR","address":"PROFESSIONAL COURTNO. 27/7,15TH CROSS, 3RD BLOCKJAYANAGARBANGALOREKARNATAKA560 011","city":"BANGALORE","district":"BANGALORE URBAN","state":"KARNATAKA"},{"ifsc":"HDFC0000286","BankDetails":{"bank_name":"HDFC BANK","BankID":0},"branch":"BANGALORE - ULSOOR","address":"MOUNT KAILASH,NO 33/5 MEANEE AVENUE }.....
 ```