package middleware

type StoredToken struct {
	tableName struct{} `pg:"public.stored_token"`
	IDUser    string   `json:"-" pg:"id_user"`
	Token     string   `json:"-" pg:"token"`
}
