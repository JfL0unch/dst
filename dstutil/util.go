package dstutil

import "github.com/JfL0unch/dst"

// Unparen returns e with any enclosing parentheses stripped.
func Unparen(e dst.Expr) dst.Expr {
	for {
		p, ok := e.(*dst.ParenExpr)
		if !ok {
			return e
		}
		e = p.X
	}
}
