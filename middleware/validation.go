package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()

	// Register custom validation functions here if needed
	validate.RegisterValidation("albumid", validateAlbumID)
}

func validateAlbum() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "POST" || c.Request.Method == "PUT" {
			var album struct {
				ID     string  `json:"id" validate:"required,albumid"`
				Title  string  `json:"title" validate:"required"`
				Artist string  `json:"artist" validate:"required"`
				Price  float64 `json:"price" validate:"required,gt=0"`
			}

			if err := c.ShouldBindJSON(&album); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Invalid JSON format",
				})
				c.Abort()
				return
			}

			// Validate the struct
			if err := validate.Struct(album); err != nil {
				errors := make([]string, 0)

				for _, err := range err.(validator.ValidationErrors) {
					errors = append(errors, getValidationErrorMessage(err))
				}

				c.JSON(http.StatusBadRequest, gin.H{
					"error":   "Validation failed",
					"details": errors,
				})
				c.Abort()
				return
			}

			// Store validated data in context
			c.Set("validatedAlbum", album)
		}

		c.Next()
	}
}

// ValidateID middleware for ID parameter validation
func ValidateID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "ID parameter is required",
			})
			c.Abort()
			return
		}

		if !isValidAlbumID(id) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid ID format. ID must be alphanumeric and 1-10 characters long",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// Custom validation function for album ID
func validateAlbumID(fl validator.FieldLevel) bool {
	return isValidAlbumID(fl.Field().String())
}

// Helper function to validate album ID format
func isValidAlbumID(id string) bool {
	if len(id) < 1 || len(id) > 10 {
		return false
	}

	// Check if ID contains only alphanumeric characters
	for _, char := range id {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')) {
			return false
		}
	}

	return true
}

// Get user-friendly validation error messages
func getValidationErrorMessage(err validator.FieldError) string {
	field := err.Field()
	tag := err.Tag()

	switch tag {
	case "required":
		return field + " is required"
	case "min":
		if field == "Price" {
			return field + " must be greater than or equal to 0"
		}
		return field + " must be at least " + err.Param() + " characters long"
	case "max":
		if field == "Price" {
			return field + " must be less than or equal to " + err.Param()
		}
		return field + " must be at most " + err.Param() + " characters long"
	case "albumid":
		return field + " must be alphanumeric and 1-10 characters long"
	default:
		return field + " is invalid"
	}
}
