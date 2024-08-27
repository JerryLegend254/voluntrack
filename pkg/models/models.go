package models

type Volunteer struct {
	FirstName string
	LastName  string
	Email     string
	Available bool
	Password  []byte
	Skills    []Skill
}

type Organization struct {
	Name          string
	Description   string
	Email         string
	Password      []byte
	Opportunities []Opportunity
}

type Opportunity struct {
	ID          int
	Title       string
	Description string
	Status      string
	Skills      []Skill
}

type Skill struct {
	ID    int
	Title string
}
