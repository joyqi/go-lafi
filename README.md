# go-lafi

**lafi = lark + feishu**

go-lafi is a Go client library for accessing the Lark/Feishu API. 

> As of now I am writing this library, there is no official Lark/Feishu SDK for Go.
> Although Bytedance is the fastest growing tech company in China,
> seems like they don't want to hire a Go developer to write a Go SDK for Lark/Feishu. 😂

The goals of this library are:

1. Provide a complete and idiomatic Go client for the Feishu API.
2. Ensure that the library is well tested and documented.
3. Practice my Go skills.🤖️

## Installation

```bash
go get github.com/joyqi/go-lafi
```

## Usage

### OAuth2 Authentication

Initialize Lark client with OAuth2 authentication:

```go
import "github.com/joyqi/go-lafi/oauth2"

var conf = &oauth2.Config{
    AppID:        "your-client-id",
    AppSecret:    "your-client-secret",
    RedirectURL:  "your-redirect-url",
}
```

For Feishu, you can specify the `Type` field to `TypeFeishu`:

```go
import "github.com/joyqi/go-lafi/oauth2"

var conf = &oauth2.Config{
    AppID:        "your-client-id",
    AppSecret:    "your-client-secret",
    RedirectURL:  "your-redirect-url",
    Type:         oauth2.TypeFeishu,
}
```

Get the authorization URL:

```go
url := conf.AuthCodeURL("your-state")
```

Exchange the authorization code for an access token:

```go
token, err := conf.Exchange(ctx, code)
```

Use TokenSource to persist the token and refresh it as needed:

```go
ts := conf.TokenSource(ctx, token)
token, err := ts.Token()
```

### Lark/Feishu API

Features:

- contact
  - [ ] User
  - [ ] Department
  - [x] Group
  - [ ] Unit
  - [ ] EmployeeTypeEnums
  - [ ] CustomAttr
  - [ ] Scope

Initialize API client:

```go
import "github.com/joyqi/go-lafi/oauth2"

client := conf.TenantTokenSource(ctx).Client()
```

Use the client to access the API. For example, to list all groups:

```go
import (
    "github.com/joyqi/go-lafi/oauth2"
    "github.com/joyqi/go-lafi/api/contact"
)

client := conf.TenantTokenSource(ctx).Client()
api := &contact.Group{Client: client}

api.SimpleList(&contact.GroupSimpleListParams{
    PageSize:  100,
})
```