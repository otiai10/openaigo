# ChatGPT: Chat Completions API Client

This is a usable example of `github.com/otiai10/openaigo`, especially for [`Chat Completions`](https://platform.openai.com/docs/guides/gpt/chat-completions-api) API.

```go
ai := chatgpt.New(token)
conversations, err := ai.Chat(ctx, []chatgpt.Message{
    chatgpt.System("You are an expert of snowboarding"),
    chatgpt.User("Hello! How's it going?"),
})
```