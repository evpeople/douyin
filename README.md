# douyin

## Introduction

a demo backend of demo-douyin,the demo is divided into ? main sections.

| Service Name    |  Usage    | Framework    | protocol    | Path                   | IDL                                      |
| --------------- | ------------ | ---------- | -------- | ---------------------- | ----------------------------------------- |
| demoapi         | http interface | kitex/gin  | http     | cmd/api  |                                           |
| demouser | user data management | kitex/gorm | protobuf | cmd/core/user |        idl/user.proto  |

### call relations
```
后续补充
```

### catalog introduce

| catalog       | introduce      |
| ---------- | ---------------- |
| pkg/constants   | constant        |
| pkg/errno      | customized error number |
| pkg/middleware | RPC middleware     |
| pkg/tracer  | init jaeger     |
| dal   | db operation              |
| pack       | data pack         |
| service    | business logic   |

## Quick Start

### 1.Setup Basic Dependence
```shell
docker-compose up
```
### 2.Run User RPC Server
```shell
cd cmd/core/user
sh build.sh
sh output/bootstrap.sh
```
### 3.Run API Server
```shell
cd cmd/api
chmod +x run.sh
./run.sh
```
### 4.Jaeger 

visit `http://127.0.0.1:16686/` on  browser.

## API request

[the doc of api ](https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18901232)