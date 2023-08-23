package main

import (
	"iwexlmsapi/routes/auth"
	"iwexlmsapi/routes/branch"
	auth.SetupAuthRoute(mainRouter)
	branch.SetupBranchRoutes(&mainRouter)
}
