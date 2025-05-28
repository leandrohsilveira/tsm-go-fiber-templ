package guards

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leandrohsilveira/tsm/dao"
	"github.com/rs/zerolog/log"
)

func AdminUserGuard(c *fiber.Ctx) error {
	logger := log.Ctx(c.UserContext()).With().Str("logger", "AdminUserGuard").Logger()

	info := GetCurrentUser(c)

	if info == nil {
		logger.Warn().Msg("Not authenticated")
		return redirectLogin(c)
	}

	if info.Role != dao.UserRoleAdminUser {
		logger.Warn().Str("role", string(info.Role)).Msgf("User role is not %s", dao.UserRoleAdminUser)
		return redirectLogin(c)
	}

	return c.Next()
}
