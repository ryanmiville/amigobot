# amigo-bot
discord bot for mnt amigos

# To Add A New Command...
1. Create a new package
2. Add your command to the package
3. Create your function to handle the command. The signature _must_ be `func(*discordgo.Session, *discordgo.MessageCreate)`
4. Add your command and handler function to the `commands` map in `main.go`

See the `greet` package as a very simple example.