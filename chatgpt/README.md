# ChatGPT: Chat Completions API Client

This is a usable example of `github.com/otiai10/openaigo`, especially for [`Chat Completions`](https://platform.openai.com/docs/guides/gpt/chat-completions-api) API.

```go
package main

import (
    "github.com/otiai10/openaigo/chatgpt"
    fc "github.com/otiai10/openaigo/functioncall"
)

func main() {
    ai := chatgpt.New(token, "gpt-4-0613")

    ai.Functions = fc.Funcs{
        "get_user_location": {/* */},
        "get_current_date": {/* */},
        "get_weather": {/* */},
    }

    conversation, err := ai.Chat(ctx, []chatgpt.Message{
        chatgpt.User("Should I take my umbrella tomorrow?"),
    })
    // AI calls necessary functions sequentially,
    // and finally reply to the user's question.
}
```
This conversation will look like this:

0. User asked "Should I take my umbrella tomorrow?"
1. Assistant wanted to call "get_user_location"
2. Function replied "Tokyo" to the assistant
3. Assistant wanted to call "get_current_date"
4. Function replied "20230707" to the assistant
5. Assistant wanted to call "get_weather" with ["Tokyo","20230707"]
6. Function replied "sunny"
7. Assistant replied "No you don't need to" to the user

and step 1~6 are done automatically.
