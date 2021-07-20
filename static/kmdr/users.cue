package kmdr

#users: {
	current: string
	available: [alias=string]: {
		name: alias
	}
}

// users: #users & {
// 	current: "foo"
// 	available: {
// 		foo: {
// 			name: "foo"
// 		}
// 	}
// }