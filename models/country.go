package models

type Country struct {
	Name struct {
		Common   string `json:"common"`
		Official string `json:"official"`
	} `json:"name"`
	Capital []string `json:"capital"`
	// Add other fields as needed
}
