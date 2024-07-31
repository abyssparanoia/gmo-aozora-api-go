# GMO Aozora Net Bank Open Api GO SDK

## About

GMO あおぞらネット銀行について

https://gmo-aozora.com/

GMO あおぞらネット銀行 API 開発者ポータルについて

https://api.gmo-aozora.com/ganb/developer/

## Version

1.0.0

## Requirements

Golang 1.8+

## Installation

- get the repository from _Github_\
  `$ go get github.com/abyssparanoia/gmo-aozora-api-go`

- get a package with the following command \
  `$ go get github.com/antihax/optional`

## Getting started

### Enviroment

Add the configuration below into your config file\
 `vi conf.json`

- stg

  conf.json

  ```json
  {
    "AUTH_BASE_URL": "https://stg-api.gmo-aozora.com/ganb/api/auth/v1",
    "JWT_ISSUER": "https://stg-api.gmo-aozora.com/",
    "AUTH_PATH": "/authorization",
    "TOKEN_PATH": "/token",
    "SALT": "PleaseDefineYourself"
  }
  ```

  [configuration.go - Personal ](./personalclient/configuration.go)

  ```go
  	BasePath:      "https://stg-api.gmo-aozora.com/ganb/api/personal/v1",
  ```

  [configuration.go - Corporate ](./corporateclient/configuration.go)

  ```go
  	BasePath:      "https://stg-api.gmo-aozora.com/ganb/api/corporation/v1",
  ```

  [configuration.go - Webhook ](./webhookclient/configuration.go)

  ```go
  	BasePath:      "https://stg-api.gmo-aozora.com/ganb/api/webhooks/v1",
  ```

- prod

  conf.json

  ```json
  {
    "AUTH_BASE_URL": "https://api.gmo-aozora.com/ganb/api/auth/v1",
    "JWT_ISSUER": "https://api.gmo-aozora.com/",
    "AUTH_PATH": "/authorization",
    "TOKEN_PATH": "/token",
    "SALT": "PleaseDefineYourself"
  }
  ```

  [configuration.go - Personal ](./personalclient/configuration.go)

  ```go
  	BasePath:      "https://api.gmo-aozora.com/ganb/api/personal/v1",
  ```

  [configuration.go - Corporate ](./corporateclient/configuration.go)

  ```go
  	BasePath:      "https://api.gmo-aozora.com/ganb/api/corporation/v1",
  ```

  [configuration.go - Webhook ](./webhookclient/configuration.go)

  ```go
  	BasePath:      "https://api.gmo-aozora.com/ganb/api/webhooks/v1",
  ```

## Documentation

- [Authorization](docs/)
- [Corporate](corporateclient/docs/)
- [Personal](personalclient/docs/)
- [Webhook](webhookclient/docs/)

## Autor

GMO Aozora Net Bank, Ltd. (open-api@gmo-aozora.com)

## Licence

[MIT](https://github.com/abyssparanoia/gmo-aozora-api-go/blob/master/LICENSE)
