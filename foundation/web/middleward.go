package web

type Middleware func(Handler) Handler

func wrapMiddleware(mw []Middleware, handler Handler) Handler {
	for i := 0; i < len(mw); i++ {
		h := mw[i]
		if h != nil {
			handler = h(handler)
		}
	}
	return handler
}
