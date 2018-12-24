package main

import (
	"github.com/ryanmiville/amigobot"
	"github.com/ryanmiville/amigobot/decide"
	"github.com/ryanmiville/amigobot/greet"
	"github.com/ryanmiville/amigobot/mfp/cals"
	"github.com/ryanmiville/amigobot/mfp/htmlparse"
	"github.com/ryanmiville/amigobot/mfp/macros"
	"github.com/ryanmiville/amigobot/remindme"
	"github.com/ryanmiville/amigobot/spoiler"
	"github.com/ryanmiville/amigobot/yn"
)

//Handlers is the list of Handlers that will be checked for every message
//sent in the channel (except the ones amigobot sends itself)
var Handlers = []amigobot.Handler{
	&cals.Handler{Fetcher: htmlparse.Fetcher{}},
	&macros.Handler{Fetcher: htmlparse.Fetcher{}},
	&yn.Handler{},
	&greet.Handler{},
	&remindme.Handler{},
	&decide.Handler{},
	&spoiler.Handler{},
}
