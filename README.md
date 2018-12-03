# amigobot
[![GoDoc Widget]][GoDoc]

discord bot for mnt amigos
# Commands
## ?yn [prompt]
Send a yes/no question to **@everyone** with prepopulated 👍 👎 reactions

![?yn screenshot](https://user-images.githubusercontent.com/2359050/38431566-a448d60e-3992-11e8-8f07-0c017d839bbc.png)
## ?cals [username]
Display a table of the current day's foods and calories from the MyFitnessPal account for the given username (your account _must_ be public for this to work)

![?cals screenshot](https://user-images.githubusercontent.com/2359050/38431591-b908c16c-3992-11e8-82a7-2272a7133183.png)
## ?macros [username]
Display a table of the current day's macros from the MyFitnessPal account for the given username (your account _must_ be public for this to work) 

![?macros screenshot](https://user-images.githubusercontent.com/2359050/38431608-c639a45a-3992-11e8-8696-b8e2a9d14e29.png)
## ?greet [someone]
Greet whomever is specified. This command is mostly for example purposes for adding commands to amigo-bot rather than providing any real utility

![?greet screenshot](https://user-images.githubusercontent.com/2359050/38431625-d3920ade-3992-11e8-91d0-3bb0b22d3f99.png)

## ?remindme duration [subject]
Remind the user after a specified time delay. The time should be given in the format specified [here](https://golang.org/pkg/time/#ParseDuration), e.g. 10h35m21s. The bot will acknowledge the request with the specified time.

![?remindme acknowledgement](https://user-images.githubusercontent.com/42191246/43987787-99a95fe6-9cf4-11e8-84ae-f3b06cd131d5.PNG)

If no message is supplied, the invoking message will be pinned after the time has elapsed. 

![?remindme with no subject screenshot](https://user-images.githubusercontent.com/42191246/43987786-999c9d92-9cf4-11e8-833c-47fef41bfde5.PNG)

If a message is supplied,a message will be sent with the included reminder.

![?remindme with subject screenshot](https://user-images.githubusercontent.com/42191246/43987788-99b82116-9cf4-11e8-8fe9-407febd0b850.PNG)

## ?decide option [ or option...] 
Decide between the given options, delimited by " or "

![?decide screenshot](https://user-images.githubusercontent.com/42191246/44006290-06b7a35c-9e50-11e8-9007-281e72530a9d.png)

## ?help
Lists the usage of each command

![?help screenshot](https://user-images.githubusercontent.com/2359050/49352669-10f04600-f687-11e8-88ea-e67f9a2fe50f.png)

# Contributing

## Getting Started
This project requires Go 1.11+ modules for dependency management. [Here](https://github.com/golang/go/wiki/Modules) are docs for modules, including usage, adding, and upgrading modules. Simply `git clone` the project **outside** of your $GOPATH, and run `go build ./...` in the root of the project to download all necessary modules. Run `go test ./...` to verify everything is working properly.

## To Add A New Command...
1. Create a new package
2. Implement the `Handler` interface found [here](handler.go).
3. Write a companion test for your new `Handler`
3. Add an instance of your `Handler` implementation to the `handlers` array in [main.go](cmd/amigobot/main.go)

See [greet.go](greet/greet.go) as a very simple example.

## Tests
run all tests with `go test ./...` to verify you haven't broken any command. Again follow the [greet example](greet/greet_test.go) to see how to mock the use of a real discord session.

## Running Locally
In the `.../amigobot/cmd/amigobot` directory, run `go install`

Now you should be able to run the app with `amigobot -t [your-bot-token]`

[GoDoc]: https://godoc.org/github.com/ryanmiville/amigobot
[GoDoc Widget]: https://godoc.org/github.com/ryanmiville/amigobot?status.svg
