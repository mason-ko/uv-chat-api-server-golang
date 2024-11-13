package ctx

import "context"

type Context struct {
	context.Context

	UserID int
}
