# amigobot
discord bot for mnt amigos

# Commands
## ?yn [prompt]
Send a yes/no question to **@everyone** with prepopulated 👍 👎 reactions

![?yn screenshot](https://user-images.githubusercontent.com/2359050/38430040-1ed46794-398e-11e8-8ee3-8c4a6ba82347.png)
## ?cals [username]
Display a table of the current day's foods and calories from the MyFitnessPal account for the given username (your account _must_ be public for this to work)

![?cals screenshot](https://user-images.githubusercontent.com/2359050/38430094-4bae7020-398e-11e8-8508-80ba99e3800b.png)
## ?macros [username]
Display a table of the current day's macros from the MyFitnessPal account for the given username (your accoutn _must_ be public for this to work) 

![?macros screenshot](https://user-images.githubusercontent.com/2359050/38429988-fffe73be-398d-11e8-9788-715dc68e2dac.png)
## ?greet [someone]
Greet whomever is specified. This command is mostly for example purposes for adding commands to amigo-bot rather than providing any real utility

![?greet screenshot](https://user-images.githubusercontent.com/2359050/38430068-31047134-398e-11e8-93b7-5fd501b60d3a.png)
# To Add A New Command...
1. Create a new package
2. Implement the `Handler` interface [here](handler.go).
3. Add an instance of your `Handler` implementation to the `handlers` array in [main.go](cmd/amigobot/main.go)

See [greet.go](greet/greet.go) as a very simple example.