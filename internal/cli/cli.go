package cli

import (
	"fmt"
	"github.com/kumak1/gh-futon/internal/gh"
	"github.com/spf13/pflag"
	"os"
	"time"
)

var (
	User string
	From time.Time
	To   time.Time
)

func init() {
	var err error

	help := pflag.BoolP("help", "h", false, "help")

	username := gh.GetLoginUser()
	pflag.StringVarP(&User, "user", "u", username.Login, "specify github username")

	to := pflag.String("to", time.Now().Format(time.DateOnly), "")
	from := pflag.String("from", time.Now().AddDate(0, -6, 0).Format(time.DateOnly), "")

	pflag.Parse()

	To, err = time.Parse(time.DateOnly, *to)
	if err != nil {
		panic(err)
	}

	From, err = time.Parse(time.DateOnly, *from)
	if err != nil {
		panic(err)
	}

	if *help {
		pflag.PrintDefaults()
		os.Exit(1)
	}

	if From.After(To) {
		fmt.Println("invalid date specified.")
		os.Exit(1)
	}
}
