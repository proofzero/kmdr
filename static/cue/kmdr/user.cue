package kmdr

#user: {
	apiVersion: string
	kind: "User"
	metadata: #metadata & {
		username: string
	}
	data: {
		name: {
			first: string
			last:  string
		}
		email: string // TODO: regex for email
		publicKey: _
		signatureKey: _
	}
}