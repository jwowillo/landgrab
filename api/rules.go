package api

import (
	"github.com/jwowillo/landgrab/convert"
	"github.com/jwowillo/landgrab/game"
	"github.com/jwowillo/trim"
	"github.com/jwowillo/trim/application"
	"github.com/jwowillo/trim/response"
)

// rulesPath is the rulesController's path.
const rulesPath = "/rules"

// rulesController is a trim.Controller used to retrieve the standard
// game.Rules.
type rulesController struct{}

func (c rulesController) Trimmings() []trim.Trimming {
	return []trim.Trimming{newValidateToken()}
}

// Path returns rulesPath.
func (c rulesController) Path() string {
	return rulesPath
}

// Description of the rulesController located at rulesDescriptionPath.
func (c rulesController) Description() *application.ControllerDescription {
	return &application.ControllerDescription{
		Get: &application.MethodDescription{
			Response:       "Rules which are used",
			Authentication: "must provide Token",
			Limiting:       "limit of the Token",
		},
	}
}

// Handle the trim.Request by returning the standard game.Rules.
func (c rulesController) Handle(r trim.Request) trim.Response {
	return response.NewJSON(map[string]convert.JSONRules{
		"rules": convert.RulesToJSONRules(game.StandardRules),
	}, trim.CodeOK)
}
