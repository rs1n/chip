package template

import (
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

func ParseWalk(
	t *template.Template, templateRoot, templateExt string,
) (*template.Template, error) {
	err := filepath.Walk(
		templateRoot, func(path string, info os.FileInfo, err error) error {
			if strings.HasSuffix(path, templateExt) {
				_, err = t.ParseFiles(path)
				return err
			}

			return err
		},
	)

	return t, err
}
