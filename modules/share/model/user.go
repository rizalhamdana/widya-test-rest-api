package model

// User is a model or entity in this project
type User struct {
	tableName struct{} `pg:"public.user"`
	ID        string   `json:"id" pg:"id"`
	Username  string   `json:"username" pg:"username"`
	Password  string   `json:"password,omitempty" pg:"password"`
	Email     string   `json:"email" pg:"email"`
	Gender    string   `json:"gender" pg:"gender"`
}

type StoredToken struct {
	tableName struct{} `pg:"public.stored_token"`
	IDUser    string   `json:"-" pg:"id_user"`
	Token     string   `json:"-" pg:"token"`
}
