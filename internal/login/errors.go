package login

type InvalidUserError string

func (err InvalidUserError) Error() string {
	return string(err)
}
func newInvalidUserError() InvalidUserError {
	return "invalid username or password"
}
