package api

import (
	"crypto/sha256"
	"crypto/subtle"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
)

var (
	apiKey        = "correct horse battery staple"
	anonymousUrls = []*regexp.Regexp{
		regexp.MustCompile("^/login$"),
		regexp.MustCompile("/static/*"),
	}
)

func validateAPIKey(c *fiber.Ctx, key string) (bool, error) {
	hashedAPIKey := sha256.Sum256([]byte(apiKey))
	hashedKey := sha256.Sum256([]byte(key))

	if subtle.ConstantTimeCompare(hashedAPIKey[:], hashedKey[:]) == 1 {
		return true, nil
	}
	return false, keyauth.ErrMissingOrMalformedAPIKey
}

func authFilter(c *fiber.Ctx) bool {
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
		Next:      authFilter,
		KeyLookup: "cookie:access_token",
		Validator: validateAPIKey,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Redirect("/login")
		},
	})
	return callback(ctx)
}
