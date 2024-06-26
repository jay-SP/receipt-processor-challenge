# Receipt Processor API

>The Receipt Processor API is a simple webservice designed to calculate reward points based on receipt data. It provides endpoints for submitting receipts and retrieving awarded points.

>Built with: Gin web framework (Go)
Storage: In-memory (data does not persist after restart)

>API Specification \
Endpoint: Process Receipts \
Path: /receipts/process \
Method: POST \
Description: Submits a receipt for processing. 

>Endpoint: Get Points \
Path: /receipts/{id}/points \
Method: GET \
Description: Retrieves points for a processed receipt. \
Path Parameter:
id (string): The ID of the receipt obtained from the /receipts/process endpoint.

# To Run

```
docker build -t receipt .  
docker run -p 8080:8080 -it receipt 
``` 
