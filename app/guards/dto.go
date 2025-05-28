package guards

import "github.com/leandrohsilveira/tsm/dao"

type CurrentUserDto struct {
	ID    string       `json:"id"`
	Name  string       `json:"name"`
	Email string       `json:"email"`
	Role  dao.UserRole `json:"role"`
}
