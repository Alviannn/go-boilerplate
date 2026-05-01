package customerror

type ErrorJSON struct {
	Message            string   `json:"message"`
	SourceErrorMessage string   `json:"sourceErrorMessage"`
	StackLine          string   `json:"stackLine,omitempty"`
	Stack              []string `json:"stack,omitempty"`
}
