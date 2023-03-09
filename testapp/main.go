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

var OPENAI_API_KEY string

var scenarios = []Scenario{
	{
		Name: "completion",
		Run: func() (any, error) {
			client := openaigo.NewClient(OPENAI_API_KEY)
			request := openaigo.CompletionRequestBody{
				Model:  "text-davinci-003",
				Prompt: []string{"Say this is a test"},
			}
			res, err := client.Completion(nil, request)
			if err != nil {
				return res, err
			}
			if len(res.Choices[0].Text) == 0 {
				return res, fmt.Errorf("no response given, that's indeed an error in this case")
			}
			return res, err
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
	{
		Name: "chat_completion_stream",
		Run: func() (any, error) {
			client := openaigo.NewClient(OPENAI_API_KEY)
			request := openaigo.ChatCompletionRequestBody{
				Stream: true,
				Model:  "gpt-3.5-turbo",
				Messages: []openaigo.ChatMessage{
					{Role: "user", Content: "Please write a short novel with 200 words in Haruki Murakami's style."},
					{Role: "user", Content: "Please add \"にゃん\" in every line of the content."},
				},
			}
			res, err := client.Chat(nil, request)
			for a := range res.Stream() {
				fmt.Print(a.Choices[0].Delta.Content)
			}
			fmt.Print("\n")
			return res, err
		},
	},
}

func init() {
	flag.Parse()
	OPENAI_API_KEY = os.Getenv("OPENAI_API_KEY")
}

func main() {
	match := flag.Arg(0)
	total := 0
	var dur time.Duration = 0
	errors := []error{}
	for i, scenario := range scenarios {
		fmt.Printf("\033[1;34m[%03d] %s\033[0m\n", i+1, scenario.Name)
		if match != "" && !strings.Contains(scenario.Name, match) {
			fmt.Printf("====> Skip\n\n")
			continue
		}
		begin := time.Now()
		res, err := scenario.Run()
		elapsed := time.Since(begin)
		if err != nil {
			fmt.Printf("\033[31mError:\033[0m %v\nTime: ", err)
			errors = append(errors, err)
		} else {
			fmt.Printf("[RESULT] %+v\n\033[32mTime:\033[0m ", res)
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
