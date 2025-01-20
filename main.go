package main

// написать сервер из уроков

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Message struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// var messages []Message
var messages = make(map[int]Message)
var nextID = 1

func GetHandler(c echo.Context) error {
	var msgSlice []Message
	for _, msg := range messages {
		msgSlice = append(msgSlice, msg)
	}
	return c.JSON(http.StatusOK, &messages)
}

func PostHandler(c echo.Context) error {
	var message Message

	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not add the message",
		})
	}

	message.ID = nextID

	messages[message.ID] = message
	// messages = append(messages, message)

	nextID++

	return c.JSON(http.StatusOK, Response{
		Status:  "ok",
		Message: "Success",
	})
}

func PatchHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Неверный id",
		})
	}

	var updatedMessage Message
	if err := c.Bind(&updatedMessage); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not update the message",
		})
	}

	// for i, message := range messages {
	// 	if message.ID == id {
	// 		updatedMessage.ID = id
	// 		messages[i] = updatedMessage
	// 		updated = true
	// 		break
	// 	}
	// }

	if _, exists := messages[id]; !exists {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Message was not found",
		})
	}

	updatedMessage.ID = id
	messages[id] = updatedMessage

	return c.JSON(http.StatusOK, Response{
		Status:  "ok",
		Message: "Message was updated",
	})
}

func DeleteHandle(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Неверный id",
		})
	}

	if _, exist := messages[id]; !exist {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Message was not found",
		})
	}
	delete(messages, id)
	return c.JSON(http.StatusOK, Response{
		Status:  "ok",
		Message: "Message was successfully deleted",
	})
}

// var counter int
// func GetHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodGet {
// 		fmt.Println(w, "Counter равен", strconv.Itoa(counter))
// 	} else {
// 		fmt.Println(w, "Поддерживается только GET")
// 	}
// }
// func PostHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodPost {
// 		counter++
// 		fmt.Println(w, "Counter увеличен на 1")
// 	} else {
// 		fmt.Println(w, "Поддерживается только метод POST")
// 	}
// }

func main() {
	e := echo.New()

	e.GET("/messages", GetHandler)
	e.POST("/messages", PostHandler)
	e.PATCH("/messages/:id", PatchHandler)
	e.DELETE("/messages/:id", DeleteHandle)

	e.Start(":8080")

	// http.HandleFunc("/get", GetHandler)
	// http.HandleFunc("/post", PostHandler)
	// http.ListenAndServe("localhost:8080", nil)
}

////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////

// package main

// import (
// 	"log"

// 	"github.com/IbadT/go_server/database"
// 	"github.com/gofiber/fiber/v2/middleware/limiter"
// 	"github.com/gofiber/fiber/v2/middleware/recover"
// 	"github.com/gofiber/fiber/v3"
// 	"github.com/gofiber/fiber/v3/middleware/compress"
// 	"github.com/gofiber/fiber/v3/middleware/logger"
// )

// func main() {

// 	if err := database.Connect(); err != nil {
// 		log.Fatalf("Ошибка подключения к базе данных: %v", err)
// 	}
// 	app := fiber.New(fiber.Config{
// 		// // Prefork позволяет создавать несколько процессов.
// 		// Prefork:       true,  // включаем предварительное форкование для увеличения производительности на многоядерных процессорах
// 		ServerHeader:  "Fiber", // добавляем заголовок для идентификации сервера
// 		CaseSensitive: true,    // включаем чувствительность к регистру в URL
// 		StrictRouting: true,    // включаем строгую маршрутизацию
// 	})

// 	// Подключаем middleware
// 	app.Use(logger.New())   // Логирование запросов
// 	app.Use(compress.New()) // Сжатие ответов
// 	app.Use(recover.New())  // Восстановление после паники
// 	app.Use(limiter.New())  // Лимит запросов для предотвращения DDOS атак

// 	routes.RegisterProductRoutes(app)

// 	//

// 	// Можно легко добавить middleware для обработки авторизации, логирования или различной защиты.
// 	// app.Use(func(c fiber.Ctx) error {
// 	// 	println("Запрос получен")
// 	// 	return c.Next() // передаем управление дальше
// 	// })

// 	// app.Get("/", func(c fiber.Ctx) error {
// 	// 	return c.SendString("Hello, Fiber!")
// 	// })

// 	// app.Get("/users/:id", func(c fiber.Ctx) error {
// 	// 	id := c.Params("id")
// 	// 	return c.SendString("User ID: " + id)
// 	// })

// 	// app.Get("/async", func(c fiber.Ctx) error {
// 	// 	go func() {
// 	// 		time.Sleep(2 * time.Second)
// 	// 		println("Асинхронная задача завершена")
// 	// 	}()
// 	// 	return c.SendString("Запрос принят, задача выполняется в фоне!")
// 	// })

// 	// app.Post("/submit", func(c fiber.Ctx) error {
// 	// 	data := new(struct {
// 	// 		Name string `json:"name"`
// 	// 	})
// 	// 	if err := c.Bind().Body(data); err != nil {
// 	// 		return err
// 	// 	}
// 	// 	return c.JSON(fiber.Map{"message": "Привет, " + data.Name})
// 	// })

// 	// Одним из больших плюсов Fiber является его совместимость с низкоуровневыми возможностями Go.
// 	// Fiber предоставляет интерфейс для работы с нативными net/http хендлерами,
// 	// что позволяет комбинировать его с уже существующими решениями:
// 	// httpHandler := func(w http.ResponseWriter, r *http.Request) {
// 	// 	w.Write([]byte("Привет из net/http"))
// 	// }
// 	// app.Get("/legacy", func(c fiber.Ctx) error {
// 	// 	httpHandler(c.Context().Response().Writer, c.Context().Request())
// 	// 	return nil
// 	// })

// 	// Query-параметры, которые передаются в строке запроса (например, ?sort=desc), также легко извлекаются через Fiber:
// 	// app.Get("/query", func(c fiber.Ctx) error {
// 	// 	query := c.Query("q", "default") // получаем значение параметра "q", задаём "default" как значение по умолчанию
// 	// 	return c.SendString("Search string: " + query)
// 	// })

// 	// Если нужно работать с заголовками запроса, Fiber имеет удобный API для их извлечения:
// 	// app.Get("/headers", func(c fiber.Ctx) error {
// 	// 	userAgent := c.Get("User-Agent") // извлекаем заголовок User-Agent
// 	// 	return c.SendString("Your User-Agent is: " + userAgent)
// 	// })

// 	app.Listen(":8080", fiber.ListenConfig{EnablePrefork: true}) // EnablePrefork - это Prefork в v3
// }

// // package main

// // import (
// // 	"log"

// // 	"github.com/IbadT/go_server/database"
// // 	"github.com/IbadT/go_server/routes"
// // 	"github.com/gofiber/fiber/v3"
// // 	"github.com/gofiber/fiber/v3/middleware/logger"
// // )

// // type User struct {
// // 	ID       string `json:"id"`
// // 	Name     string `json:"name"`
// // 	Login    string `json:"login"`
// // 	Password string `json:"password"`
// // }

// // // var users = []User{
// // // 	{ID: "1", Name: "Eduard", Login: "ibadtoff@gmail.com", Password: "gts530200"},
// // // }

// // func main() {
// // 	app := fiber.New()

// // 	database.Connect()

// // 	// middleware
// // 	app.Use(logger.New())

// // 	routes.Setup(app)

// // 	// // xh http://localhost:3000/api
// // 	// app.Get("/api", func(c fiber.Ctx) error {
// // 	// 	return c.SendString("Hello, World 👋!")
// // 	// })

// // 	// // xh http://localhost:3000/api/Eduard
// // 	// app.Get("/api/:name", func(c fiber.Ctx) error {
// // 	// 	msg := fmt.Sprintf("Hello %s 👋", c.Params("name"))
// // 	// 	return c.SendString(msg)
// // 	// })

// // 	// app.Get("/api/users", getUsers)

// // 	// middleware
// // 	// app.Use(func(c fiber.Ctx) error {
// // 	// 	return c.SendStatus(404)
// // 	// })

// // 	// app.Listen(":3000")
// // 	log.Fatal(app.Listen(":3000"))
// // }

// // // go mod init github.com/your/repo
// // // go get -u github.com/gofiber/fiber/v3
// // // https://github.com/IbadT/go_server.git
// // // go get -u gorm.io/gorm
// // // go get -u gorm.io/driver/postgres
