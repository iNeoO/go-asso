package commonapi

type ErrorEnvelope struct {
	Status bool      `json:"status"`
	Data   *struct{} `json:"data"`
	Error  string    `json:"error"`
}
