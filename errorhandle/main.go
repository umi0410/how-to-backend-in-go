package main

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

var (
	ErrResourceNotFound = errors.New("존재하지 않는 리소스")
	users = []*User{
		&User{Name: "우미잉", Age: 25, Username: "umi0410"},
	}
)

type User struct {
	Name string
	Age int
	Username string
}

type UserService struct{}

func (svc *UserService) GetUserWorst(username string) (*User, error) {
	for _, user := range users {
		if user.Username == username {
			return user, nil
		}
	}

	return nil, errors.New("유저가 존재하지 않는 유저(username=" + username + ")")
}

func (svc *UserService) GetUserUngracefully(username string) (*User, error) {
	for _, user := range users {
		if user.Username == username {
			return user, nil
		}
	}

	return nil, ErrResourceNotFound
}

func (svc *UserService) GetUserGracefully(username string) (*User, error) {
	for _, user := range users {
		if user.Username == username {
			return user, nil
		}
	}

	return nil, fmt.Errorf("입력된 username=%s: %w", username, ErrResourceNotFound)
}

func main() {
	userService := &UserService{}
	app := fiber.New(fiber.Config{})
	app.Use(customErrorHandler)

	app.Get("/worst/:username", func(ctx *fiber.Ctx) error {
		username := ctx.Params("username", "")
		user, err := userService.GetUserWorst(username)
		if err != nil {
			return err
		}

		return ctx.JSON(user)
	})

	app.Get("/ungraceful/:username", func(ctx *fiber.Ctx) error {
		username := ctx.Params("username", "")
		user, err := userService.GetUserUngracefully(username)
		if err != nil {
			return err
		}

		return ctx.JSON(user)
	})

	app.Get("/graceful/:username", func(ctx *fiber.Ctx) error {
		username := ctx.Params("username", "")
		user, err := userService.GetUserGracefully(username)
		if err != nil {
			return err
		}

		return ctx.JSON(user)
	})

	app.Listen(fmt.Sprintf("0.0.0.0:%d", 8000))
}

func customErrorHandler(ctx *fiber.Ctx) error {
	// 다음 핸들러를 호출한 뒤
	// 본 미들웨어의 작업은 사후처리!
	err := ctx.Next()
	if err != nil {
		log.Error(err)
		if errors.Is(err, ErrResourceNotFound) {
			// fiber의 Default Error Handler가 알아들을 수 있는 형태의 fiber Error을 리턴
			return fiber.NewError(404, err.Error())
		}
	}

	return err
}