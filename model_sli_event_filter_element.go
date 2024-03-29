package instana

type SliEventFilterElement struct {
	Type         *string      `json:"type"`
	Name         *string      `json:"name"`
	StringValue  *string      `json:"stringValue"`
	NumberValue  *int64       `json:"numberValue"`
	BooleanValue *bool        `json:"booleanValue"`
	Key          *string      `json:"key"`
	Value        *interface{} `json:"value"`
	Operator     *string      `json:"operator"`
	Entity       *string      `json:"entity"`
}
