package web

type SecurityCaseCreateRequest struct {
	Subject string `json:"subject"`
	Risk    string `json:"risk"`
	Analyst string `json:"analyst"`
}

type SecurityCaseUpdateRequest struct {
	ID      string `json:"id"`
	Subject string `json:"subject"`
	Risk    string `json:"risk"`
	Analyst string `json:"analyst"`
}
