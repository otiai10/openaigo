package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/otiai10/openaigo"
	fc "github.com/otiai10/openaigo/functioncall"
)

type Scenario struct {
	Name string
	Run  func() (any, error)
}

const (
	SKIP = "\033[0;33m====> SKIP\033[0m\n\n"
)

func GetWeather(location string, date float64) (string, error) {
	return "sunny", nil
}

func GetDate() int {
	now := time.Now()
	return now.Year()*10000 + int(now.Month())*100 + now.Day()
}

var (
	OPENAI_API_KEY string

	scenarios = []Scenario{
		{
			Name: "completion",
			Run: func() (any, error) {
				client := openaigo.NewClient(OPENAI_API_KEY)
				request := openaigo.CompletionRequestBody{
					Model:  openaigo.TextDavinci003,
					Prompt: []string{"Say this is a test"},
				}
				return client.Completion(nil, request)
			},
		},
		{
			Name: "image_edit",
			Run: func() (any, error) {
				client := openaigo.NewClient(OPENAI_API_KEY)
				f, err := os.Open("./testdata/baby-sea-otter.png")
				if err != nil {
					return nil, err
				}
				defer f.Close()
				request := openaigo.ImageEditRequestBody{
					Image:  f,
					Prompt: "A cute baby sea otter with big cheese",
					Size:   openaigo.Size256,
				}
				return client.EditImage(nil, request)
			},
		},
		{
			Name: "image_variation",
			Run: func() (any, error) {
				client := openaigo.NewClient(OPENAI_API_KEY)
				f, err := os.Open("./testdata/baby-sea-otter.png")
				if err != nil {
					return nil, err
				}
				defer f.Close()
				request := openaigo.ImageVariationRequestBody{
					Image: f,
					Size:  openaigo.Size256,
				}
				return client.CreateImageVariation(nil, request)

			},
		},
		{
			Name: "chat_completion",
			Run: func() (any, error) {
				client := openaigo.NewClient(OPENAI_API_KEY)
				request := openaigo.ChatRequest{
					Model: openaigo.GPT3_5Turbo,
					Messages: []openaigo.Message{
						{Role: "user", Content: "Hello!"},
					},
				}
				return client.Chat(nil, request)
			},
		},
		{
			// https://platform.openai.com/docs/models/gpt-4
			Name: "[SKIP] chat_completion_GPT4",
			Run: func() (any, error) {
				client := openaigo.NewClient(OPENAI_API_KEY)
				request := openaigo.ChatRequest{
					Model: openaigo.GPT4,
					Messages: []openaigo.Message{
						{Role: "user", Content: "Who are you?"},
					},
				}
				return client.Chat(nil, request)
			},
		},
		{
			Name: "chat_completion_stream",
			Run: func() (any, error) {
				client := openaigo.NewClient(OPENAI_API_KEY)
				data := make(chan openaigo.ChatCompletionResponse)
				done := make(chan error)
				defer close(data)
				defer close(done)
				calback := func(r openaigo.ChatCompletionResponse, d bool, e error) {
					if d {
						done <- e
					} else {
						data <- r
					}
				}
				request := openaigo.ChatCompletionRequestBody{
					Model:          openaigo.GPT3_5Turbo_0613,
					StreamCallback: calback,
					Messages: []openaigo.Message{
						{
							Role:    "user",
							Content: fmt.Sprintf("What are the historical events happend on %s", time.Now().Format("01/02"))},
					},
				}
				res, err := client.ChatCompletion(context.Background(), request)
				if err != nil {
					return res, err
				}
				for {
					select {
					case payload := <-data:
						fmt.Print(payload.Choices[0].Delta.Content)
					case err = <-done:
						fmt.Print("\n")
						return res, err
					}
				}
			},
		},

		// Test case using "function_call"
		{
			Name: "function_call",
			Run: func() (any, error) {
				conversation := []openaigo.Message{
					{Role: "user", Content: "Should I bring an umbrella tomorrow? I'm living around Tokyo."},
				}
				funcs := fc.Funcs{
					"GetDate": fc.Func{GetDate, "A function to get date today", fc.Params{}},
					"GetWeather": fc.Func{GetWeather, "A function to get weather information", fc.Params{
						{"location", "string", "location of the wather", true, nil},
						{"date", "integer", "date MMDD as number", true, nil},
					}},
				}
				client := openaigo.NewClient(OPENAI_API_KEY)
				request := openaigo.ChatRequest{
					Model:     openaigo.GPT3_5Turbo_0613,
					Messages:  conversation,
					Functions: funcs,
				}
				res_1, err := client.Chat(nil, request)
				if err != nil {
					return nil, err
				}
				conversation = append(conversation, res_1.Choices[0].Message)
				if res_1.Choices[0].Message.FunctionCall != nil {
					fmt.Printf("%+v\n", res_1.Choices[0].Message.FunctionCall)
					conversation = append(conversation, openaigo.Message{
						Role:    "function",
						Name:    res_1.Choices[0].Message.FunctionCall.Name(),
						Content: funcs.Call(res_1.Choices[0].Message.FunctionCall),
					})
				}
				request.Messages = conversation
				res_2, err := client.Chat(nil, request)
				if err != nil {
					return nil, err
				}
				conversation = append(conversation, res_2.Choices[0].Message)
				if res_2.Choices[0].Message.FunctionCall != nil {
					fmt.Printf("%+v\n", res_2.Choices[0].Message.FunctionCall)
					conversation = append(conversation, openaigo.Message{
						Role:    "function",
						Name:    res_2.Choices[0].Message.FunctionCall.Name(),
						Content: funcs.Call(res_2.Choices[0].Message.FunctionCall),
					})
				}
				request.Messages = conversation
				res_3, err := client.Chat(nil, request)
				return res_3, err
			},
		},
	}

	list bool
)

func init() {
	flag.BoolVar(&list, "list", false, "List up all names of scenario")
	flag.Parse()
	OPENAI_API_KEY = os.Getenv("OPENAI_API_KEY")
}

func main() {

	if list {
		for i, scenario := range scenarios {
			fmt.Printf("% 2d %s\n", i, scenario.Name)
		}
		return
	}
	match := flag.Arg(0)
	total := 0
	var dur time.Duration = 0
	errors := []error{}
	for i, scenario := range scenarios {
		fmt.Printf("\033[1;34m[%03d] %s\033[0m\n", i+1, scenario.Name)
		if strings.HasPrefix(scenario.Name, "[SKIP]") {
			fmt.Print(SKIP)
			continue
		}
		if match != "" {
			if !strings.Contains(scenario.Name, match) && !strings.Contains(fmt.Sprintf("%d", i), match) {
				fmt.Print(SKIP)
				continue
			}
		}
		begin := time.Now()
		res, err := scenario.Run()
		elapsed := time.Since(begin)
		if err != nil {
			fmt.Printf("\033[31mError:\033[0m %+v\n", err)
			if e, ok := err.(openaigo.APIError); ok {
				fmt.Println("++++++++++++++++++++++")
				// fmt.Println("StatusCode:", e.StatusCode)
				fmt.Println("Status:    ", e.Status)
				fmt.Println("Type:      ", e.Type)
				fmt.Println("Message:   ", e.Message)
				fmt.Println("Code:      ", e.Code)
				fmt.Println("Param:     ", e.Param)
				fmt.Println("++++++++++++++++++++++")
			}
			errors = append(errors, err)
			fmt.Print("Time: ")
		} else {
			fmt.Printf("%+v\n\033[32mTime:\033[0m ", res)
		}
		fmt.Printf("%v\n\n", elapsed)
		dur += elapsed
		total++
	}
	fmt.Println("===============================================")
	fmt.Printf("Total %d scenario executed in %v.\n", total, dur)
	if len(errors) > 0 {
		os.Exit(1)
	}
}
