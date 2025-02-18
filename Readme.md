[//]: # (Copyright 2025 Canonical Ltd.)
[//]: # (SPDX-License-Identifier: AGPL-3.0)

# Identity Platform API
This repo hosts all [protocol buffers](https://protobuf.dev/) definitions for the [Admin UI](https://github.com/canonical/identity-platform-admin-ui) APIs.
Also, Golang code generated will be here.

This repo relies on [Buf](https://buf.build) for code generation and protocol buffers for APi definitions.
It also relies on the [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway?tab=readme-ov-file) Buf plugin
to add a REST API reverse proxy for the generated gRPC server.

Buf uses remote plugins available on its platform.
In order to install the Buf CLI just run
```shell
# it's going to prompt for your password since it needs `sudo` to install
# the buf command in /usr/local/bin
make install-buf
```


## Generate Go and Swagger files
To generate Golang files and the swagger file run
```shell
make generate
```

This will create all the Go files needed, and also an `api.swagger.json` file under the `openapi` folder.
To get the OpenAPI v3 version of the API, run
```shell
make openapi-v3
```

To clean up all generated files, run
```shell
make clean
```

### Tools
- [Buf](https://buf.build)
- [Protocol buffers](https://protobuf.dev/)
- [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway?tab=readme-ov-file)
