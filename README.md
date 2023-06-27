# Go REST API
REST API designed in Golang

## API

#### /schools
* `GET` : Get all schools in paginated format
* `POST` : Create a new school

#### /schools/:id
* `GET` : Get a school information with its id 
* `PUT` : Update a project with its id information
* `DELETE` : Delete a school with its id information

# REST API

The REST API endpoints and their usage example is described below.

## GET requst to check the availability of the server

### Request

`GET /v1/healthcheck`

    curl.exe -i http://localhost:4000/v1/healthcheck

### Response

    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Tue, 27 Jun 2023 10:52:49 GMT
    Content-Length: 69

    {"environment":"development","status":"available","version":"1.0.0"}

### Request

`GET /v1/schools/5`

    curl.exe localhost:4000/v1/schools/5

### Response

    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Tue, 27 Jun 2023 11:53:01 GMT
    Content-Length: 212

    {"id":5,"created_at":"2023-06-27T17:23:01.890541+05:30","name":"ApplTree","level":"High School","contact":"Anna smith","phone":"601-4411","address":"14 Apple tree","mode":["Blended/Hybrid","Online"],"version":1}


