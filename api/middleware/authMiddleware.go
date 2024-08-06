package middleware

import (
	"regexp"
	"strings"

	services "test/api/services"
	store "test/api/store"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
)

var (
	anonymousUrls = []*regexp.Regexp{
		regexp.MustCompile("^/login$"),
		regexp.MustCompile("/static/*"),
	}
	errorMessage = "Login timed out. Please login again"
)

func validateAPIKey(c *fiber.Ctx, key string) (bool, error) {
	if services.LoginUserWithJWT(key) {
		return true, nil
	}
	return false, keyauth.ErrMissingOrMalformedAPIKey
}

func applyAuthOnUrls(c *fiber.Ctx) bool {
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
		Next:      applyAuthOnUrls,
		KeyLookup: "cookie:access_token",
		Validator: validateAPIKey,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			store.LoginStore.Error = &errorMessage
			return c.Redirect("/login")
		},
	})
	return callback(ctx)
}
