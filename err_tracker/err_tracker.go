// A module for tracking errors

package err_tracker

var Tracker map[string]*failedExpression

type failedExpression struct {
	endpoint  string
	frequency int
	errType   string
}
