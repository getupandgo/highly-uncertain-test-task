package base

import "github.com/getupandgo/highly-uncertain-test-task/internal/model"

var CommandToState = map[string]string{
	"011": model.State_T,
	"110": model.State_M,
	"111": model.State_P,
}

func DecorateCommand(extend map[string]string) map[string]string {
	for baseCmd, baseState := range CommandToState {
		_, ok := extend[baseCmd]
		if ok {
			continue
		}
		extend[baseCmd] = baseState
	}

	return extend
}

var StateToExpression = map[string]model.Expression{
	model.State_M: func(data model.Data) float32 {
		return data.D + (data.D * float32(data.E) / 10)
	},
	model.State_P: func(data model.Data) float32 {
		return data.D + (data.D * float32(data.E-data.F) / 25.5)
	},
	model.State_T: func(data model.Data) float32 {
		return data.D - (data.D * float32(data.F) / 30)
	},
}

func DecorateExpression(extend map[string]model.Expression) map[string]model.Expression {
	for baseCmd, baseExpression := range StateToExpression {
		_, ok := extend[baseCmd]
		if ok {
			continue
		}
		extend[baseCmd] = baseExpression
	}

	return extend
}
