# Go API client for swagger

This specification defines the APIs provided by the application to manage/share the lab resources by reserving and releasing it.

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```
    "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:8081/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ResourceTypeApi* | [**CreateResourceType**](docs/ResourceTypeApi.md#createresourcetype) | **Post** /resourcetype | Create a Resource Type
*ResourceTypeApi* | [**DeleteResourceType**](docs/ResourceTypeApi.md#deleteresourcetype) | **Delete** /resourcetype | Delete Resource type.
*ResourceTypeApi* | [**GetResourceTypes**](docs/ResourceTypeApi.md#getresourcetypes) | **Get** /resourcetype | Get All Resource Types.


## Documentation For Models

 - [ErrorModel](docs/ErrorModel.md)
 - [NewResourceType](docs/NewResourceType.md)
 - [ResourceTypeResponse](docs/ResourceTypeResponse.md)
 - [ResourceTypesResponse](docs/ResourceTypesResponse.md)


## Documentation For Authorization
 Endpoints do not require authorization.


## Author

mithun.bs@gmail.com

