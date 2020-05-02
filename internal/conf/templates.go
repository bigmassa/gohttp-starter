package conf

import (
	"html/template"
	"sync"

	"github.com/binaryknights/gutil"
)

var templates *template.Template
var templatesOnce sync.Once

func GetTemplates() *template.Template {
	templatesOnce.Do(func() {
		templates = gutil.ParseTemplates("web/templates")
	})
	return templates
}
