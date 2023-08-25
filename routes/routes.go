package routes

import (
	//"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
)

// Function to convert markdown to HTML
func markdownToHTML(markdown []byte) []byte {
	html := blackfriday.Run(markdown)
	return html
}

// Function to check if the face vector is valid
func isValidFaceVector(vector []float64) bool {
	// Check if the vector has exactly 128 dimensions
	return len(vector) == 128
}

// Route handler for showing the Service Presentation
func GetRoot() gin.HandlerFunc {
	return GetReference()
}

// Route handler for showing the Reference Documentation
func GetReference() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Read the content of the reference.md file
		content, err := os.ReadFile("./reference/reference.md")
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to read Reference file")
			return
		}

		// Convert Markdown to HTML
		htmlContent := markdownToHTML(content)

		// Set the appropriate headers and send the content as the response
		c.Header("Content-Type", "text/html charset=utf-8")
		c.String(http.StatusOK, string(htmlContent))
	}
}

// Route handler for getting the authorization response
func PostFRV() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Define the requestData struct
		var requestData struct {
			FaceVector []float64 `json:"face_vector"`
		}

		// Parse the JSON request body into the requestData struct
		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON",
			})
			return
		}

		// Check if the "face_vector" parameter exists
		if len(requestData.FaceVector) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Missing 'face_vector' parameter",
			})
			return
		}

		// Validate the dimensions of the face vector
		if !isValidFaceVector(requestData.FaceVector) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid face vector dimensions",
			})
			return
		}

		// Set the appropriate headers and send the content as the response
		c.Header("Content-Type", "application/json charset=utf-8")
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
}

