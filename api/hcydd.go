package api

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var answers = []string{
	"i'm trying to get pregnant",
	"it interferes with the heroin",
	"i'm 12",
	"because none of you can stop me when I blackout",
	"they don't have enough booze in the bar for me ;)",
	"because my mommy said so",
	"i'm the designated astronaut",
	"i need to keep my blood pure for the Dark Lord",
}

func init() {
	rand.Seed(time.Now().Unix())
}

// Handler is the thing that now.sh will run
func Handler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, answers[rand.Intn(len(answers))])
}
