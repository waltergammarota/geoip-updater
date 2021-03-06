package maxmind_test

import (
	"testing"

	"github.com/crazy-max/geoip-updater/pkg/maxmind"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	cases := []struct {
		name     string
		wantData maxmind.Config
		wantErr  bool
	}{
		{
			name: "Empty license key",
			wantData: maxmind.Config{
				LicenseKey: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid base URL",
			wantData: maxmind.Config{
				LicenseKey: "0123456789",
				BaseURL:    "foo.bar",
			},
			wantErr: true,
		},
		{
			name: "Success",
			wantData: maxmind.Config{
				LicenseKey: "0123456789",
			},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			_, err := maxmind.New(tt.wantData)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
