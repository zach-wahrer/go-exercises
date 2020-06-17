package main

import "html/template"

const temp1 = `----------------------------
Title:   {{.Title}}
Year:    {{.Year}}
Rated:   {{.Rated}}
Runtime: {{.Runtime}}
Genre:   {{.Genre}}
----------------------------
`

var output = template.Must(template.New("movie").Parse(temp1))
