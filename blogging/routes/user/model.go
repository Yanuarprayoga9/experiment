package user

type User struct {
	ID        int     `db:"id" json:"id"`
	Email     string  `db:"email" json:"email"`
	Name      string  `db:"name" json:"name"`
	Nip       *string `db:"nip" json:"nip,omitempty"`
	Role      string  `db:"role" json:"role"`
	CreatedAt string  `db:"created_at" json:"created_at"`
	UpdatedAt *string `db:"updated_at" json:"updated_at,omitempty"`
}

type GetAllUsersResponse struct {
	Users []User `json:"users"`
	Total int    `json:"total"`
}
