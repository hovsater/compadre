package compadre

import (
	"os"
	"testing"
)

type appconfig struct {
	Int                           int
	IntWithDefault                int
	IntWithoutDefault             int
	String                        string
	StringWithDefault             string
	StringWithoutDefault          string
	StringWithMultipleExpressions string
}

func TestRead(t *testing.T) {
	var c appconfig
	if err := Read("test_fixtures/config.json", &c); err != nil {
		t.Errorf("Failed to read config: %s", err)
	}

	mustMatch(t, "Int", c.Int, 3000)
	mustMatch(t, "IntWithDefault", c.IntWithDefault, 4000)
	mustMatch(t, "IntWithoutDefault", c.IntWithoutDefault, 0)

	mustMatch(t, "String", c.String, "Hello world!")
	mustMatch(t, "StringWithDefault", c.StringWithDefault, "KevinSjoberg")
	mustMatch(t, "StringWithoutDefault", c.StringWithoutDefault, "")
	mustMatch(t, "StringWithMultipleExpressions", c.StringWithMultipleExpressions, "user= port=3000")
}

func TestReadFromEnv(t *testing.T) {
	os.Setenv("DB_SSL", "true")
	os.Setenv("DB_PORT", "5000")
	os.Setenv("DB_USER", "admin")

	var c appconfig
	if err := Read("test_fixtures/config.json", &c); err != nil {
		t.Errorf("Failed to read config: %s", err)
	}

	mustMatch(t, "IntWithDefault", c.IntWithDefault, 5000)
	mustMatch(t, "IntWithoutDefault", c.IntWithoutDefault, 5000)

	mustMatch(t, "StringWithDefault", c.StringWithDefault, "admin")
	mustMatch(t, "StringWithoutDefault", c.StringWithoutDefault, "admin")
	mustMatch(t, "StringWithMultipleExpressions", c.StringWithMultipleExpressions, "user=admin port=5000")
}

func mustMatch(t *testing.T, attribute string, value interface{}, expected interface{}) {
	if value != expected {
		t.Errorf("%s mismatch: %v wanted %v", attribute, value, expected)
	}
}
