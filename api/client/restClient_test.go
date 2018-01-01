package client

import "testing"

func TestGetRestApiURL(t *testing.T) {
	cases := []struct {
		server, restPath, want string
	}{
		{"github.com", "organizations", "https://api.github.com/organizations"},
		{"githubinc.com", "organizations", "https://githubinc.com/api/v3/organizations"},
	}
	for _, c := range cases {
		got := GetRestApiURL(c.server, c.restPath)
		if got != c.want {
			t.Errorf("GetRestApiURL(%q, %q) == %q, want %q", c.server, c.restPath, got, c.want)
		}
	}
}
