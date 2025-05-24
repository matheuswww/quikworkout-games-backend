package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	admin_router "github.com/matheuswww/quikworkout-games-backend/src/routes/admin"
	edition_router "github.com/matheuswww/quikworkout-games-backend/src/routes/edition"
	user_router "github.com/matheuswww/quikworkout-games-backend/src/routes/user"
	participant_router "github.com/matheuswww/quikworkout-games-backend/src/routes/participant"
)

func InitRoutes(r *gin.RouterGroup, database *sql.DB) {
	r.Use(func(c *gin.Context) {
		c.Header("X-Frame-Options", "DENY")
		c.Header("Content-Security-Policy", "frame-ancestors 'none'")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("Cross-Origin-Resource-Policy", "same-origin")
		c.Header("Cross-Origin-Opener-Policy", "same-origin")
		c.Header("Referrer-Policy", "no-referrer")
		c.Next()
	})

	user_router.InitUserRoutes(r, database)
	admin_router.InitAdminRoutes(r, database)
	edition_router.InitEditionRoutes(r, database)
	participant_router.InitParticipantRoutes(r, database)
}
