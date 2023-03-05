package engine

// httpNodeDecorate http node decorate
func httpNodeDecorate() map[string]func(n *httpNode, v string) {
	return map[string]func(n *httpNode, v string){
		"method": func(n *httpNode, v string) {
			n.method = v
		},
		"contenttype": func(n *httpNode, v string) {
			n.contentType = v
		},
		"endpoint": func(n *httpNode, v string) {
			n.endpoint = v
		},
		"root": func(n *httpNode, v string) {
			n.root = v
		},
		"headers": func(n *httpNode, v string) {
			n.headers = v
		},
	}
}
