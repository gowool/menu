package internal

import (
	"fmt"
	"html"
	"strings"
)

func HTMLAttribute(name string, value any) string {
	switch v := value.(type) {
	case bool:
		if !v {
			return name
		}
		value = name
	case string:
		if name == "class" && v == "" {
			return ""
		}
	}
	return fmt.Sprintf(`%s="%s"`, name, html.EscapeString(fmt.Sprintf("%s", value)))
}

func HTMLAttributes(attributes map[string]any) string {
	var b strings.Builder
	for name, value := range attributes {
		if attribute := HTMLAttribute(name, value); attribute != "" {
			b.WriteRune(' ')
			b.WriteString(attribute)
		}
	}
	return b.String()
}

func HTMLClasses(classes []string) string {
	return strings.TrimSpace(strings.Join(classes, " "))
}

func HTMLClassesAny(classes []any) string {
	classStrings := make([]string, len(classes))
	for i, class := range classes {
		classStrings[i] = fmt.Sprintf("%s", class)
	}
	return HTMLClasses(classStrings)
}
