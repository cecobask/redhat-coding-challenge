# redhat-coding-challenge
GoLang REST API service that is able to perform CRUD operations on objects organized by buckets. The service

## Instructions
1. Have go installed
2. Clone the repository
3. Start the server and database as Docker containers:
```
docker-compose up --build
```
4. Sending requests can be done using the terminal (cURL) or using Postman
## Functionality

### List all objects from a bucket
*Request: GET /objects/{bucket}*

To retrieve information about all objects from a bucket you have to specify the bucket (string) request param.

If successful, you will receive the objects as a JSON array and status 200 (OK).
```
curl --location --request GET '0.0.0.0:8080/objects/testbucket'
```

### Retrieve a single object from a bucket
*Request: GET /objects/{bucket}/{objectID}*

To retrieve a single object from a bucket you have to specify a bucket (string) and object ID (uuid/v4) request params.

If successful, you will be able to download the object and receive status 200 (OK).
If object was not found, you will receive an error and status 404 (Not found).
```
curl --location --request GET '0.0.0.0:8080/objects/testbucket/325432b6-6666-43fe-bb73-070090bc810d'
```

### Upload object to a bucket / Replace existing object
*Request: PUT /objects/{bucket}/{objectID}*

To upload a file you have to specify a bucket (string) and object ID (uuid/v4) request params, as well as file in the request body form with key *uploadObject*.

If successful, you will receive the object ID in the response and status 201 (Created).
```
curl --location --request PUT '0.0.0.0:8080/objects/testbucket/325432b6-6666-43fe-bb73-070090bc810d' \
--form 'uploadObject=@"$FILEPATH"'
```

### Delete a single object from a bucket
*Request: DELETE /objects/{bucket}/{objectID}*

To delete a single object from a bucket you have to specify a bucket (string) and object ID (uuid/v4) request params.

If successful, you will be able to download the object and receive status 200 (OK).
If object was not found, you will receive an error and status 404 (Not found).
```
curl --location --request DELETE '0.0.0.0:8080/objects/testbucket/325432b6-6666-43fe-bb73-070090bc810d'
```