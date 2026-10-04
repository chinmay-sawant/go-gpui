// Package invite is the invite page: the invite link, email invites, and the team list.
package invite

// Data is the data the invite page prints.
type Data struct {
	URL     string
	Pending []InvitePending
	Members []InviteMember
}

// InvitePending is one invite that has not been accepted yet.
type InvitePending struct {
	Email string
	When  string
}

// InviteMember is one teammate in the team list.
type InviteMember struct {
	Name  string
	Email string
	Mark  string
}

// Default returns the invite page data.
func Default() Data {
	return Data{
		URL: "https://wisprflow.ai/invite/chinmay",
		Pending: []InvitePending{
			{"ravi@example.com", "2 days ago"},
			{"meera@example.com", "Yesterday"},
		},
		Members: []InviteMember{
			{"Priya Nair", "priya@wisprflow.ai", "PN"},
			{"Arjun Mehta", "arjun@wisprflow.ai", "AM"},
			{"Maya Iyer", "maya@wisprflow.ai", "MI"},
			{"Rahul Desai", "rahul@wisprflow.ai", "RD"},
		},
	}
}
