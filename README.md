# amigobot
discord bot for mnt amigos

# To Add A New Command...
1. Create a new package
2. Implement the `Handler` interface [here](handler.go).
3. Add an instance of your `Handler` implementation to the `handlers` array in [main.go](cmd/amigobot/main.go)

See [greet.go](greet/greet.go) as a very simple example.