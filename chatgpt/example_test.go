package chatgpt

import (
	"context"
	"fmt"
	"os"

	"github.com/otiai10/openaigo"
	"github.com/otiai10/openaigo/functioncall"
)

var funcs = functioncall.Funcs{
	"get_user_locatin": functioncall.Func{
		Value: func() string {
			return "Tokyo"
		},
		Description: "Get user's location",
	},
	"get_date": functioncall.Func{
		Value: func() string {
			return "2023-07-07"
		},
		Description: "Get date of today",
	},
	"get_weather": functioncall.Func{
		Value: func(location, date string) string {
			return fmt.Sprintf("Weather in %s on %s is sunny", location, date)
		},
		Description: "Get weather of the location on the date",
		Parameters: functioncall.Params{
			{Name: "location", Type: "string", Description: "Location to get weather", Required: true},
			{Name: "date", Type: "string", Description: "Date to get weather", Required: true},
		},
	},
}

func ExampleAI() {

	key := os.Getenv("OPENAI_API_KEY")
	ai := New(key, openaigo.GPT4o)
	ai.Functions = funcs
	conv := []Message{
		User("Should I bring my umbrella tomorrow? You can use functions to get necessary information."),
	}
	res, err := ai.Chat(context.Background(), conv)
	if err != nil {
		panic(err)
	}
	for i, m := range res {
		if i != 0 {
			fmt.Print("->")
		}
		// fmt.Printf("%s (%s): %s\n", m.Role, m.Name, m.Content) // DEBUG
		fmt.Printf("[%d]%s", i, m.Role)
		continue
		// fmt.Printf("[%d] ", i)
		// // Print role name in different color
		// switch m.Role {
		// case "user":
		// 	fmt.Print("\033[36m")
		// case "assistant":
		// 	fmt.Print("\033[32m")
		// case "function":
		// 	fmt.Print("\033[33m")
		// }
		// if m.Role == "assistant" && m.FunctionCall != nil {
		// 	fmt.Printf("%s\033[0m %s\n", m.Role, "(function_call)")
		// 	fmt.Printf("  > `%s(%+v)`\n", m.FunctionCall.Name(), m.FunctionCall.Args())
		// } else {
		// 	fmt.Printf("%s\033[0m\n", m.Role)
		// 	fmt.Printf("  > %s\n", strings.Trim(m.Content, "\n"))
		// }
	}
	// Output: [0]user->[1]assistant->[2]function->[3]assistant->[4]function->[5]assistant->[6]function->[7]assistant
}
