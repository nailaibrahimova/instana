package instana

type SliEventFilterExpression struct {
	Type            string                  `json:"type"`
	LogicalOperator string                  `json:"logicalOperator"`
	Elements        []SliEventFilterElement `json:"elements"`
}
