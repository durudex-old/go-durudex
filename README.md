<div align="center">
  <a href="https://discord.gg/4qcXbeVehZ">
    <img alt="Discord" src="https://img.shields.io/discord/882288646517035028?label=%F0%9F%92%AC%20discord">
  </a>
  <a href="https://github.com/durudex/go-durudex/blob/main/LICENSE">
    <img alt="License" src="https://img.shields.io/github/license/durudex/go-durudex?label=%F0%9F%93%95%20license">
  </a>
  <a href="https://github.com/durudex/go-durudex/stargazers">
    <img alt="GitHub Stars" src="https://img.shields.io/github/stars/durudex/go-durudex?label=%E2%AD%90%20stars&logo=sdf">
  </a>
  <a href="https://github.com/durudex/go-durudex/network">
    <img alt="GitHub Forks" src="https://img.shields.io/github/forks/durudex/go-durudex?label=%F0%9F%93%81%20forks">
  </a>
</div>

<h1 align="center">Durudex Go SDK</h1>

## Setup

```
go get github.com/durudex/go-durudex/sdk
```

## Usage

An example of getting a user using id:

```go
import (
  ...

  "github.com/durudex/durudex-go/sdk"
)

func main() {
  client := sdk.NewClient(sdk.ClientConfig{
    Endpoint:      sdk.TestAPIEndpoint,
    TransportType: sdk.AuthTransportType,
    Transport:     sdk.NewDefaultTransport(),
  })

  ctx := context.Background()
  userId := ksuid.New()

  user, err := client.GetUser(ctx, userId)
  if err != nil { ... }

  ...
}
```

## ⚠️ License

Copyright © 2022 [Durudex](https://github.com/durudex). Released under the MIT license.
