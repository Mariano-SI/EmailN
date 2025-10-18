package campaing

import "time"

type Contact struct{
	Email string
}

type Campaing struct {
	ID        string
	Name      string
	CreatedOn time.Time
	Content string
	Contacts []Contact
}