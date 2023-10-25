package sandsmad_parser

// DishOfTheDay is a struct containing a dish of the day. A slice of lines describes the dish and the short description provides a summary of the main dish
type DishOfTheDay struct {
	NameOfDay        string   `json:"name_of_day"`
	ShortDescription string   `json:"short_description"`
	Lines            []string `json:"lines"`
}
