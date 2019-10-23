package custom1

import "github.com/getupandgo/highly-uncertain-test-task/internal/model"

var StateToExpression = map[string]model.Expression{
	model.State_P: func(data model.Data) float32 {
		return 2*data.D + data.D*float32(data.E)/100
	},
}
