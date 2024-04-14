package main

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestFindAuthorAndZIP(t *testing.T) {
	// Start an HTTP mock
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Mock the HTTP response
	httpmock.RegisterResponder("GET", "https://example.com/sample",
		httpmock.NewStringResponder(200, `<table summary="作家データ"><tr><td></td><td>Haruki Murakami</td></tr></table><table class="download"><a href="sample.zip">Download</a></table>`))

	author, zipURL := findAuthorAndZIP("https://example.com/sample")
	assert.Equal(t, "Haruki Murakami", author, "Expected author name does not match")
	assert.Equal(t, "sample.zip", zipURL, "Expected ZIP URL does not match")
}

func TestFindEntries(t *testing.T) {
	// Start an HTTP mock
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Mock the HTTP response
	httpmock.RegisterResponder("GET", "https://www.aozora.gr.jp/index_pages/person879.html",
		httpmock.NewStringResponder(200, `<ol><li><a href="/cards/000879/card123.html">Sample Title</a></li></ol>`))

	httpmock.RegisterResponder("GET", "http://www.aozora.gr.jp/cards/000879/card123.html",
		httpmock.NewStringResponder(200, `<table summary="作家データ"><tr><td></td><td>Haruki Murakami</td></tr></table><table class="download"><a href="book.zip">Download</a></table>`))

	entries, err := findEntries("https://www.aozora.gr.jp/index_pages/person879.html")
	assert.Nil(t, err, "Expected no error")
	assert.Len(t, entries, 1, "Expected one entry in the result")
	assert.Equal(t, "book.zip", entries[0].ZipURL, "Expected ZIP URL does not match")
}
