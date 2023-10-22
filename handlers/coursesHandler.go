package handlers

import (
	"fmt"
	"log"

	"github.com/Spacio-app/content-management-microservice/domain"
	"github.com/Spacio-app/content-management-microservice/services"
	"github.com/Spacio-app/content-management-microservice/utils"
	"github.com/gofiber/fiber/v2"
)

// Initialize the Cloudinary client using environment variables

// validar datos de entrada

func CreateCourse(c *fiber.Ctx) error {
	content := domain.CourseReq{}
	// if err := c.MultipartForm(&content) {
	// 	log.Println("Error al analizar el cuerpo de la solicitud:", err)
	// 	return err
	// }
	title := c.FormValue("title")
	description := c.FormValue("description")
	author := c.FormValue("author")
	content.Title = title
	content.Description = description
	content.Author = author
	// // Handle the videos array
	content.Videos = []domain.VideoReq{}

	for i := 0; ; i++ {
		// //verificar si hay mas videos
		// if c.FormValue(fmt.Sprintf("videos[%d][title]", i)) == "" {
		// 	break
		// }
		titleKey := fmt.Sprintf("videos[%d][title]", i)
		descKey := fmt.Sprintf("videos[%d][desc]", i)
		urlKey := fmt.Sprintf("videos[%d][url]", i)

		title := c.FormValue(titleKey)
		desc := c.FormValue(descKey)
		file, err := c.FormFile(urlKey)

		if title == "" && desc == "" && err != nil {
			// No more videos to process
			break
		}

		isVideo := true

		secureURLs, publicIDs, miniatureURL, err := utils.ProcessUploadedFiles(c, file, isVideo)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error al procesar archivos",
			})
		}

		// Create a video struct to hold title, desc, and file info
		video := domain.VideoReq{
			Title:              title,
			Description:        desc,
			MiniatureVideo:     miniatureURL,
			PublicIDCloudinary: publicIDs,
			URL:                secureURLs,
			// You may need to define a 'File' field in your Video struct
		}
		content.Videos = append(content.Videos, video)
		// videos = append(videos, video)
	}
	//enviar a servicio
	if err := services.CreateCourse(content); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al crear el curso",
		})
	}
	return c.JSON(content)
}

// secureURLs, publicIDs, miniatureURL, err := utils.ProcessUploadedFiles(c, url, isVideo)
// if err != nil {
// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 		"error": "Error al procesar archivos",
// 	})
// }

// form, err := c.MultipartForm()
// if err != nil {
// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 		"error": "Error al analizar el formulario",
// 	})
// }

// // Accede a los campos del formulario, incluidos los datos de video
// title := form.Value["title"][0]             // Acceder al valor del campo "title"
// author := form.Value["author"][0]           // Acceder al valor del campo "author"
// description := form.Value["description"][0] // Acceder al valor del campo "description"
// // log.Printf("Título: %s, Autor: %s, Descripción: %s\n", title, author, description)
// // // Inicializar content.Videos como una lista vacía
// // videos := form.Value["videos"]
// type VideoData struct {
// 	Title       string `json:"title"`
// 	Description string `json:"desc"`
// 	URL         string `json:"url"`
// }

// // En tu función CreateCourse:

// // Recibir el campo "videosData" como JSON
// videosDataJSON := form

// log.Printf("Videos: %s\n", videosDataJSON)
// // Parsear el JSON para obtener los datos de video
// var videosData []VideoData
// // if err := json.Unmarshal([]byte(videosDataJSON), &videosData); err != nil {
// // 	// Manejar error de análisis JSON
// // }
// content.Videos = []domain.VideoReq{}
// for x := range videosData {
// 	log.Printf("Video %d: %s\n", x, videosData[x])
// }
// // obtener el contenido de los videos
// //videoFiles := form.File["videos"] // Acceder a los archivos de video
// if len(videosData) == 0 {
// 	log.Println("No se proporcionaron videos")
// } else {
// 	log.Printf("Se proporcionaron %d videos para el curso %s\n", len(videosData), title)

// 	for _, videoInfo := range videosData {
// 		isVideo := true

// 		videoTitle := videoInfo.Title
// 		videoDescription := videoInfo.Description
// 		videoURL := videoInfo.URL

// 		// Crear una nueva estructura de video
// 		videoReq := domain.VideoReq{
// 			Title:       videoTitle,
// 			Description: videoDescription,
// 		}

// 		// Procesar y cargar archivos
// 		secureURLs, publicIDs, miniatureURL, err := utils.ProcessUploadedFiles(c, videoURL, isVideo)
// 		if err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"error": "Error al procesar archivos",
// 			})
// 		}

// 		// Asignar los resultados a la estructura de video
// 		videoReq.URL = secureURLs[0]
// 		videoReq.PublicIDCloudinary = publicIDs[0]
// 		videoReq.MiniatureVideo = miniatureURL

// 		// Agregar el video a la lista de videos en content
// 		content.Videos = append(content.Videos, videoReq)
// 	}
// }

// // Asignar otros datos del formulario a la estructura content
// content.Title = title
// content.Author = author
// content.Description = description

// log.Println("Creando un nuevo curso...")

// if err := services.CreateCourse(content); err != nil {
// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 		"error": "Error al crear el curso",
// 	})
// }
// return c.JSON(content)
// }

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
