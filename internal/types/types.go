package types

type Students struct {
	ID    string
	Name  string `validate:"required"`
	Age   int    `validate:"required"`
	Email string `validate:"required"`
}
