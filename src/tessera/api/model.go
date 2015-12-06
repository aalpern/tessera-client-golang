package api

// Placeholder
type Dashboard struct {
    ID int32 `json:"id"`
    Href string `json:"href"`
    ViewHref string `json:"view_href"`
    DefinitionHref string `json:"definition_href"`
    // creation_date string `json:"creation_date"`
    // last_modified_date string `json:"last_modified_date"`
    ImportedFrom string `json:"imported_from"`
    Title string `json:"title"`
    Category string `json:"category"`
    Summary string `json:"summary"`
    Description string `json:"description"`
    // Definition Definition `json:"definition"`
    // Tags []Tag `json:"tags"`
}
