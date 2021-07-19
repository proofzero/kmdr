package kmdr

// TODO: use the cue.mod to import defintions

#metadata: {
        name: =~"^[0-9a-zA-Z]{2,}$" | "0"
        [_]:  string
}
metadata: #metadata
#node: {
        apiVersion: string
        kind:       string
        metadata:   metadata
        data?: [_]: _
}