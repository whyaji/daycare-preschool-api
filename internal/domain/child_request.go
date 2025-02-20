package domain

// Define a request struct to correctly parse JSON
type CreateChildRequest struct {
	Name             string `json:"name"`
	Nickname         string `json:"nickname"`
	BirthPlace       string `json:"birthPlace"`
	BirthDate        string `json:"birthDate"` // Keep as string for parsing
	Gender           string `json:"gender"`
	AlergyInfo       string `json:"alergyInfo"`
	Notes            string `json:"notes"`
	NumberOfSiblings int    `json:"numberOfSiblings"`
	LivingWith       string `json:"livingWith"`
	RegisteredDate   string `json:"registeredDate"` // Keep as string for parsing
	Parents          string `json:"parents"`
	Teachers         string `json:"teachers"`
}
