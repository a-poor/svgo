package svgo

import (
	"fmt"
	"strings"
)

func formatAttributes(attrs map[string]string) string {
	var sattrs string
	for k, v := range attrs {
		sattrs += fmt.Sprintf("%s=\"%s\" ", k, v)
	}
	return strings.TrimSpace(sattrs)
}
