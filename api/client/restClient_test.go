package client

import (
	"testing"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{}
var server string

func TestGetRestApiURL(t *testing.T) {
	rootCmd.PersistentFlags().StringVarP(&server, "server", "s", "github.com", "Hostname of the GitHub Enterprise server. Using github.com if omitted")

	cases := []struct {
		server, restPath, want string
	}{
		{"github.com", "organizations", "https://api.github.com/organizations"},
		{"", "/organizations", "https://api.github.com/organizations"},
		{"githubinc.com", "/admin/organizations", "https://githubinc.com/api/v3/admin/organizations"},
		{"githubinc.com", "organizations", "https://githubinc.com/api/v3/organizations"},
	}

	for _, c := range cases {
		if c.server != "" {
			viper.Set("server", c.server)
		}
		got := GetRestApiURL(c.restPath)
		if got != c.want {
			t.Errorf("GetRestApiURL(%q, %q) == %q, want %q", c.server, c.restPath, got, c.want)
		}
	}
}
