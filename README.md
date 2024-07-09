<div align="center">
    <h1>Fiber Todo Application</h1>
    <p>Todo Application using <a alt="fiber" href="https://github.com/gofiber/fiber">Fiber v3</a></p>
</div>

## Info.

`Fiber Todo Application`은 **fiber v3** 프레임워크 기반의 Go 언어로 작성되었습니다.

## Feature.

#### 1. Clean Architecture 적용

- **Hexagonal Architecture (Port and Adapter)** 구조 적용

    ```
    .
    ├── cmd
    ├── config
    ├── docs
    ├── internal
    │   ├── adapter
    │   │   ├── persistence
    │   │   └── router
    │   ├── core
    │   │   ├── application
    │   │   ├── domain
    │   │   └── port
    │   ├── utils
    │   ├── api.go
    │   └── server.go
    ├── migrations
    ├── sql
    │   ├── queries
    │   └── schema.sql
    ├── docker-compose.yml
    ├── env.yaml
    ├── env-template.yaml
    ├── go.mod
    ├── go.sum
    ├── README.md
    └── sqlc.yaml
    ```

#### 2. OAS(OpenAPI Specification) 3 적용

- [swag v2](https://github.com/swaggo/swag/tree/v2)를 사용하여 OAS 3.1 형식의 json, yaml 생성

    - `swag v2 cli` 설치
        ```bash
        go install github.com/swaggo/swag/v2/cmd/swag@latest
        ```

    - `oas` 파일 생성
        ```bash
        swag init -d cmd,internal --v3.1
        ```

- [gofiber/swagger](https://github.com/gofiber/swagger)를 `fiber v3`, `swag v2`에 맞게 수정하여 Swagger Handler 적용

    - [수정 코드](https://github.com/himitery/fiber-todo/tree/main/config/oas)

#### 3. SQL Generator 사용

- [sqlc](https://github.com/sqlc-dev/sqlc)를 사용하여 SQL 스키마 및 쿼리를 Go 언어로 생성

    ```bash
    sqlc generate
    ```

## Reference

- [https://github.com/bluecheat/microapp-fiber-kit](https://github.com/bluecheat/microapp-fiber-kit)
