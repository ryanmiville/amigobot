# amigo-bot
discord bot for mnt amigos

# To Add A New Command...
1. Create a new package
2. Implement the `MessageHandler` interface from `main.go`
3. Add an instance of your `MessageHandler` implementation to the `handlers` array in `main.go`

See the `greet` package as a very simple example.