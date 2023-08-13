package repository

import (
	"testing"
)

func TestInit(t *testing.T) {
	// Create a temporary directory for testing
	_ = t.TempDir()

	// Test cases
	testCases := []struct {
		name    string
		path    string
		force   bool
		wantErr bool
	}{
		//add some cases
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := Init(tc.path, tc.force)
			if tc.wantErr && err == nil {
				t.Errorf("Init(%s, %v) - expected error but got none", tc.path, tc.force)
			}
			if !tc.wantErr && err != nil {
				t.Errorf("Init(%s, %v) - unexpected error: %v", tc.path, tc.force, err)
			}
		})
	}
}
