package router

import (
	"github.com/somashekhar-m25/task-manager-application/internal/database"
)

func StartApplication() {
	_, _ = database.ConnectToDB()
}
