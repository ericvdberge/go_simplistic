package middleware

import (
	"regexp"
	"strings"

	store "test/internal/api/store"
	services "test/pkg/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
)

var (
	anonymousUrls = []*regexp.Regexp{
		regexp.MustCompile("^/login$"),
		regexp.MustCompile("/static/*"),
	}
	errorMessage = "Login is not valid anymore. Please login again"
)

func validateAPIKey(c *fiber.Ctx, key string) (bool, error) {
	if services.LoginUserWithJWT(key) {
		return true, nil
	}
	return false, keyauth.ErrMissingOrMalformedAPIKey
}

func allowAnonymousUrls(c *fiber.Ctx) bool {
	originalURL := strings.ToLower(c.OriginalURL())
	for _, pattern := range anonymousUrls {
		if pattern.MatchString(originalURL) {
			return true
		}
	}
	return false
}

func HandleAuthMiddleware(ctx *fiber.Ctx) error {
	callback := keyauth.New(keyauth.Config{
		Next:      allowAnonymousUrls,
		KeyLookup: "cookie:access_token",
		Validator: validateAPIKey,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			c.ClearCookie("access_token")
			store.LoginStore.Error = &errorMessage
			return c.Redirect("/login")
		},
	})
	return callback(ctx)
}
