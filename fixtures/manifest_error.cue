manifests: {
    bar: {
        apiVersion: "kubelt://v1alpha1"
        kind: "Space"
        something: "error"
        metadata: {
            name: "sm"
            change: "bar" // bork the validation
        }
        data:
            "prop1": "baz"
    }
}
