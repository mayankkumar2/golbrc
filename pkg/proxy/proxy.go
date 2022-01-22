package proxy

import "github.com/valyala/fasthttp"

func ProxyMiddleware(ctx *fasthttp.RequestCtx) {
	ctx.Request.SetHost("127.0.0.1:8000")
	err := fasthttp.Do(&ctx.Request, &ctx.Response)
	if err != nil {
		return
	}
}

