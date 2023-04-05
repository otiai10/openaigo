package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/otiai10/openaigo"
)

type Scenario struct {
	Name string
	Run  func() (any, error)
}

var (
	OPENAI_API_KEY string

	scenarios = []Scenario{
		{
			Name: "completion",
			Run: func() (any, error) {
				client := openaigo.NewClient(OPENAI_API_KEY)
				request := openaigo.CompletionRequestBody{
					Model:  "text-davinci-003",
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
				request := openaigo.ChatCompletionRequestBody{
					Model: "gpt-3.5-turbo",
					Messages: []openaigo.ChatMessage{
						{Role: "user", Content: "Hello!"},
					},
				}
				return client.Chat(nil, request)
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
		if match != "" {
			if !strings.Contains(scenario.Name, match) && !strings.Contains(fmt.Sprintf("%d", i), match) {
				fmt.Printf("====> Skip\n\n")
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
