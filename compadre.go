package compadre

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"text/template"
)

var funcMap = template.FuncMap{
	"i": intParse,
	"s": stringParse,
}

func Read(filename string, c interface{}) error {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	t := template.Must(template.New("compadreJSON").Funcs(funcMap).Parse(string(content)))

	var output bytes.Buffer

	if err = t.Execute(&output, nil); err != nil {
		return err
	}

	if err = json.Unmarshal(output.Bytes(), c); err != nil {
		return err
	}

	return nil
}

func parseExp(exp string) string {
	parts := make([]string, 2)

	for i, p := range strings.SplitN(exp, ":", 2) {
		parts[i] = p
	}

	if os.Getenv(parts[0]) == "" {
		return parts[1]
	} else {
		return os.Getenv(parts[0])
	}
}

func intParse(exp string) int {
	value := parseExp(exp)
	intValue, _ := strconv.Atoi(value)
	return intValue
}

func stringParse(exps ...string) string {
	if len(exps) == 1 {
		return strconv.Quote(parseExp(exps[0]))
	} else {
		format := exps[len(exps)-1]
		values := make([]interface{}, len(exps)-1)

		for i, exp := range exps[:len(exps)-1] {
			values[i] = parseExp(exp)
		}

		return strconv.Quote(fmt.Sprintf(format, values...))
	}
}
