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

# Contributing
This project uses dep for dependency management. If you need to add a new dependency, [here](https://golang.github.io/dep/docs/installation.html) are instructions for installing it. Reference [the docs](https://golang.github.io/dep/docs/daily-dep.html#adding-a-new-dependency) for how to add dependencies. The current dependencies are packaged with the repo in the [vendor](https://github.com/ryanmiville/amigobot/tree/master/vendor) directory.


Download the repo with `go get github.com/ryanmiville/amigobot`

## To Add A New Command...
1. Create a new package
2. Implement the `Handler` interface found [here](handler.go).
3. Add an instance of your `Handler` implementation to the `handlers` array in [main.go](cmd/amigobot/main.go)

See [greet.go](greet/greet.go) as a very simple example.

## Running Locally
In the `.../amigobot/cmd/amibot` directory, run `go install`

Now you should be able to run the app with `amigobot -t [your-bot-token]

[GoDoc]: https://godoc.org/github.com/ryanmiville/amigobot
[GoDoc Widget]: https://godoc.org/github.com/ryanmiville/amigobot?status.svg
