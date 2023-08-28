package routes

import (
	"fmt"
	"net/http"
	"os"
	"encoding/json"
	"bytes"

	"github.com/gin-gonic/gin"
	"github.com/miguelamello/golang-faceid-api/database"
	"github.com/russross/blackfriday/v2"
)

// Define the requestData struct
type requestData struct {
	Vector []float64 `json:"vector"`
}

// Define the DBResponse struct
type DBPayload struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Role  string `json:"role"`
}

// Define the DBResponse struct
type DBResult struct {
	ID      string  `json:"id"`
	Version int     `json:"version"`
	Score   float64 `json:"score"`
	Payload DBPayload `json:"payload"`
	Vector  []float64 `json:"vector"`
}

// Define the DBResponse struct
type DBResponse struct {
	Result []DBResult `json:"result"`
	Status string   `json:"status"`
	Time   float64  `json:"time"`
}
 
// Define the APIResult struct
type APIResult struct {
	ID      string  `json:"id"`
	Payload DBPayload `json:"payload"`
}

// Define the APIResponse struct
type APIResponseFull struct {
	Result 	[]APIResult `json:"result"`
	Grant 	bool `json:"grant"`
}

// Define the APIResponse struct
type APIResponseEmpty struct {
	Result 	[]string `json:"result"`
	Grant 	bool `json:"grant"`
}

// Function to convert markdown to HTML
func markdownToHTML(markdown []byte) []byte {
	html := blackfriday.Run(markdown)
	return html
}

// Function to check if the vector is valid
func isValidVector(vector []float64) bool {
	// Check if the vector has exactly 128 dimensions
	return len(vector) == 128
}

// Function to adjust the result from database 
// and keep the API response consistent
func adjustResult(result string) string {
	var dbresponse DBResponse
	err := json.Unmarshal([]byte(result), &dbresponse)
	if err != nil {
		return ""
	}
	if len(dbresponse.Result) > 0 {
		var apiresponse APIResponseFull
		apiresponse.Result = make([]APIResult, 1)
		apiresponse.Result[0].ID = dbresponse.Result[0].ID
		apiresponse.Result[0].Payload = dbresponse.Result[0].Payload
		apiresponse.Grant = true
		resultBytes, err := json.Marshal(apiresponse)
		if err != nil {
			return ""
		}
		return string(resultBytes)
	} else if len(dbresponse.Result) == 0 {
		var apiresponse APIResponseEmpty
		apiresponse.Result = make([]string, 0)
		apiresponse.Grant = false
		resultBytes, err := json.Marshal(apiresponse)
		if err != nil {
			return ""
		}
		return string(resultBytes)
	}	else {
		return result
	}
}

// Route handler for showing the Reference Documentation
func GetReference(c *gin.Context) {
	// Read the content of the reference.md file
	content, err := os.ReadFile("./reference/reference.md")
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to read Reference file")
		return
	}

	// Convert Markdown to HTML
	htmlContent := markdownToHTML(content)

	// Read the content of the style.html file
	cssContent, err := os.ReadFile("./reference/style.html")
	if err != nil {
		cssContent = []byte("")
	}

	// Concatenate HTML and CSS content
	var concatenatedContent bytes.Buffer
	concatenatedContent.Write(cssContent)
	concatenatedContent.Write(htmlContent)

	htmlHead := "<!DOCTYPE html><html><head>"
	htmlHead += "<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">"
	htmlHead += "<title>FaceID Reference</title></head><body>"
	htmlTail := "</body></html>"
	htmlDoc := htmlHead + concatenatedContent.String() + htmlTail
	
	// Set the appropriate headers and send the content as the response
	c.Header("Content-Type", "text/html charset=utf-8")
	c.Header("Content-Length", fmt.Sprint(len(htmlDoc)))
	c.String(http.StatusOK, htmlDoc)
}

// Route handler for getting the authorization response
func SearchVector(c *gin.Context) {
	// Parse the JSON request body into the requestData struct
	var requestData requestData
	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON",
		})
		return
	}

	// Check if the "vector" parameter exists
	if len(requestData.Vector) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing 'vector' parameter",
		})
		return
	}

	// Validate the dimensions of the vector
	if !isValidVector(requestData.Vector) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid vector dimensions",
		})
		return
	}

	// Search the vector in the database
	result, err := db.SearchPoint(requestData.Vector)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to search vector",
		})
		return
	}

	// Set the appropriate headers and send the content as the response
	c.Header("Content-Type", "application/json charset=utf-8")
	c.String(http.StatusOK, adjustResult(result))
}

