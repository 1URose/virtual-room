package handlers

import (
	"git.ai-space.tech/coursework/backend/internal/application"
	"git.ai-space.tech/coursework/backend/internal/presentation/http/sponsors/sponsor_dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SponsorHandler struct {
	Service *application.SponsorService
}

func NewHandler(service *application.SponsorService) *SponsorHandler {
	return &SponsorHandler{Service: service}
}

// CreateSponsor handles the creation of a new sponsor.
// @Summary Create a sponsor
// @Description Creates a new sponsor with the provided details
// @Tags sponsors
// @Accept json
// @Produce json
// @Param sponsor body sponsor_dto.CreateSponsor true "Sponsor details"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /sponsors/create [post]
func (s *SponsorHandler) CreateSponsor(ctx *gin.Context) {
	dto := sponsor_dto.CreateSponsor{}

	if err := ctx.ShouldBindJSON(&dto); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
	}

	if err := s.Service.Create(&dto); err != nil {

		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, gin.H{"sponsor creation is successful": dto.Name})
}

// GetSponsorByName retrieves a sponsor by its name.
// @Summary Get sponsor by name
// @Description Fetches the sponsor details by name
// @Tags sponsors
// @Produce json
// @Param name path string true "Sponsor name"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /sponsors/{name} [get]
func (s *SponsorHandler) GetSponsorByName(ctx *gin.Context) {
	name := ctx.Param("name")

	sponsor, err := s.Service.GetSponsor(name)

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, err)

		return
	}

	ctx.JSON(http.StatusOK, sponsor)
}

// GetAllSponsors retrieves all sponsors.
// @Summary Get all sponsors
// @Description Fetches a list of all sponsors
// @Tags sponsors
// @Produce json
// @Success 200 {array} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /sponsors/all [get]
func (s *SponsorHandler) GetAllSponsors(ctx *gin.Context) {
	sponsors, err := s.Service.GetAllSponsors()

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, err)

		return
	}

	ctx.JSON(http.StatusOK, sponsors)
}

// DeleteSponsor deletes a sponsor by its name.
// @Summary Delete a sponsor
// @Description Deletes the sponsor specified by the name
// @Tags sponsors
// @Produce json
// @Param name path string true "Sponsor name"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /sponsors/{sponsorName} [delete]
func (s *SponsorHandler) DeleteSponsor(ctx *gin.Context) {
	sponsorName := ctx.Param("sponsorName")

	if err := s.Service.DeleteSponsor(sponsorName); err != nil {

		ctx.JSON(http.StatusInternalServerError, err)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"sponsor deletion is successful": sponsorName})
}

// UpdateSponsor updates a sponsor's details.
// @Summary Update a sponsor
// @Description Updates the sponsor information such as name, contact, and event association
// @Tags sponsors
// @Accept json
// @Produce json
// @Param name path string true "Sponsor name"
// @Param sponsor body map[string]interface{} true "Sponsor updates"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /sponsors/{name} [put]
func (s *SponsorHandler) UpdateSponsor(ctx *gin.Context) {
	name := ctx.Param("name")

	existingSponsor, err := s.Service.GetSponsor(name)

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, err)

		return
	}

	var input map[string]interface{}

	if err = ctx.ShouldBindJSON(&input); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})

		return
	}

	if newName, ok := input["new_name"].(string); ok {
		existingSponsor.Name = newName
	}
	if newContactInfo, ok := input["new_contact_info"].(string); ok {
		existingSponsor.ContactInfo = newContactInfo
	}
	if ContributionAmount, ok := input["new_contribution"].(float64); ok {
		existingSponsor.ContributionAmount = ContributionAmount
	}
	if newEventName, ok := input["new_event_name"].(string); ok {
		event, err := s.Service.EventRepo.GetEventByName(newEventName)

		if err != nil {

			ctx.JSON(http.StatusInternalServerError, err)

			return
		}
		existingSponsor.EventID = event.ID
	}

	if err = s.Service.UpdateSponsor(existingSponsor); err != nil {

		ctx.JSON(http.StatusInternalServerError, err)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"sponsor update is successful": name})

}
