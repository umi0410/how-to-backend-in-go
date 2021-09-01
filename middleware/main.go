package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	log "github.com/sirupsen/logrus"
	"time"
)

// injectLoggerMiddleware 는 요청 유저의 정보를 주입한 logger를
// 요청에 대한 context에 삽입합니다.
var injectLoggerMiddleware = func (ctx *fiber.Ctx) error{
	// basic auth middleware가 주입한 username을 이용합니다.
	username := ctx.Locals("username")
	logger := log.WithField("user", username)
	// context에 logger라는 키로 유저 정보를 주입한 logger를 전달합니다!
	ctx.Locals("logger", logger)

	// 다음 미들웨어(혹은 요청 핸들러)를 수행합니다.
	return ctx.Next()
}

func main(){
	// 두 가지 Config를 이용해 각각의 서버 8000, 8001에 띄웁니다.
	for i, config := range []*basicauth.Config{
		newBasicAuthConfigAlwaysAllow(),
		newBasicAuthConfigAllowOnlyAdmin(),
	}{
		app := fiber.New(fiber.Config{})
		// injectLoggerMiddleware는 basicauth가 context에 주입한 username을 이용하기 때문에
    	// 꼭 basicauth middleware 다음에 수행되어야합니다!
		app.Use(basicauth.New(*config))
		app.Use(injectLoggerMiddleware)
		// JWT Middleware
		app.Get("", func(ctx *fiber.Ctx) error {
			logger := ctx.Locals("logger").(*log.Entry)
			logger.Info("유저가 접속했습니다")
			return ctx.SendString("Welcome!\n")
		})

		go app.Listen(fmt.Sprintf("0.0.0.0:%d", 8000 +i))
	}
	for {
		time.Sleep(10 * time.Second)
	}
}

// newBasicAuthConfigAlwaysAllow 는 언제나 인증에 성공하는 Authorizer를
// 이용하는 Config를 만듭니다.
func newBasicAuthConfigAlwaysAllow() *basicauth.Config{
	return &basicauth.Config{
		Authorizer: func(username string, password string) bool {
			return true
		},
	}
}

// newBasicAuthConfigAlwaysAllow 는 Users에 존재하는 유저 정보에 대해서만
// 인증에 성공하는 Authorizer를 이용하는 Config를 만듭니다.
func newBasicAuthConfigAllowOnlyAdmin() *basicauth.Config{
	return &basicauth.Config{
		Users: map[string]string{
			"foo": "bar",
		},
	}
}

