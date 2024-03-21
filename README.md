# Getting started

YouremailAPI GO Sdk is a tool to interact with the [YouremailAPI api](https://youremailapi.com)

To start using the SDK, you must first have an account on the platform (you can make one for free [here](https://youremailapi.com/auth/sign-up)).

When you already have your account, you will need to upload an SMTP account with which you want to send emails and you will also have to create your first template.
For more details, we recommend reading this [getting started](https://docs.youremailapi.com/docs/getting-started)

## Install this package

```text
go get -u github.com/youremailapi/go-sdk
```

## Start using it

How to send and email:

```go
import (
	"fmt"

	youremailapigosdk "github.com/youremailapi/go-sdk"
)

func main() {

	var variables = make(map[string]string)
	var bcc = []string{"Federico Juretich <fedejuret@gmail.com>"};

	variables["%user_first_name%"] = "Federico"
	variables["%app_name%"] = "YouremailAPI"

	response, err := youremailapigosdk.SendEmail("<YOUR API KEY>", youremailapigosdk.SendEmailData{
		Template:    "<YOUR TEMPLATE TOKEN>",
		SmtpAccount: "<YOUR SMTP ACCOUNT TOKEN>",
		To:          "contact@youremailapi.com",
		Subject:     "Some Subject",
		Variables:   variables,
		Bcc: bcc
	})

	if err != nil {
		panic(err)
	}

	messageId := response.MessageID
	accepted := response.Accepted

	fmt.Println(messageId)
	fmt.Println(accepted)
}
```

## Info

The project is still in development, but right now, the version you're looking at is usable. It is being used productively in some projects.
Soon more functionalities will be implemented that allow a customizable iteration to be able to create accounts and templates directly from your platform.

If you have any questions or suggestions, you can write to this email: [contact@youremailapi.com](mailto:contact@youremailapi.com)
