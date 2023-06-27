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


