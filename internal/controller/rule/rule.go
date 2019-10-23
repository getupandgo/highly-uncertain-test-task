package rule

import (
	"github.com/getupandgo/highly-uncertain-test-task/internal/controller/rule/base"
	"github.com/getupandgo/highly-uncertain-test-task/internal/controller/rule/custom1"
	"github.com/getupandgo/highly-uncertain-test-task/internal/controller/rule/custom2"
	"github.com/getupandgo/highly-uncertain-test-task/internal/model"
)

type Rule struct {
	CommandToState    map[string]string
	StateToExpression map[string]model.Expression
}

var RegisteredRules = map[string]Rule{
	"base": {
		CommandToState:    base.CommandToState,
		StateToExpression: base.StateToExpression,
	},
	"custom1": {
		CommandToState:    base.CommandToState,
		StateToExpression: base.DecorateExpression(custom1.StateToExpression),
	},
	"custom2": {
		CommandToState:    base.DecorateCommand(custom2.CommandToState),
		StateToExpression: base.DecorateExpression(custom2.StateToExpression),
	},
}
