package util

import (
	"github.com/russross/blackfriday"
	"strings"
)

func Markdown(markdown string) map[string]string {
	parts := strings.SplitN(strings.Replace(markdown, "\r", "", -1), "\n\n", 2)
	params := make(map[string]string)
	params["markdown"] = markdown
	params["template"] = "post"
	for _, line := range strings.Split(parts[0], "\n") {
		pair := strings.SplitN(line, ":", 2)
		if len(pair) > 1 {
			key := strings.TrimSpace(pair[0])
			value := strings.TrimSpace(pair[1])
			if key != "" && value != "" {
				params[key] = value
			}
		}
	}

	if len(parts) > 1 {
		params["content"] = string(blackfriday.MarkdownBasic([]byte(parts[1])))
		params["content"] = strings.Replace(params["content"], "<pre><code>", "<pre>", -1)
		params["content"] = strings.Replace(params["content"], "</code></pre>", "</pre>", -1)
		params["content"] = strings.Replace(params["content"], "\n</pre>", "</pre>", -1)
	}

	return params
}