# go-grpc-gateway-openapi

This is a sample project to demonstrate how to use [buf](https://buf.build/) to generate go, grpc, gateway and openapi files. The project also includes a sample swagger ui to test the generated files.
This swagger ui is working with the grpc server and the grpc gateway on the same port. It also has a grpc client to test the grpc server.

## Prerequisites
Install go, protoc, buf, grpc-gateway, swaggerui.

It has googleapis proto files already, so you don't need to download them.

SwaggerUI can be downloaded from their [GitHub Repo](https://github.com/swagger-api/swagger-ui). Once downloaded, place the content of dist folder somewhere in your Go project. For example, swaggerui.
After that, move grpc.swagger.json file to swaggerui/swagger.json folder and change swagger-initializerr.js file to load swagger.json file from swaggerui/swagger.json folder in the url.

```javascript
    url: "./swagger.json"
```

## To Run
1. `buf generate` to generate the go, grpc, gateway, swagger, and openapi files.
2. Copy the generated `api/gen/v1/grpc.swagger.json` to `swaggerui/swagger.json`
3. Run `go server/main.go` to start the grpc server.
4. Run `go run main.go` to start the http client.
5. Navigate to `http://localhost:8080/swaggerui/` to see the swagger ui.
6. Execute the `GET` and `POST` requests to see the results.