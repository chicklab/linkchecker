package main

import (
    "errors"
    "net/http"
    "net/url"
    "golang.org/x/net/html"
)

type Document struct {
    *Selection
    url      *url.URL
    statusCode      int
    rootNode *html.Node
}

func NewRequest(url string) (*Document, error) {

    res, e := http.Get(url)
    if e != nil {
        return nil, e
    }
    return NewRequestFromResponse(res)
}

func NewRequestFromResponse(res *http.Response) (*Document, error) {
    if res == nil {
        return nil, errors.New("Response is nil")
    }
    defer res.Body.Close()
    if res.Request == nil {
        return nil, errors.New("Response.Request is nil")
    }

    root, e := html.Parse(res.Body)
    if e != nil {
        return nil, e
    }

    return NewDocument(root, res.StatusCode, res.Request.URL), nil
}

func NewDocument(root *html.Node, statusCode int, url *url.URL) *Document {

    doc := &Document{nil, url, statusCode, root}
    doc.Selection = newSingleSelection(root, doc)
    return doc
}


type Selection struct {
	Nodes    []*html.Node
	document *Document
}
// Helper constructor to create a selection of only one node
func newSingleSelection(node *html.Node, doc *Document) *Selection {
	return &Selection{[]*html.Node{node}, doc}
}

