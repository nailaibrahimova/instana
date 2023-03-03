package instana

type SliEventFilterElement struct {
	Type         *string      `json:"type"`
	Name         *string      `json:"name"`
	StringValue  *string      `json:"stringValue"`
	NumberValue  *interface{} `json:"numberValue"`
	BooleanValue *bool        `json:"booleanValue"`
	Key          *string      `json:"key"`
	Value        *string      `json:"value"`
	Operator     *string      `json:"operator"`
	Entity       *string      `json:"entity"`
}
