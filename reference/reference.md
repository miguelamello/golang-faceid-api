OpenAPI Specification v3.1.0

# FaceID API
Version 1.0.0

## 1) Introduction

The FaceID API is designed to establish a backend infrastructure for user authentication through Face Recognition. This API functions by receiving a Face Representation Vector (FRV) from a remote client, then verifying whether this provided FRV corresponds to an existing one stored in the database. Based on this verification, access is either granted or denied to the client.

The FaceID API is accessible to any client capable of sending an HTTP request. Developed using Golang and following the REST architecture, the API offers flexibility in terms of client implementation. Clients can be developed in any programming language and can originate from various platforms such as web, mobile, or desktop. This seamless integration capability enables the incorporation of Face Recognition functionality into a diverse array of applications. Notably, the process of obtaining the FRV from the client falls outside the scope of the FaceID API's responsibilities.

## 2) Servers

```
http://faceid.orionsoft.site
```

## 3) Routes

#### 3.1) POST /vector

This route is used to send a FRV - Face Representation Vector to the API. The FRV is sent as a JSON object in the request body. The API will then verify whether this FRV corresponds to an existing one stored in the database. The API will return a JSON object containing the required data, which can be used by the client to determine whether access should be granted or denied.

Syntax 						| Description 												
---- | ----
Method:						|	POST			
Route:						|	/vector
Parameters:				|	none
Header:						|	Content-Type: application/json
Body:							|	{ "vector": []float64 }
Success Response:	|	{ "result": [{ "id": string, "payload": { "email": string, "name": string, "role": string } }] }
Error Response:		|	{ "statusCode": number, "error": string }
  
#### 3.1.1) HTTP signature

```
POST /vector HTTP/1.1
Host: http://faceid.orionsoft.site
Content-Type: application/json

{
	"vector": [ float64, float64, ... ]
}
```
Note: The vector must be a 128-dimensional array of type float64.
  
#### 3.1.2) Example using Curl

Request:
```
curl --location 'http://faceid.orionsoft.site/vector' \
--header 'Content-Type: application/json' \
--data '{
	"vector": [
		0.14768389, -0.035013203, 0.06823086, -0.11805733, 0.08483968,
		0.018404376, -0.101448506, -0.05162203, -0.018404376, 0.13107507,
		-0.14768389, -0.035013203, 0.11805733, 0.035013203, -0.08483968,
		0.018404376, 0.06823086, 0.05162203, -0.13107507, -0.018404376,
		0.11805733, 0.14768389, -0.035013203, 0.08483968, -0.101448506,
		-0.035013203, 0.08483968, 0.14768389, -0.13107507, -0.05162203,
		0.018404376, 0.035013203, 0.14768389, 0.018404376, -0.035013203,
		-0.08483968, 0.05162203, -0.06823086, -0.13107507, 0.11805733,
		0.035013203, 0.08483968, -0.018404376, -0.14768389, -0.06823086,
		-0.05162203, 0.018404376, 0.11805733, -0.035013203, 0.13107507,
		0.05162203, -0.035013203, 0.14768389, 0.06823086, 0.018404376,
		-0.08483968, 0.05162203, -0.13107507, -0.08483968, 0.11805733,
		-0.035013203, 0.14768389, 0.035013203, -0.018404376, 0.08483968,
		0.05162203, -0.018404376, 0.035013203, -0.06823086, 0.14768389,
		-0.13107507, 0.11805733, 0.035013203, -0.035013203, 0.08483968,
		-0.018404376, 0.11805733, -0.14768389, 0.06823086, -0.05162203,
		0.08483968, 0.035013203, 0.14768389, -0.018404376, 0.13107507,
		-0.11805733, 0.05162203, -0.035013203, 0.14768389, -0.05162203,
		0.018404376, -0.06823086, 0.08483968, -0.035013203, 0.11805733,
		-0.13107507, 0.14768389, 0.035013203, -0.018404376, 0.06823086,
		-0.08483968, -0.05162203, 0.13107507, 0.11805733, 0.035013203,
		-0.14768389, -0.035013203, 0.08483968, -0.018404376, -0.035013203,
		0.14768389, -0.13107507, 0.05162203, 0.06823086, -0.11805733, 
		0.035013203, 0.08483968, 0.018404376, -0.035013203, -0.05162203, 
		0.14768389, 0.08483968, -0.035013203, -0.018404376, -0.13107507, 
		0.11805733, 0.05162203, 0.06823086
	]
}'
```

Response:
```
/// Successful response

{
	"result": [
		{
			"id": "bb54eff4-4440-11ee-bb8e-706979a86d88",
			"payload": {
				"email": "francissinatrao@gmail.com",
				"name": "Francis Albert Sinatra",
				"role": "customer"
			}
		}
	],
	"grant": true
}

Note: When the vector is found in the database, "grant" 
field will be `true` and "result" will contain an 
array with with vector id and the user payload 
associated with the vector.
```

```
// Error response

{
    "result": [],
    "grant": false
}

Note: When the vector is NOT found in the database, 
"grant" field will be `false` and "result" will 
contain an empty array.
```

## 4) Support

For any questions or concerns, please contact the developer at `miguelangeomello@gmail.com`. This API is currently in development and is subject to change, and sould not be used in production environments. If you need a stable version of this API, please contact the developer.

