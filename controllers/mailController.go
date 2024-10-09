package controllers

import (
	"net/http"

	"github.com/Renan-Parise/codium-mail/entities"
	"github.com/Renan-Parise/codium-mail/services"
	"github.com/Renan-Parise/codium-mail/utils"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := email.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ec.emailService.PublishEmail(email)
	if err != nil {
		utils.GetLogger().WithError(err).Error("Failed to publish email")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email request received"})
}
