# Receipt Processor Service

## Overview

The Receipt Processor is a web service that processes receipt data and calculates points based on specific rules. The service provides two main endpoints:

1. `/receipts/process`: Accepts receipt data and returns an ID for the processed receipt.
2. `/receipts/{id}/points`: Returns the number of points awarded for a receipt given its ID.

This service is built with Go and is containerized using Docker for easy deployment and testing.

## Requirements

- Docker
- (Optional) Go - if you want to run the service locally without Docker

## Getting Started

### Building the Docker Image

To build the Docker image, navigate to the root directory of the project and run:

```bash
docker build -f build/Dockerfile -t receipt-processor .
```
This command builds the Docker image using the Dockerfile located in the build/ directory and tags it as receipt-processor.

### Running the Service
Once the image is built, you can run the service using:
    
```bash
docker run -p 8080:8080 receipt-processor
```
This command starts the Receipt Processor service and exposes it on port 8080.

## Using the Service
### Process Receipt Endpoint
To process a receipt, send a POST request to /receipts/process with the receipt data in JSON format.

Example request:
```json
{
  "retailer": "M&M Corner Market",
  "purchaseDate": "2022-03-20",
  "purchaseTime": "14:33",
  "items": [
    {
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    }
  ],
  "total": "9.00"
}
```

### Get Points Endpoint
To retrieve the points for a processed receipt, send a GET request to /receipts/{id}/points, replacing {id} with the receipt ID returned from the process receipt endpoint.

Example request:
```bash
GET /receipts/{id}/points
```

## Testing
To ensure the service is functioning as expected, automated tests are included. These tests can be run in the Docker build process.