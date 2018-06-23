# slack-message

♯♯♯ minimum ♯♯♯

Post a message using slack incoming webhook.

```go
package main

import (
	"fmt"

	"github.com/yyoshiki41/slack-message"
)

func main() {
	p := slack.Payload{
		Text:        "Hello, worlds.",
		Channel:     "#general",
	}
	err := slack.Post("https://hooks.slack.com/******", p)
	if err != nil {
		fmt.Println(err)
	}
}
```
