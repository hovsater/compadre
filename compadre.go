package compadre

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	lineMatcher       = "\"[^\"]*#[dts]{.+}\""
	expressionMatcher = "#([dts]){([^:}]+)(?::((?:\\}|[^}])+))?}"
)

func Read(filename string, c interface{}) error {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	lines := regexp.MustCompile(lineMatcher)

	content = []byte(lines.ReplaceAllStringFunc(string(content), func(s string) string {
		expressions := regexp.MustCompile(expressionMatcher).FindAllStringSubmatch(s, -1)

		if len(expressions) == 1 && strconv.Quote(expressions[0][0]) == s {
			return build(expressions[0])
		} else {
			return buildMultiple(s, expressions)
		}
	}))

	if err := json.Unmarshal(content, c); err != nil {
		return err
	}

	return nil
}

func build(e []string) string {
	v := stringDefault(os.Getenv(e[2]), e[3])
	return typeDefault(e[1], v)
}

func buildMultiple(value string, exps [][]string) string {
	for _, e := range exps {
		v := stringDefault(os.Getenv(e[2]), e[3])
		value = strings.Replace(value, e[0], v, 1)
	}

	return value
}

func typeDefault(t string, v string) string {
	switch t {
	case "d":
		return stringDefault(v, "0")
	case "t":
		return stringDefault(v, "false")
	default:
		return strconv.Quote(v)
	}
}

func stringDefault(value string, defaultValue string) string {
	if value == "" {
		value = defaultValue
	}
	return value
}
