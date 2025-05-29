package model

type DealOwner struct {
	Email string
	Name  string
	ID    int
}

// DealOwners is the list of all configured deal owners.
var DealOwners = []DealOwner{
	{"ana@mccusa.co", "Ana Nova", 12},
	{"mateus@mccusa.co", "Mateus Lins de Carvalho", 45},
	{"contacto@mccusa.co", "Paola Andrea Carrillo", 2},
	{"elkin@mccusa.co", "Elkin Cárdenas", 7},
	{"ttrujillo@mccusa.co", "Tatiana Trujillo Trujillo", 13},
	{"mariac@mccusa.co", "Maria C Lievano", 25},
	{"coord_agencies@mccusa.co", "Jonathan David Alcantara Cortes", 29},
	{"adriana@mccusa.co", "Adriana Restrepo", 39},
	{"william@mccusa.co", "William Urrego", 43},
	{"vitor@mccusa.co", "Vitor Freire De Almeida", 49},
	{"dfernandez@mccusa.co", "Darlean Fernandez", 76},
	{"meeting@mccusa.co", "MCC USA EB-3 Virtual Consultant", 80},
	{"meetingpresencial@mccusa.co", "Consultor Presencial", 81},
}

// GetByID returns the DealOwner matching the given ActiveCampaign ID, or nil if none.
func GetByID(id int) *DealOwner {
	for _, o := range DealOwners {
		if o.ID == id {
			return &o
		}
	}
	return nil
}

// GetByEmail returns the DealOwner matching the given email, or nil if none.
func GetByEmail(email string) *DealOwner {
	for _, o := range DealOwners {
		if o.Email == email {
			return &o
		}
	}
	return nil
}
