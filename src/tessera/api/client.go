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
    if err:= json.NewDecoder(res.Body).Decode(&dashboards); err != nil {
        return dashboards, err
    }

    return dashboards, nil
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
