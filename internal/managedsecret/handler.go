package managedsecret

import (
	"context"
	"fmt"
	"net/http"
	"smanager/internal/httputils"

	"github.com/gin-gonic/gin"
)

type ManagedSecretHandler struct {
	ManagedSecretServ IManagedSecretService
}
type IManagedSecretService interface {
	CreateManagedSecret(ctx context.Context, secretName, secretValue string) error
	ListManagedSecret(ctx context.Context) ([]ManagedSecretDto, error)
	GetSecret(ctx context.Context, name string) (string, error)
	EditManagedSecret(ctx context.Context, name, value string) error
}

func (msh *ManagedSecretHandler) CreateManagedSecret(c *gin.Context) {
	var request CreateManagedSecretRequest
	er1 := c.ShouldBindJSON(&request)
	var responseDto httputils.HttpResponseDto[*CreateManagedSecretResponse]
	if er1 != nil {
		fmt.Print(er1.Error())
		responseDto.ErrorMessage = "Bad request"
		c.JSON(http.StatusBadRequest, responseDto)
		return
	}
	err := msh.ManagedSecretServ.CreateManagedSecret(c.Request.Context(), request.Name, request.Value)
	if err != nil {
		fmt.Println(fmt.Errorf("ManagedSecretHandler %w", err))
		responseDto.ErrorMessage = "Error creating secret"
		c.JSON(http.StatusInternalServerError, responseDto)
		return
	}
	responseDto.Data = &CreateManagedSecretResponse{request.Name}
	c.JSON(http.StatusOK, responseDto)
}

func (msh *ManagedSecretHandler) ListManagedSecret(c *gin.Context) {
	secrets, err := msh.ManagedSecretServ.ListManagedSecret(c.Request.Context())
	if err != nil {
		fmt.Println(fmt.Errorf("ListManagedSecret %w", err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	responseDto := httputils.HttpResponseDto[[]ManagedSecretDto]{
		Data: secrets,
	}
	c.JSON(http.StatusOK, responseDto)

}

func (msh *ManagedSecretHandler) GetSecret(c *gin.Context) {
	secretName := c.Param("name")
	svalue, err := msh.ManagedSecretServ.GetSecret(c.Request.Context(), secretName)
	if err != nil {
		fmt.Println(fmt.Errorf("GetSecret %w", err))
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	responseDto := httputils.HttpResponseDto[string]{
		Data: svalue,
	}
	c.JSON(http.StatusOK, responseDto)

}

func (msh *ManagedSecretHandler) EditManagedSecret(c *gin.Context) {
	var request CreateManagedSecretRequest
	er1 := c.ShouldBindJSON(&request)
	if er1 != nil {
		fmt.Print(er1.Error())
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	err2 := msh.ManagedSecretServ.EditManagedSecret(c.Request.Context(), request.Name, request.Value)
	if err2 != nil {
		fmt.Print(er1.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, nil)
}

type CreateManagedSecretRequest struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type CreateManagedSecretResponse struct {
	Name string `json:"name"`
}
