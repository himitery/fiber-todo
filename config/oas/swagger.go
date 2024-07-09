package oas

import (
	"fmt"
	"html/template"
	"path"
	"strings"
	"sync"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/gofiber/utils/v2"
	swaggerFiles "github.com/swaggo/files/v2"
	"github.com/swaggo/swag/v2"
)

const (
	defaultDocURL = "swagger.json"
	defaultIndex  = "index.html"
)

var HandlerDefault = New()

func New(config ...Config) fiber.Handler {
	cfg := configDefault(config...)

	index, err := template.New("index.html").Parse(indexTmpl)
	if err != nil {
		panic(fmt.Errorf("fiber: swagger middleware error -> %w", err))
	}

	var (
		prefix string
		once   sync.Once
		fs     = static.New("/docs", static.Config{
			FS: swaggerFiles.FS,
		})
	)

	return func(ctx fiber.Ctx) error {
		once.Do(
			func() {
				prefix = strings.ReplaceAll(ctx.Route().Path, "*", "")

				forwardedPrefix := getForwardedPrefix(ctx)
				if forwardedPrefix != "" {
					prefix = forwardedPrefix + prefix
				}

				// Set doc url
				if len(cfg.URL) == 0 {
					cfg.URL = path.Join(prefix, defaultDocURL)
				}
			},
		)

		switch ctx.Path(utils.CopyString(ctx.Params("*"))) {
		case defaultIndex:
			ctx.Type("html")
			return index.Execute(ctx, cfg)
		case defaultDocURL:
			var doc string
			if doc, err = swag.ReadDoc(cfg.InstanceName); err != nil {
				return err
			}
			return ctx.Type("json").SendString(doc)
		case "", "/":
			return ctx.Redirect().Status(fiber.StatusMovedPermanently).To(path.Join(prefix, defaultIndex))
		default:
			return fs(ctx)
		}
	}
}

func getForwardedPrefix(c fiber.Ctx) string {
	header := c.GetReqHeaders()["X-Forwarded-Prefix"]

	if len(header) == 0 {
		return ""
	}

	prefix := ""

	for _, rawPrefix := range header {
		endIndex := len(rawPrefix)
		for endIndex > 1 && rawPrefix[endIndex-1] == '/' {
			endIndex--
		}

		if endIndex != len(rawPrefix) {
			prefix += rawPrefix[:endIndex]
		} else {
			prefix += rawPrefix
		}
	}

	return prefix
}
