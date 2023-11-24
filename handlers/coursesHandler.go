package handlers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Spacio-app/content-management-microservice/domain"
	"github.com/Spacio-app/content-management-microservice/services"
	"github.com/Spacio-app/content-management-microservice/utils"
	"github.com/gofiber/fiber/v2"
)

// Initialize the Cloudinary client using environment variables

// validar datos de entrada

type User struct {
	Name  string `json:"name"`
	Image string `json:"image"`
	Email string `json:"email"`
	// Add more fields as needed
}

func CreateCourse(c *fiber.Ctx) error {
	content := domain.CourseReq{}
	// if err := c.MultipartForm(&content) {
	// 	log.Println("Error al analizar el cuerpo de la solicitud:", err)
	// 	return err
	// }

	// fmt.Println("headers USER", User)

	title := c.FormValue("title")
	description := c.FormValue("description")

	// author := c.FormValue("author")

	// announcementStr := c.FormValue("announcement")
	// fmt.Println("announcement", announcementStr)
	content.Title = title
	content.Description = description
	// announcement, err := strconv.ParseBool(announcementStr)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return err
	// }
	announcement := c.FormValue("createAnnouncement")
	fmt.Println("announcement", announcement)
	content.CreateAnnouncement = announcement == "true"

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

	// // Handle the videos array
	content.Videos = []domain.VideoReq{}
	//subir miniatura
	URLMiniature, publicIDMiniature, err := uploadMiniature(c)
	if err != nil {
		content.Miniature = ""
	}
	content.Miniature = URLMiniature
	content.PublicIDMiniature = publicIDMiniature
	var firstVideoMiniatureURL string
	for i := 0; ; i++ {
		// //verificar si hay mas videos
		// if c.FormValue(fmt.Sprintf("videos[%d][title]", i)) == "" {
		// 	break
		// }
		//primera iteracion preguntar si hay miniatura

		titleKey := fmt.Sprintf("videos[%d][title]", i)
		descKey := fmt.Sprintf("videos[%d][desc]", i)
		urlKey := fmt.Sprintf("videos[%d][url]", i)
		miniatureKey := fmt.Sprintf("videos[%d][miniature]", i)

		title := c.FormValue(titleKey)
		desc := c.FormValue(descKey)
		file, err := c.FormFile(urlKey)
		miniatureFile, err := c.FormFile(miniatureKey)

		if title == "" && desc == "" && err != nil {
			// No more videos to process
			break
		}

		isVideo := true

		secureURLs, publicIDs, miniatureURL, PublicIDMiniature, err := utils.ProcessUploadedFiles(c, file, isVideo, miniatureFile)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error al procesar archivos",
			})
		}
		if i == 0 {
			// Esta es la primera iteración, así que establece la miniatura del primer video como la miniatura principal
			firstVideoMiniatureURL = miniatureURL
		}
		// Create a video struct to hold title, desc, and file info
		video := domain.VideoReq{
			Title:              title,
			Description:        desc,
			MiniatureVideo:     miniatureURL,
			PublicIDCloudinary: publicIDs,
			URL:                secureURLs,
			PublicIDMiniature:  PublicIDMiniature,
			// You may need to define a 'File' field in your Video struct
		}
		content.Videos = append(content.Videos, video)
		// videos = append(videos, video)
	}
	// If no miniature was uploaded, use the first video's miniature
	if content.Miniature == "" && firstVideoMiniatureURL != "" {
		content.Miniature = firstVideoMiniatureURL
	}
	//enviar a servicio
	if err := services.CreateCourse(content); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al crear el curso",
		})
	}

	if content.CreateAnnouncement {
		announcement := createAnnouncementFromCourse(content)
		if err := services.CreateAnnouncement(announcement); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error al crear el anuncio",
			})
		}
	}
	return c.JSON(content)

}

func GetAllCoursesHandler(c *fiber.Ctx) error {
	content, err := services.GetAllCourses()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al obtener los cursos",
		})
	}
	return c.JSON(content)
}

func UpdateCourseHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	content := domain.CourseReq{}
	if err := c.BodyParser(&content); err != nil {
		log.Println("Error al analizar el cuerpo de la solicitud:", err)
		return err
	}
	err := services.UpdateCourse(id, content)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al actualizar el curso",
		})
	}
	return c.JSON(content)
}

func getUserHeader(c *fiber.Ctx) (User, error) {
	UserHeader := c.Get("User")

	fmt.Println("UserHeader", UserHeader)
	var user User

	if err := json.Unmarshal([]byte(UserHeader), &user); err != nil {
		fmt.Println("Error:", err)
		return user, err
	}
	return user, nil
}
func uploadMiniature(c *fiber.Ctx) (string, string, error) {
	file, err := c.FormFile("miniature")
	if err != nil {
		return "", "", err
	}
	isVideo := false
	secureURLs, publicIDs, _, _, err := utils.ProcessUploadedFiles(c, file, isVideo, nil)
	if err != nil {
		return "", "", err
	}
	return secureURLs, publicIDs, nil
}
func createAnnouncementFromCourse(content domain.CourseReq) domain.FeedReq {
	announcement := domain.FeedReq{
		Title:       content.Title,
		Description: "Se ha creado un nuevo curso: " + content.Title + "\n \n \n" + "Descripcion: " + content.Description,
		Author:      content.Author,
		Comments:    []domain.FeedCommentsReq{},
	}
	return announcement
}
