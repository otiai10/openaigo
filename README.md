# openaigo

[![Go](https://github.com/otiai10/openaigo/actions/workflows/go.yml/badge.svg)](https://github.com/otiai10/openaigo/actions/workflows/go.yml)
[![CodeQL](https://github.com/otiai10/openaigo/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/otiai10/openaigo/actions/workflows/codeql-analysis.yml)

Yet another API client for `api.openai.com`.

This library is community-maintained, NOT officially supported by OpenAI.

# Usage Example

```go
package main

import (
  "fmt"
  "github.com/otiai10/openaigo"
)

func main() {
  client := openaigo.NewClient(os.Getenv("OPENAI_APIKEY"))
  request := openaigo.CompletionRequestBody{
    Model:  "text-davinci-003",
    Prompt: []string{"Say this is a test"},
  }
  response, err := client.Completion(nil, request)
  fmt.Println(response, err)
}
```

if you just want to try, hit commands below.

```shell
git clone git@github.com:otiai10/openaigo.git
cd openaigo
OPENAI_APIKEY=YourAPIKey go run ./testapp/main.go
```

# Endpoint Support

- Models
  - [x] [List models](https://beta.openai.com/docs/api-reference/models/list)
  - [x] [Retrieve model](https://beta.openai.com/docs/api-reference/models/retrieve)
- Completions
  - [x] [Create completion](https://beta.openai.com/docs/api-reference/completions/create)
- Edits
  - [x] [Create edits](https://beta.openai.com/docs/api-reference/edits/create)
- Images
  - [x] [Create image](https://beta.openai.com/docs/api-reference/images/create)
  - [x] [Create image edit](https://beta.openai.com/docs/api-reference/images/create-edit)
  - [x] [Create image variation](https://beta.openai.com/docs/api-reference/images/create-variation)
- Embeddings
  - [x] [Create embeddings](https://beta.openai.com/docs/api-reference/embeddings/create)
- Files
  - [ ] [List files](https://beta.openai.com/docs/api-reference/files/list)
  - [ ] [Upload file](https://beta.openai.com/docs/api-reference/files/upload)
  - [ ] [Delete file](https://beta.openai.com/docs/api-reference/files/delete)
  - [ ] [Retrieve file](https://beta.openai.com/docs/api-reference/files/retrieve)
  - [ ] [Retrieve file content](https://beta.openai.com/docs/api-reference/files/retrieve-content)
- Fine-tunes
  - [ ] [Create fine-tune (beta)](https://beta.openai.com/docs/api-reference/fine-tunes/create)
  - [ ] [List fine-tunes (beta)](https://beta.openai.com/docs/api-reference/fine-tunes/list)
  - [ ] [Retrieve fine-tune (beta)](https://beta.openai.com/docs/api-reference/fine-tunes/retrieve)
  - [ ] [Cance fine-tune (beta)](https://beta.openai.com/docs/api-reference/fine-tunes/cancel)
  - [ ] [List fine-tune events (beta)](https://beta.openai.com/docs/api-reference/fine-tunes/events)
  - [ ] [Delete fine-tune model (beta)](https://beta.openai.com/docs/api-reference/fine-tunes/delete-model)
- Moderation
  - [ ] [Create moderation](https://beta.openai.com/docs/api-reference/moderations/create)
- ~~Engines~~ *(deprecated)*
  - ~~[List engines](https://beta.openai.com/docs/api-reference/engines/list)~~
  - ~~[Retrieve engine](https://beta.openai.com/docs/api-reference/engines/retrieve)~~
