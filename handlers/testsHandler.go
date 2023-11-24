package handlers

import (
	"log"

	"github.com/Spacio-app/content-management-microservice/domain"
	"github.com/Spacio-app/content-management-microservice/services"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTest(c *fiber.Ctx) error {
	content := domain.TestReq{}
	if err := c.BodyParser(&content); err != nil {
		log.Println("Error al analizar el cuerpo de la solicitud:", err)
		return err
	}
	user, error := getUserHeader(c)
	if error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al obtener el usuario",
		})
	}
	author := domain.AuthorReq{
		Name:  user.Name,
		Photo: user.Image,
		Email: user.Email,
	}
	// author := domain.AuthorReq{
	// 	Name:  "pruebaPostman",
	// 	Photo: "pruebaPostman",
	// }
	content.Author = author
	//enviar a servicio
	err := services.CreateTest(content)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al crear el test",
		})
	}
	if content.CreateAnnouncement {
		announcement := createAnnouncementFromTest(content)
		if err := services.CreateAnnouncement(announcement); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error al crear el anuncio",
			})
		}
	}
	return c.JSON(content)
}

// get all tests
func GetAllTestsHandler(c *fiber.Ctx) error {
	content, err := services.GetAllTests()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al obtener los tests",
		})
	}

	return c.JSON(content)
}

// update test
func UpdateTestHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	content := domain.TestReq{}
	if err := c.BodyParser(&content); err != nil {
		log.Println("Error al analizar el cuerpo de la solicitud:", err)
		return err
	}
	err := services.UpdateTest(id, content)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al actualizar el test",
		})
	}
	return c.JSON(content)
}

// calificar test
func CalificarTestHandler(c *fiber.Ctx) error {
	contentID := c.Params("contentID")

	content := domain.TestReq{}
	if err := c.BodyParser(&content); err != nil {
		log.Println("Error al analizar el cuerpo de la solicitud:", err)
		return err
	}
	user, error := getUserHeader(c)
	if error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al obtener el usuario",
		})
	}
	author := domain.AuthorReq{
		Name:  user.Name,
		Photo: user.Image,
		Email: user.Email,
	}
	// Suponiendo que tienes un TestResult con las respuestas del usuario
	userAnswers := content.UserAnswers

	// Suponiendo que tienes las respuestas correctas en el contenido del test
	correctAnswers := content.Questions

	numQuestions := len(correctAnswers)
	numCorrect := 0

	// Comparar las respuestas del usuario con las respuestas correctas
	for i, question := range correctAnswers {
		if i < len(userAnswers) {
			userAnswer := userAnswers[i]

			// Buscar la opción correcta en las opciones de la pregunta
			var correctOption *domain.Option
			for _, option := range question.Options {
				if option.IsCorrect {
					correctOption = &option
					break
				}
			}

			// Comparar la respuesta del usuario con la opción correcta
			if correctOption != nil && userAnswer.AnswerText == correctOption.OptionText {
				numCorrect++
				userAnswer.IsCorrect = true
			} else {
				userAnswer.IsCorrect = false
			}
		}
	}

	var calification float64
	var percentageCorrect float64
	// Calcular el porcentaje correcto y la calificación
	percentageCorrect = (float64(numCorrect) / float64(numQuestions)) * 100
	calification = calculateCalification(percentageCorrect)

	// Guardar la calificación en la base de datos u otro servicio
	// Puedes utilizar domain.TestResult o crear una nueva estructura de resultado

	result := domain.TestResultReq{
		ContentID:         contentID,
		Author:            author, // Reemplaza con el ID del usuario
		Calification:      calification,
		PercentageCorrect: percentageCorrect,
	}

	// Enviar a servicio para guardar el resultado
	services.SaveTestResult(result)

	// Devolver la calificación como respuesta
	return c.JSON(fiber.Map{
		"calification": calification,
	})
}

func calculateCalification(percentage float64) float64 {
	// Formula para calcular la calificación
	calification := float64(7 * (percentage / 100))

	// Limitar la calificación entre 1 y 7
	if calification < 1 {
		return 1
	} else if calification > 7 {
		return 7
	}

	return calification
}
func HasRatedTestHandler(c *fiber.Ctx) error {
	contentID := c.Params("contentID")

	user, err := getUserHeader(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al obtener el usuario",
		})
	}

	author := domain.AuthorReq{
		Name:  user.Name,
		Photo: user.Image,
		Email: user.Email,
	}

	objectID, err := primitive.ObjectIDFromHex(contentID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de contenido inválido",
		})
	}

	result, err := services.GetTestResult(objectID, author.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al obtener el resultado del test",
		})
	}

	hasRated := result > 0
	return c.JSON(fiber.Map{
		"hasRated": hasRated,
	})
}

func createAnnouncementFromTest(content domain.TestReq) domain.FeedReq {
	announcement := domain.FeedReq{
		Title:       content.Title,
		Description: "Se ha creado un nuevo Test, prueba tus conocimientos!: " + content.Title + "\n" + content.Description,
		Author:      content.Author,
	}
	return announcement
}
