package state

type Domain struct {
	Name string
	ID   string
}
type AuthDomains struct {
	Corporate Domain
	System    Domain
	User      Domain
}
