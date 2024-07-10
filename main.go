package main

import (
	"fmt"
	link "html-link-parser/parser"
	"strings"
)

var htmlExample = `
<html>
<body>
  <a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>
</body>
</html>
`

func main() {
	r := strings.NewReader(htmlExample)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("% + v\n", links)
}
