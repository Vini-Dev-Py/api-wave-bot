package domain

// User representa o modelo de domínio do usuário.
type User struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"-"`            // campo usado apenas para criar usuário
	PasswordHash string `json:"passwordHash"` // senha criptografada vinda do banco
	CompanyID    string `json:"companyId"`
}
