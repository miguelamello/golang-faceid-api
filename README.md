# golang-faceid-api
This API aims to create a backend infrastructure to allow user authentication by Face Recognition. The API receives a FRV (Face Representation Vector) from remote client and verifies if the given object matchs an existing FRV in the vector database. The API simply grants or deny access to the client.

This API can be used from any client that can send a HTTP request. The API is written in Golang and uses the REST architecture. The client can be written in any language, and can come from any platform (web, mobile, desktop, etc). This facilitates the integration of Face Recognition in any application. The process of getting the FRV from the client is not part of this API. 


