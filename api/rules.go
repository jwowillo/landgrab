package api

import (
	"github.com/jwowillo/landgrab/convert"
	"github.com/jwowillo/landgrab/game"
	"github.com/jwowillo/pack"
	"github.com/jwowillo/trim"
	"github.com/jwowillo/trim/application"
	"github.com/jwowillo/trim/controller"
	"github.com/jwowillo/trim/response"
)

const (
	// rulesPath is the rulesController's path.
	rulesPath = "/rules"
	// rulesDescriptionPath is the path to the rulesController's
	// description.
	rulesDescriptionPath = descriptionBase + "rules.json"
)

// rulesController is a trim.Controller used to retrieve the standard
// game.Rules.
type rulesController struct {
	controller.Bare
}

// Path returns rulesPath.
func (c rulesController) Path() string {
	return rulesPath
}

// Description of the rulesController located at rulesDescriptionPath.
func (c rulesController) Description() *application.ControllerDescription {
	return must(read(rulesDescriptionPath))
}

// Handle the trim.Request by returning the standard game.Rules.
func (c rulesController) Handle(r trim.Request) trim.Response {
	return response.NewJSON(pack.AnyMap{
		"rules": convert.RulesToJSONRules(game.StandardRules),
	}, trim.CodeOK)
}
