package models

type Book struct {
	Title            string   `json:"title"`
	AuthorName       []string `json:"author_name,omitempty"`
	FirstPublishYear int      `json:"first_publish_year,omitempty"`
	ISBN             []string `json:"isbn,omitempty"`
}

type OpenLibraryResponse struct {
	NumFound int    `json:"numFound"`
	Start    int    `json:"start"`
	Docs     []Book `json:"docs"`
}

type APIResponse struct {
	Query string `json:"query"`
	Books []Book `json:"books"`
	Total int    `json:"total"`
}
