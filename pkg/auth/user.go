package auth

import (
	"errors"
	"fmt"

	"github.com/motephyr/longcare/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/session/v2"
	"github.com/motephyr/longcare/app"
	"github.com/motephyr/longcare/models"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func User(c *fiber.Ctx) (*models.User, error) {
	store := app.Http.Session.Get(c)
	userID := store.Get("user_id")
	if userID == nil {
		return nil, errors.New("User Not Logged In")
	}

	user, err := models.Users(Where("id = ?", userID)).OneG()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UserID(c *fiber.Ctx) int {
	store := app.Http.Session.Get(c)
	return store.Get("user_id").(int)
}

func IsLoggedIn(c *fiber.Ctx) bool {
	store := app.Http.Session.Get(c)
	userID := store.Get("user_id")
	if userID == nil {
		DeleteSession(store)
		c.ClearCookie()
		return false
	}
	token := c.Cookies("Verify-Rest-Token")
	if token == "" {
		tokenHash := store.Get("user_token")
		c.Cookie(&fiber.Cookie{
			Name:     "Verify-Rest-Token",
			Value:    fmt.Sprintf("%s", tokenHash),
			Secure:   false,
			HTTPOnly: true,
		})
	}

	return true
}

func Login(c *fiber.Ctx, userID int, secret string) (*config.Token, error) {
	store := app.Http.Session.Get(c) // get/create new session
	store.Set("user_id", userID)     // save to storage
	c.Locals("user_id", userID)
	token, err := app.Http.Token.CreateToken(c, userID, secret)
	if err == nil {
		store.Set("user_token", token.Hash)
		store.Set("token_expiry", token.Expire)
	}
	store.Save()
	return token, err
}

func Logout(c *fiber.Ctx) error {
	store := app.Http.Session.Get(c)
	DeleteSession(store)
	c.ClearCookie()
	return nil
}

func AuthCookie(c *fiber.Ctx) error {
	IsLoggedIn(c)
	return c.Next()
}

func DeleteSession(store *session.Store) {
	store.Delete("user_id")
	store.Delete("user_token")
	err := store.Save()
	if err != nil {
		panic(err)
	}
}

func CheckLogin(l models.User) (*models.User, error) {

	user, err := GetVerifiedUserByUsername(l.Username)
	if err != nil {
		return nil, errors.New("invalid Username or Password")
	}
	match, err := app.Http.Hash.Match(l.Password, user.Password)
	if !match {
		return nil, errors.New("invalid Username or Password")
	}
	return user, nil
}

func GetVerifiedUserByUsername(username string) (*models.User, error) {
	user, err := models.Users(Where("username = ?", username)).OneG()
	if err != nil {
		return nil, err
	}
	return user, nil //nolint:wsl
}
