package controller

import (
	"errors"
	"github.com/getupandgo/highly-uncertain-test-task/internal/controller/rule"
	"github.com/getupandgo/highly-uncertain-test-task/internal/model"
)

type Controller struct{}

func (c Controller) CalculateSubstitution(rule string, input model.Input) (string, float32, error) {
	if input.D == 0 || input.E == 0 || input.F == 0 {
		return "", 0, errors.New("numbers cannot be zero")
	}

	command := boolToString(input.A) + boolToString(input.B) + boolToString(input.C)

	return c.substitutionForCommand(rule, command, model.Data{
		D: input.D,
		E: input.E,
		F: input.F,
	})
}

func (c Controller) substitutionForCommand(ruleset string, command string, data model.Data) (string, float32, error) {
	r, ok := rule.RegisteredRules[ruleset]
	if !ok {
		return "", 0, errors.New("no such rule registered")
	}

	state, ok := r.CommandToState[command]
	if !ok {
		return "", 0, errors.New("no such command registered")
	}

	expr, ok := r.StateToExpression[state]
	if !ok {
		return "", 0, errors.New("no expressions for state")
	}

	return state, expr(data), nil
}

func boolToString(b bool) string {
	if b {
		return "1"
	} else {
		return "0"
	}
}
