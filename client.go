package api

import (
    "io"
    "fmt"
    "encoding/json"
    "strings"
    "net/url"
    "net/http"
)

type Client struct {
    RootURI  url.URL
    *http.Client
}

const kApplicationJson = "application/json; charset=utf-8"
const kUserAgent       = "monster.partyhat.co/tessera/api/client"

// New creates a new Tessera client instance. The root URL provided
// will be normalized, stripping a trailing forward slash if one
// exists.
func New(rootURI string) (*Client, error) {
    uri, err := url.Parse(strings.TrimRight(rootURI, "/"))
    if err != nil {
        return nil, err
    }
    return &Client {
        *uri,
        &http.Client {},
    }, nil
}

//
// newRequest is a private method to centralize creating new HTTP
// requests for all Tessera API calls.
//
func (client *Client) newRequest(method, path string, body io.Reader) (*http.Request, error) {
    url := client.RootURI
    url.Path = fmt.Sprintf("%v%v", url.Path, path)

    request, err := http.NewRequest(method, url.String(), body)
    if err != nil {
        return nil, err
    }

    request.Header.Add("Accept", kApplicationJson)
    request.Header.Add("Content-Type", kApplicationJson)
    request.Header.Add("User-Agent", kUserAgent)

    return request, nil
}

// -----------------------------------------------------------------------------
// Dashboard methods
// -----------------------------------------------------------------------------

func (client *Client) getDashboards(req *http.Request, definition bool) ([]Dashboard, error) {
    var dashboards = make([]Dashboard, 0)

    if definition {
        req.URL.RawQuery = "definition=true"
    }

    // Send it
    res, err := client.Do(req)
    if err != nil {
        return dashboards, err
    }
    defer res.Body.Close()

    // Read and parse the response body
    err = json.NewDecoder(res.Body).Decode(&dashboards)
    return dashboards, nil
}

//
// Fetch a single dashboard by ID.
//
func (client *Client) GetDashboard(id int32, definition bool) (Dashboard, error) {
    var dashboard = Dashboard{}

    req, err := client.newRequest("GET", fmt.Sprintf("/api/dashboard/%d", id), nil)
    if err != nil {
        return dashboard, err
    }

    if definition {
        req.URL.RawQuery = "definition=true"
    }

    res, err := client.Do(req)
    if err != nil {
        return dashboard, err
    }
    defer res.Body.Close()

    err = json.NewDecoder(res.Body).Decode(&dashboard)
    return dashboard, err
}

//
// List all dashboards in the system. If the definition parameter is
// true, return the dashboards with their complete definition. If it's
// false, only the dashboard metadata will be returned.
//
// In the event of an error, an empty array and an error will be returned.
//
func (client *Client) ListDashboards(definition bool) ([]Dashboard, error) {
    req, err := client.newRequest("GET", "/api/dashboard/", nil)
    if err != nil {
        return make([]Dashboard, 0), err
    }

    return client.getDashboards(req, definition)
}

//
// List all dashboards tagged with a specific tag. If the definition
// parameter is true, return the dashboards with their complete
// definition. If it's false, only the dashboard metadata will be
// returned.
//
// In the event of an error, an empty array and an error will be returned.
//
func (client *Client) ListDashboardsByTag(tag string, definition bool) ([]Dashboard, error) {
    req, err := client.newRequest("GET", fmt.Sprintf("/api/dashboard/tagged/%s", tag), nil)
    if err != nil {
        return make([]Dashboard, 0), err
    }

    return client.getDashboards(req, definition)
}

//
// List all dashboards belonging to a specific category. If the
// definition parameter is true, return the dashboards with their
// complete definition. If it's false, only the dashboard metadata
// will be returned.
//
// In the event of an error, an empty array and an error will be returned.
//
func (client *Client) ListDashboardsByCategory(category string, definition bool) ([]Dashboard, error) {
    // Construct the request
    req, err := client.newRequest("GET", fmt.Sprintf("/api/dashboard/category/%s", category), nil)
    if err != nil {
        return make([]Dashboard, 0), err
    }

    return client.getDashboards(req, definition)
}

// -----------------------------------------------------------------------------
// Auxiliary API Methods
// -----------------------------------------------------------------------------

//
// List all tags that exist on dashboards, and the number of
// dashboards tagged with each one.
//
func (client *Client) ListTags() ([]Tag, error) {
    var tags = make([]Tag, 0)

    req, err := client.newRequest("GET", "/api/tag/", nil)
    if err != nil {
        return tags, err
    }

    res, err := client.Do(req)
    if err != nil {
        return tags, err
    }
    defer res.Body.Close()

    err = json.NewDecoder(res.Body).Decode(&tags)
    return tags, err
}

//
// List all categories that dashboards are organized into, and the
// number of dashboards in each one.
//
func (client *Client) ListCategories() ([]Category, error) {
    var categories = make([]Category, 0)

    req, err := client.newRequest("GET", "/api/dashboard/category/", nil)
    if err != nil {
        return categories, err
    }

    res, err := client.Do(req)
    if err != nil {
        return categories, err
    }
    defer res.Body.Close()

    err = json.NewDecoder(res.Body).Decode(&categories)
    return categories, err
}
