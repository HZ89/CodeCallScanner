package git

import (
	"testing"
)

func TestRepo_Clone(t *testing.T) {
	type fields struct {
		config GitArgs
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := New(tt.fields.config)
			if err != nil {
				t.Fatalf("New() error = %v", err)
			}
			if err := r.Clone(); (err != nil) != tt.wantErr {
				t.Errorf("Repo.Clone() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := r.CheckOut(); (err != nil) != tt.wantErr {
				t.Errorf("Repo.Checkout() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
