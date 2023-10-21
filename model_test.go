package msc1929

import "testing"

var testResponse = &Response{
	Contacts:    []Contact{{Email: "contact@example.com", MatrixID: "@contact:example.com"}},
	Admins:      []Contact{{Email: "admin@example.com"}},
	SupportPage: "https://example.com",
	hasContent:  true,
	emails:      []string{"contact@example.com", "admin@example.com"},
	mxids:       []string{"@contact:example.com"},
}

func TestIsEmpty(t *testing.T) {
	if testResponse.IsEmpty() {
		t.Fail()
	}
}

func TestIsEmpty_True(t *testing.T) {
	if !(&Response{}).IsEmpty() {
		t.Fail()
	}
}

func TestIsEmpty_Nil(t *testing.T) {
	var resp *Response
	if !resp.IsEmpty() {
		t.Fail()
	}
}

func TestResponse(t *testing.T) {
	emails := []string{"contact@example.com", "admin@example.com"}
	mxids := []string{"@contact:example.com"}
	supportPage := "https://example.com"
	resp := *testResponse
	resp.hydrate()

	if resp.IsEmpty() {
		t.Fail()
	}
	if supportPage != resp.SupportPage {
		t.Fail()
	}
	for i, email := range resp.Emails() {
		if emails[i] != email {
			t.Fail()
		}
	}
	for i, mxid := range resp.MatrixIDs() {
		if mxids[i] != mxid {
			t.Fail()
		}
	}
}
