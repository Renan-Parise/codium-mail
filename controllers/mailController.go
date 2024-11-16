package controllers

import (
	"fmt"
	"net/http"

	"github.com/Renan-Parise/mail/entities"
	"github.com/Renan-Parise/mail/services"
	"github.com/Renan-Parise/mail/utils"
	"github.com/gin-gonic/gin"
)

type EmailController struct {
	emailService services.EmailService
}

func NewEmailController() *EmailController {
	return &EmailController{
		emailService: services.NewEmailService(),
	}
}

func (ec *EmailController) Send(c *gin.Context) {
	var email entities.Email
	if err := c.ShouldBindJSON(&email); err != nil {
		utils.GetLogger().WithError(err).Error("Failed to bind JSON in SendEmail")
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid request payload: %s", err)})
		return
	}

	if err := email.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid email: %s", err)})
		return
	}

	err := ec.emailService.PublishEmail(email)
	if err != nil {
		utils.GetLogger().WithError(err).Error("Failed to publish email")
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to publish email: %s", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email request received with success"})
}
