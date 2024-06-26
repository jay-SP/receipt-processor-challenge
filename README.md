# receipt-processor-challenge
The Receipt Processor API is a simple webservice designed to calculate reward points based on receipt data. It provides endpoints for submitting receipts and retrieving awarded points.

Built with: Gin web framework (Go)
Storage: In-memory (data does not persist after restart)
API Specification
Endpoint: Process Receipts
Path: /receipts/process
Method: POST
Description: Submits a receipt for processing.
Request Body:
JSON
{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    }
    // ... other items
  ],
  "total": "35.35"
}
Use code with caution.
content_copy
retailer (string, required): Name of the retailer.
purchaseDate (string, required): Date of purchase (YYYY-MM-DD).
purchaseTime (string, required): Time of purchase (HH:MM).
items (array, required): Array of item objects (see Item model below).
total (string, required): Total amount of the purchase (e.g., "35.35").
Response:
Success (200 OK):
JSON
{"id": "f24d1e97-108d-4530-a880-33f1e91243c1"} 
Use code with caution.
content_copy
Error (400 Bad Request):
JSON
{"error": "Invalid receipt data"} 
Use code with caution.
content_copy
Endpoint: Get Points
Path: /receipts/{id}/points
Method: GET
Description: Retrieves points for a processed receipt.
Path Parameter:
id (string): The ID of the receipt obtained from the /receipts/process endpoint.
Response:
Success (200 OK):
JSON
{"points": 100}
Use code with caution.
content_copy
Error (404 Not Found):
JSON
{"error": "No receipt found for that id"}
Use code with caution.
content_copy
