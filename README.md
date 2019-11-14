# Moloon - distributed serverless Edge

## Building

```bash
docker-compose build
```

## Configuration

Moolon can be configured using a config file or enviroment variables, for instance

```bash
moloon controller --config ./etc/.moloon.yaml
```


### Environment Variables

Name | Type | Default | Description
---|---|---|---
PORT | string | localhost:3000 | http address (accepts also port number only for heroku compability)  
LOG_LEVEL | string | debug | log level
LOG_TEXTLOGGING | bool | false | defaults to json logging
DATABASE_URL | string | postgres://postgres:postgres<br>@localhost:5432/gobase?sslmode=disable | PostgreSQL connection string
AUTH_LOGIN_URL | string | http://localhost:3000/login | client login url as sent in login token email
AUTH_LOGIN_TOKEN_LENGTH | int | 8 | length of login token
AUTH_LOGIN_TOKEN_EXPIRY | time.Duration | 11m | login token expiry
AUTH_JWT_SECRET | string | random | jwt sign and verify key - value "random" creates random 32 char secret at startup (and automatically invalidates existing tokens on app restarts, so during dev you might want to set a fixed value here)
AUTH_JWT_EXPIRY | time.Duration | 15m | jwt access token expiry
AUTH_JWT_REFRESH_EXPIRY | time.Duration | 1h | jwt refresh token expiry
DISCOVERY_CONFIG | string | kubernetes, or file 

## Kubernetes dependencies

Make sure you have the dependency set to the right version, according to [this](https://github.com/kubernetes/client-go/blob/master/INSTALL.md#go-modules). As an example for using Kubernetes 1.15 client use

```bash
go get k8s.io/client-go@kubernetes-1.15.0
```

## Contributing

Any feedback and pull requests are welcome and highly appreciated. Please open an issue first if you intend to send in a larger pull request or want to add additional features.
