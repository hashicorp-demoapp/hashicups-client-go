# Hashicups Client (Go)

A Go client package provides a transient connection between [terraform-provider-hashicups](https://github.com/hashicorp/terraform-provider-hashicups) and [product-api](https://github.com/hashicorp-demoapp/product-api-go). You do not need to compile this package, rather the HashiCups provider uses it as a dependency. 


![Hashicups diagram](./docs/terraform-provider-diagram.jpeg)

Using this module, Hashicups provider will be able to establish new client and send the HTTP(s) request over to product api application to perform CRUD operations. It also handles data mapping from user's inputs to `models.go`. The Hashicups URL defined as `http://localhost:19090` and this can be configured [here](https://github.com/hashicorp-demoapp/hashicups-client-go/blob/main/client.go#L11) in case that port has been configured differently. This module also handles API calls to sign up, sign in and sign out for users authentication to use this application.