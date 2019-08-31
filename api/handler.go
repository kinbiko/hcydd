package api

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const page = `
<html style="
	border-color: #e8e6e3;
	background: #181a1b;
">
<body>
<a href="/" style="color: white;">
	<div>
		<h2 style="
			left: 0;
			font-family: courier;
			line-height: 200px;
			margin-top: -100px;
			position: absolute;
			text-align: center;
			top: 50%%;
			width: 100%%;
		">%s</h1>
	</div>
</a>
</body>
</html>
`

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
	fmt.Fprintf(res, page, answers[rand.Intn(len(answers))])
}
