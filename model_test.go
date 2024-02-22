package msc1929

import "testing"

var (
	testResponseOK = &Response{
		Contacts: []*Contact{
			{Email: "admin@example.com", MatrixID: "@admin:example.com", Role: RoleAdmin},
			{Email: "security@example.com", MatrixID: "@security:example.com", Role: RoleSecurity},
			{Email: "just@example.com", MatrixID: "@just:example.com", Role: "some-other-role"},
			{Email: "missing-role@example.com", MatrixID: "@missing-role:example.com"},
			{MatrixID: "@missing-email:example.com", Role: "missing-email"},
			{Email: "missing-mxid@example.com", Role: "missing-mxid"},
			{Role: "missing-email-mxid"},
			{},
		},
		SupportPage: "https://example.com",
	}
	testResponseInvalid = &Response{
		Contacts: []*Contact{
			{Email: "@invalid:email", MatrixID: "invalid-mxid"},
		},
		SupportPage: "invalid url;://",
	}
)

func TestIsEmpty(t *testing.T) {
	if testResponseOK.Clone().IsEmpty() {
		t.Error("expected non-empty response from testResponseOK, got empty response")
	}
	if !testResponseInvalid.Clone().IsEmpty() {
		t.Errorf("expected empty response from testResponseInvalid, got %+v", testResponseInvalid)
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

func TestResponseValid(t *testing.T) {
	adminEmails := []string{"admin@example.com"}
	adminMxids := []string{"@admin:example.com"}
	securityEmails := []string{"security@example.com"}
	securityMxids := []string{"@security:example.com"}
	allEmails := []string{"admin@example.com", "security@example.com", "just@example.com", "missing-role@example.com", "missing-mxid@example.com"}
	allMXIDs := []string{"@admin:example.com", "@security:example.com", "@just:example.com", "@missing-role:example.com", "@missing-email:example.com"}
	supportPage := "https://example.com"

	r := testResponseOK.Clone()

	if r.SupportPage != supportPage {
		t.Errorf("expected support page %s, got %s", supportPage, r.SupportPage)
	}
	slicesEqual(t, r.AdminEmails(), adminEmails)
	slicesEqual(t, r.AdminMatrixIDs(), adminMxids)
	slicesEqual(t, r.SecurityEmails(), securityEmails)
	slicesEqual(t, r.SecurityMatrixIDs(), securityMxids)
	slicesEqual(t, r.AllEmails(), allEmails)
	slicesEqual(t, r.AllMatrixIDs(), allMXIDs)
}

func TestResponseInvalid(t *testing.T) {
	r := testResponseInvalid.Clone()
	r.Sanitize()

	if len(r.Contacts) != 0 {
		t.Errorf("expected no contacts, got %d", len(r.Contacts))
	}
	if r.SupportPage != "" {
		t.Errorf("expected empty support page, got %s", r.SupportPage)
	}
}

func slicesEqual(t *testing.T, a, b []string) {
	t.Helper()

	if len(a) != len(b) {
		t.Errorf("slices %+v and %+v have different lengths: %d and %d", a, b, len(a), len(b))
	}
	for i, v := range a {
		if b[i] != v {
			t.Errorf("slices %+v and %+v differ at index %d: %s and %s", a, b, i, v, b[i])
		}
	}
}
