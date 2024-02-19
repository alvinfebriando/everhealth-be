package handler

import (
	"net/http"

	"github.com/alvinfebriando/project-batman-be/dto"
	"github.com/alvinfebriando/project-batman-be/usecase"
	"github.com/gin-gonic/gin"
)

type DrugFormHandler struct {
	drugFormUsecase usecase.DrugFormUsecase
}

func NewDrugFormHandler(u usecase.DrugFormUsecase) *DrugFormHandler {
	return &DrugFormHandler{drugFormUsecase: u}
}

func (h *DrugFormHandler) GetAllDrugForm(c *gin.Context) {
	drugForms, err := h.drugFormUsecase.FindAllDrugForm(c.Request.Context())
	if err != nil {
		_ = c.Error(err)
		return
	}
	drugFormsRes := []*dto.DrugFormRes{}
	for _, drugForm := range drugForms {
		drugFormres := dto.NewDrugFormres(drugForm)
		drugFormsRes = append(drugFormsRes, &drugFormres)
	}
	c.JSON(http.StatusOK, drugFormsRes)
}
