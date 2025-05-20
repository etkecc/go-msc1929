package msc1929

import (
	"testing"
)

func TestContact_IsEmpty(t *testing.T) {
	tests := []struct {
		name    string
		contact *Contact
		want    bool
	}{
		{"nil contact", nil, true},
		{"empty contact", &Contact{}, true},
		{"contact with email", &Contact{Email: "test@example.com"}, false},
		{"contact with matrix ID", &Contact{MatrixID: "@user:example.com"}, false},
		{"contact with both", &Contact{Email: "test@example.com", MatrixID: "@user:example.com"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.contact.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContact_Roles(t *testing.T) {
	contact := &Contact{Role: RoleAdmin}
	if !contact.IsAdmin() {
		t.Errorf("expected IsAdmin to return true")
	}
	contact.Role = RoleModerator
	if !contact.IsModerator() {
		t.Errorf("expected IsModerator to return true")
	}
	contact.Role = RoleModeratorUnstable
	if !contact.IsModerator() {
		t.Errorf("expected IsModerator to return true for unstable role")
	}
	contact.Role = RoleDPO
	if !contact.IsDPO() {
		t.Errorf("expected IsDPO to return true")
	}
	contact.Role = RoleDPOUnstable
	if !contact.IsDPO() {
		t.Errorf("expected IsDPO to return true for unstable role")
	}
	contact.Role = RoleSecurity
	if !contact.IsSecurity() {
		t.Errorf("expected IsSecurity to return true")
	}
}

func TestResponse_Sanitize(t *testing.T) {
	resp := &Response{
		Contacts: []*Contact{
			{Email: "invalid-email", MatrixID: "invalid"},
			{Email: "valid@example.com", MatrixID: "@az09._=/-+:example.com"},
		},
		SupportPage: "http://valid.url",
	}
	resp.Sanitize()
	if len(resp.Contacts) != 1 {
		t.Errorf("expected 1 valid contact, got %d", len(resp.Contacts))
	}
	if resp.Contacts[0].Email != "valid@example.com" {
		t.Errorf("expected valid email, got %s", resp.Contacts[0].Email)
	}
	if resp.Contacts[0].MatrixID != "@az09._=/-+:example.com" {
		t.Errorf("expected valid matrix ID, got %s", resp.Contacts[0].MatrixID)
	}
}

func TestResponse_IsEmpty(t *testing.T) {
	tests := []struct {
		name     string
		response *Response
		want     bool
	}{
		{"nil response", nil, true},
		{"empty response", &Response{}, true},
		{"response with support page", &Response{SupportPage: "http://valid.url"}, false},
		{"response with valid contact", &Response{Contacts: []*Contact{{Email: "test@example.com"}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.response.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_Clone(t *testing.T) {
	resp := &Response{
		Contacts:    []*Contact{{Email: "test@example.com", MatrixID: "@user:example.com"}},
		SupportPage: "http://valid.url",
	}
	clone := resp.Clone()
	if &resp.Contacts == &clone.Contacts {
		t.Errorf("expected clone to have a different contacts slice reference")
	}
	if clone.SupportPage != resp.SupportPage {
		t.Errorf("expected cloned SupportPage to be the same")
	}
}

func TestResponse_AllEmails(t *testing.T) {
	resp := &Response{
		Contacts: []*Contact{
			{Email: "one@example.com"},
			{Email: "two@example.com"},
		},
	}
	emails := resp.AllEmails()
	if len(emails) != 2 {
		t.Errorf("expected 2 emails, got %d", len(emails))
	}
}

func TestResponse_AllMatrixIDs(t *testing.T) {
	resp := &Response{
		Contacts: []*Contact{
			{MatrixID: "@one:example.com"},
			{MatrixID: "@two:example.com"},
		},
	}
	ids := resp.AllMatrixIDs()
	if len(ids) != 2 {
		t.Errorf("expected 2 matrix IDs, got %d", len(ids))
	}
}

func TestResponse_AdminEmails(t *testing.T) {
	resp := &Response{
		Contacts: []*Contact{{Role: RoleAdmin, Email: "admin@example.com"}}}
	if emails := resp.AdminEmails(); len(emails) != 1 || emails[0] != "admin@example.com" {
		t.Errorf("expected [admin@example.com], got %v", emails)
	}
}

func TestResponse_AdminMatrixIDs(t *testing.T) {
	resp := &Response{
		Contacts: []*Contact{{Role: RoleAdmin, MatrixID: "@admin:example.com"}}}
	if ids := resp.AdminMatrixIDs(); len(ids) != 1 || ids[0] != "@admin:example.com" {
		t.Errorf("expected [@admin:example.com], got %v", ids)
	}
}

func TestResponse_ModeratorEmails(t *testing.T) {
	resp := &Response{
		Contacts: []*Contact{{Role: RoleModerator, Email: "mod@example.com"}}}
	if emails := resp.ModeratorEmails(); len(emails) != 1 || emails[0] != "mod@example.com" {
		t.Errorf("expected [mod@example.com], got %v", emails)
	}
}

func TestResponse_ModeratorMatrixIDs(t *testing.T) {
	resp := &Response{
		Contacts: []*Contact{{Role: RoleModerator, MatrixID: "@mod:example.com"}}}
	if ids := resp.ModeratorMatrixIDs(); len(ids) != 1 || ids[0] != "@mod:example.com" {
		t.Errorf("expected [@mod:example.com], got %v", ids)
	}
}

func TestResponse_DPOEmails(t *testing.T) {
	resp := &Response{
		Contacts: []*Contact{{Role: RoleDPO, Email: "dpo@example.com"}}}
	if emails := resp.DPOEmails(); len(emails) != 1 || emails[0] != "dpo@example.com" {
		t.Errorf("expected [dpo@example.com], got %v", emails)
	}
}

func TestResponse_DPOMatrixIDs(t *testing.T) {
	resp := &Response{
		Contacts: []*Contact{{Role: RoleDPO, MatrixID: "@dpo:example.com"}}}
	if ids := resp.DPOMatrixIDs(); len(ids) != 1 || ids[0] != "@dpo:example.com" {
		t.Errorf("expected [@dpo:example.com], got %v", ids)
	}
}

func TestResponse_SecurityEmails(t *testing.T) {
	resp := &Response{
		Contacts: []*Contact{{Role: RoleSecurity, Email: "security@example.com"}}}
	if emails := resp.SecurityEmails(); len(emails) != 1 || emails[0] != "security@example.com" {
		t.Errorf("expected [security@example.com], got %v", emails)
	}
}

func TestResponse_SecurityMatrixIDs(t *testing.T) {
	resp := &Response{
		Contacts: []*Contact{{Role: RoleSecurity, MatrixID: "@security:example.com"}}}
	if ids := resp.SecurityMatrixIDs(); len(ids) != 1 || ids[0] != "@security:example.com" {
		t.Errorf("expected [@security:example.com], got %v", ids)
	}
}
