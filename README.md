# go-feishu

go-feishu is a Go client library for accessing the Feishu API. 

> As of now I am writing this library, there is no official Feishu SDK for Go.
> Although Bytedance is the fastest growing tech company in China,
> seems like they don't want to hire a Go developer to write a Go SDK for Feishu. 😂

The goals of this library are:

1. Provide a complete and idiomatic Go client for the Feishu API.
2. Provide a test suite to ensure API compatibility over time.
3. Practice my Go skills.🤖️

## Installation

```bash
go get github.com/joyqi/go-feishu
```

## Usage

### OAuth2 Authentication

Initialize Feishu client with OAuth2 authentication:

```go
import "github.com/joyqi/go-feishu/oauth2"

var conf = &oauth2.Config{
    AppID:        "your-client-id",
    AppSecret:    "your-client-secret",
    RedirectURL:  "your-redirect-url",
}
```
