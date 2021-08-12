package kmdr

#metadata: {
	name: =~"^[0-9a-zA-Z]{2,}$" | "0"
	[_]:  string
}

#node: {
	apiVersion: string
	kind:       string
	metadata:   #metadata
	data?: [_]: _
}

#manifests & {
	[_=string]: #node
}
