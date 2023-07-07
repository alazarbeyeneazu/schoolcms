package constant

type Domain string

const (
	Corporate Domain = "corporate"
	User      Domain = "user"
	System    Domain = "system"
)

type ContextKey any
