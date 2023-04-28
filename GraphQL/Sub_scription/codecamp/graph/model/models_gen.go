// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewUser struct {
	Email string `json:"email"`
}

type Notification struct {
	ID    string `json:"id"`
	Seen  bool   `json:"seen"`
	Text  string `json:"text"`
	Title string `json:"title"`
}

type Subscription struct {
	NotificationAdded *User `json:"notificationAdded"`
}

type UpdateNotification struct {
	ID     string `json:"id"`
	UserID string `json:"userID"`
	Seen   bool   `json:"seen"`
}

type UpdateUser struct {
	ID    string  `json:"id"`
	First *string `json:"first,omitempty"`
	Last  *string `json:"last,omitempty"`
	Email *string `json:"email,omitempty"`
}

type User struct {
	ID            string          `json:"id"`
	First         string          `json:"first"`
	Last          string          `json:"last"`
	Email         string          `json:"email"`
	Notifications []*Notification `json:"notifications"`
}
