package api

type RawFields map[string]interface{}

type Tag struct {
    ID int32 `json:"id"`
    Href string `json:"href"`
    Name string `json:"name"`
    Description string `json:"description"`
    Color string `json:"color"`
    Count int32 `json:"count"`
}

type Category struct {
    Name string `json:"name"`
    Count int32 `json:"count"`
}

type Dashboard struct {
    ID int32 `json:"id"`
    Href string `json:"href"`
    ViewHref string `json:"view_href"`
    DefinitionHref string `json:"definition_href"`
    // string for now
    creation_date string `json:"creation_date"`
    // string for now
    last_modified_date string `json:"last_modified_date"`
    ImportedFrom string `json:"imported_from"`
    Title string `json:"title"`
    Category string `json:"category"`
    Summary string `json:"summary"`
    Description string `json:"description"`
    Definition Definition `json:"definition"`
    Tags []Tag `json:"tags"`
}

type DashboardItem struct {
    ItemID string `json:"item_id"`
    ItemType string `json:"item_type"`
    CssClass string `json:"css_class"`
    Height uint8 `json:"height"`
    Style string `json:"style"`
    Title string `json:"title"`
}

type Container struct {
    /* How to represent polymorphic dashboard items here is tricky */
    Items []DashboardItem `json:"items"`
    *DashboardItem
}

type Query struct {
    Name string `json:"name"`
    Targets []string `json:"targets"`
}

type Definition struct {
    Queries map[string]Query `json:"queries"`
    *Container
}
