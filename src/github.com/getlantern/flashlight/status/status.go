package status

import (
	"bytes"
	"errors"
	"html/template"
	"strings"
)

type cannotFindServerT struct {
	ServerName   string
	ErrorMessage string
}

func normalizeError(err error) string {
	if err != nil {
		content := strings.SplitN(strings.TrimSpace(err.Error()), "\n", 2)
		return strings.TrimSpace(content[0])
	}
	return ""
}

// CannotFindServer creates and returns a generic "cannot find server" error.
func CannotFindServer(server string, errMessage error) ([]byte, error) {
	var err error
	var buf []byte
	var tmpl *template.Template

	if errMessage == nil {
		errMessage = errors.New("Unknown error.")
	}

	buf, err = Asset("generic_error.html")

	if err != nil {
		return nil, err
	}

	tmpl, err = template.New("status_error").Parse(string(buf))
	if err != nil {
		return nil, err
	}

	data := cannotFindServerT{
		ServerName:   server,
		ErrorMessage: normalizeError(errMessage),
	}

	out := bytes.NewBuffer(nil)

	if err = tmpl.Execute(out, data); err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}
