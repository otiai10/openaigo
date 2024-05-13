# openaigo

[![Go](https://github.com/otiai10/openaigo/actions/workflows/go.yml/badge.svg)](https://github.com/otiai10/openaigo/actions/workflows/go.yml)
[![CodeQL](https://github.com/otiai10/openaigo/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/otiai10/openaigo/actions/workflows/codeql-analysis.yml)
[![App Test over API](https://github.com/otiai10/openaigo/actions/workflows/api.yml/badge.svg)](https://github.com/otiai10/openaigo/actions/workflows/api.yml)
[![License](https://img.shields.io/github/license/otiai10/openaigo)](https://github.com/otiai10/openaigo/blob/main/LICENSE)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fotiai10%2Fopenaigo.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fotiai10%2Fopenaigo?ref=badge_shield)
<br>[![Maintainability](https://api.codeclimate.com/v1/badges/20c434b47940ce8ed511/maintainability)](https://codeclimate.com/github/otiai10/openaigo/maintainability)
[![Go Report Card](https://goreportcard.com/badge/github.com/otiai10/openaigo)](https://goreportcard.com/report/github.com/otiai10/openaigo)
[![codecov](https://codecov.io/github/otiai10/openaigo/branch/main/graph/badge.svg?token=mfAYgn6Uto)](https://codecov.io/github/otiai10/openaigo)
<br>[![Reference](https://img.shields.io/github/v/tag/otiai10/openaigo?sort=semver)](https://pkg.go.dev/github.com/otiai10/openaigo)
[![GoDoc](https://pkg.go.dev/badge/github.com/otiai10/openaigo)](https://pkg.go.dev/github.com/otiai10/openaigo)

Yet another API client for `api.openai.com`.

This library is community-maintained, NOT officially supported by OpenAI.

# Usage Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/otiai10/openaigo"
)

func main() {
	client := openaigo.NewClient(os.Getenv("OPENAI_API_KEY"))
	request := openaigo.ChatRequest{
		Model: "gpt-4o",
		Messages: []openaigo.Message{
			{Role: "user", Content: "Hello!"},
		},
	}
	ctx := context.Background()
	response, err := client.Chat(ctx, request)
	fmt.Println(response, err)
}

```

if you just want to try, hit commands below.

```shell
git clone https://github.com/otiai10/openaigo.git
cd openaigo
OPENAI_API_KEY=YourAPIKey go run ./testapp/main.go
```

See [test app](https://github.com/otiai10/openaigo/blob/main/testapp/main.go) as a working example.

# API Keys?

Visit https://beta.openai.com/account/api-keys and you can create your own API key to get started [for free](https://openai.com/api/pricing/).

# Endpoint Support

- Models
  - [x] [List models](https://beta.openai.com/docs/api-reference/models/list)
  - [x] [Retrieve model](https://beta.openai.com/docs/api-reference/models/retrieve)
- Text Completions
  - [x] [Create completion](https://beta.openai.com/docs/api-reference/completions/create)
- **Chat Completions**
  - [x] [Create Chat Completions](https://platform.openai.com/docs/api-reference/chat/create)
    - [x] [with function_call](https://openai.com/blog/function-calling-and-other-api-updates) <- New
- Edits
  - [x] [Create edits](https://beta.openai.com/docs/api-reference/edits/create)
- Images
  - [x] [Create image (beta)](https://beta.openai.com/docs/api-reference/images/create)
  - [x] [Create image edit (beta)](https://beta.openai.com/docs/api-reference/images/create-edit)
  - [x] [Create image variation (beta)](https://beta.openai.com/docs/api-reference/images/create-variation)
- Embeddings
  - [x] [Create embeddings](https://beta.openai.com/docs/api-reference/embeddings/create)
- Files
  - [x] [List files](https://beta.openai.com/docs/api-reference/files/list)
  - [x] [Upload file](https://beta.openai.com/docs/api-reference/files/upload)
  - [x] [Delete file](https://beta.openai.com/docs/api-reference/files/delete)
  - [x] [Retrieve file](https://beta.openai.com/docs/api-reference/files/retrieve)
  - [x] [Retrieve file content](https://beta.openai.com/docs/api-reference/files/retrieve-content)
- Fine-tunes
  - [x] [Create fine-tune](https://beta.openai.com/docs/api-reference/fine-tunes/create)
  - [x] [List fine-tunes](https://beta.openai.com/docs/api-reference/fine-tunes/list)
  - [x] [Retrieve fine-tune](https://beta.openai.com/docs/api-reference/fine-tunes/retrieve)
  - [x] [Cancel fine-tune](https://beta.openai.com/docs/api-reference/fine-tunes/cancel)
  - [x] [List fine-tune events](https://beta.openai.com/docs/api-reference/fine-tunes/events)
  - [x] [Delete fine-tune model](https://beta.openai.com/docs/api-reference/fine-tunes/delete-model)
- Moderation
  - [x] [Create moderation](https://beta.openai.com/docs/api-reference/moderations/create)
- ~~Engines~~ *(deprecated)*
  - ~~[List engines](https://beta.openai.com/docs/api-reference/engines/list)~~
  - ~~[Retrieve engine](https://beta.openai.com/docs/api-reference/engines/retrieve)~~

# Need `function_call`?

```go
request := openaigo.ChatRequest{
  Messages: []openaigo.Message{
    {Role: "user", Content: "How's the weather today in Tokyo?"},
  },
  Functions: []openaigo.Function{
    {
      Name: "get_weather",
      Parameters: openaigo.Parameters{
        Type:       "object",
        Properties: map[string]map[string]any{
          "location": {"type": "string"},
          "date":     {"type": "string", "description": "ISO 8601 date string"},
        },
        Required: []string{"location"},
      },
    }
  },
}
```

if you want **shorthand**, use [`functioncall`](https://pkg.go.dev/github.com/otiai10/openaigo@v1.4.0/functioncall).

```go
import fc "github.com/otiai10/openaigo/functioncall"

request.Functions = fc.Funcs{
  "get_weather": {GetWeather, "Get weather of the location", fc.Params{
    {"location", "string", "location of the weather", true},
    {"date", "string", "ISO 8601 date string", true},
  }},
}
```

See [test app](https://github.com/otiai10/openaigo/blob/main/testapp/main.go) as a working example.

# Need `stream`?

```go
client := openaigo.NewClient(OPENAI_API_KEY)
request := openaigo.ChatRequest{
  Stream: true,
  StreamCallback: func(res ChatCompletionResponse, done bool, err error) {
    // Do what you want!
    // You might need chan handling here.
    // See test app how you can do it.
    // https://github.com/otiai10/openaigo/search?q=chat_completion_stream
  },
}
```

# Need Proxy?

```go
client := openaigo.NewClient(OPENAI_API_KEY)
// You can set whatever you want
transport := &http.Transport{ Proxy: http.ProxyFromEnvironment }
client.HTTPClient = &http.Client{ Transport: transport }
// Done!
```

# Issues

Report any issue here or any feedback is welcomed.

* https://github.com/otiai10/openaigo/issues
