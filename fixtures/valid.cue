manifests: {
    foo: {
        apiVersion: "kubelt://v1alpha1"
        kind: "Space"
        metadata: {
            name:         "foo"
        }
        data: {
            index: false
        }
    }
    bar: {
        apiVersion: "kubelt://v1alpha1"
        kind: "Space"
        metadata: {
            description: "hi"
        }
    }
}
