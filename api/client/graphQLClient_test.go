package client

import "testing"

func TestGetGraphQLApiURL(t *testing.T) {
	cases := []struct {
		server, want string
	}{
		{"github.com", "https://api.github.com/graphql"},
		{"githubinc.com", "https://githubinc.com/api/graphql"},
	}
	for _, c := range cases {
		got := GetGraphQLApiURL(c.server)
		if got != c.want {
			t.Errorf("GetGraphQLApiURL(%q) == %q, want %q", c.server, got, c.want)
		}
	}
}
