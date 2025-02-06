package components

import (
	"fmt"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)


type AlpineStore map[string]string

func AlpineStoreInit(as AlpineStore) Node {
	script := "document.addEventListener('alpine:init', () => {"

	for k, v := range as {
		script += fmt.Sprintf("Alpine.store('%s', %s)", k, v)
	}

	script += "})"

	return Script(Raw(script))
}
