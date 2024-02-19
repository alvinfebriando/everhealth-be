package handler

import (
	"net/http"

	"github.com/alvinfebriando/project-batman-be/dto"
	"github.com/alvinfebriando/project-batman-be/usecase"
	"github.com/gin-gonic/gin"
)

type DoctorSpecialistHandler struct {
	doctorSpecialisUsecase usecase.DoctorSpecialistUsecase
}

func NewDoctorSpecialistHandler(u usecase.DoctorSpecialistUsecase) *DoctorSpecialistHandler {
	return &DoctorSpecialistHandler{doctorSpecialisUsecase: u}
}

func (h *DoctorSpecialistHandler) GetAllDoctorSpecialist(c *gin.Context) {
	doctorSpecialists, err := h.doctorSpecialisUsecase.FindAllDoctorSpecialist(c.Request.Context())
	if err != nil {
		_ = c.Error(err)
		return
	}
	doctorSpecialistsRes := []*dto.DoctorSpecialistRes{}
	for _, doctorSpecialist := range doctorSpecialists {
		doctorSpecialistres := dto.NewDoctorSpecialistres(doctorSpecialist)
		doctorSpecialistsRes = append(doctorSpecialistsRes, &doctorSpecialistres)
	}
	c.JSON(http.StatusOK, doctorSpecialistsRes)
}
