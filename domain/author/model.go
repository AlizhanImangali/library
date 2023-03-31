package author

type Author struct {
	ID        string  `json:"id" db:"id"`
	FullName  string  `json:"fullName" db:"full_name"`
	Pseudonym *string `json:"pseudonym" db:"pseudonym"`
	Specialty *string `json:"specialty" db:"specialty"`
}
