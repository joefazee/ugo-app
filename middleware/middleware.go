package middleware

import (
	"github.com/joefazee/ladiwork/data"
	"github.com/joefazee/ugo"
)

type Middleware struct {
	App    *ugo.Ugo
	Models data.Models
}
