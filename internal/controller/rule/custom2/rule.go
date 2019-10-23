package custom2

import "github.com/getupandgo/highly-uncertain-test-task/internal/model"

var CommandToState = map[string]string{
	"101": model.State_M,
	"110": model.State_T,
}

var StateToExpression = map[string]model.Expression{
	model.State_M: func(data model.Data) float32 {
		return float32(data.F) + data.D + (data.D * float32(data.E) / 100)
	},
}
