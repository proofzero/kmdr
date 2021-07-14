// spaces: foo: {
//     apiVersion: "kubelt.com/v1alpha1"
//     metadata: {
//         name:         "foo"
//     }
//     data: {
//         index: true
//     }
// }
// bar: {
//     apiVersion: "kubelt.com/v1alpha1"
//     metadata: {
//         change: "bar" // bork the validation
//     }
// }


// package manager for CNDs?
// definitions: {
//   rob: "..."    
// }

// materlizations: {
//  path:schema => localfs:plugin(csv):generate
// }

// scripts: {
//     name: {
//         //do things with program flows?
//     }
// }

manifests: {
    bar: {
        apiVersion: "kubelt://v1alpha1"
        // kind: "Space"
        something: "error"
        metadata: {
            name: "sm"
            change: "bar" // bork the validation
        }
        data:
            "prop1": "baz"
    }
}
