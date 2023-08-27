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

3.1) POST /vector

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



