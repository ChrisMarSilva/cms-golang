package controllers

import (
	"errors"
	"net/http"

	"cms.golang.tnb.api/database"
	"cms.golang.tnb.api/entities"
	"cms.golang.tnb.api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SituacaoRepo struct {
	Db *gorm.DB
}

func New() *SituacaoRepo {
	return &SituacaoRepo{Db: database.GetDatabase()}
}

func (repository *SituacaoRepo) GetSituacoes(c *gin.Context) {
	var situacoes []entities.CriptoEmpresaSituacao
	var situacaoModel models.CriptoEmpresaSituacaoModel
	err := situacaoModel.GetSituacoes(repository.Db, &situacoes)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	// log.Println("Situacoes", len(situacoes))
	// for _, situacao := range situacoes {
	// 	log.Println(situacao.ToString())
	// }
	c.JSON(http.StatusOK, situacoes)
}

func (repository *SituacaoRepo) GetSituacao(c *gin.Context) {
	codigo, _ := c.Params.Get("codigo")
	var situacao entities.CriptoEmpresaSituacao
	var situacaoModel models.CriptoEmpresaSituacaoModel
	err := situacaoModel.GetSituacao(repository.Db, &situacao, codigo)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, situacao)
}

// func (repository *SituacaoRepo) CreateSituacao(c *gin.Context) {
// 	var user entities.CriptoEmpresaSituacao
// 	c.BindJSON(&user)
// 	err := models.CreateSituacao(repository.Db, &user)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
// 		return
// 	}
// 	c.JSON(http.StatusOK, user)
// }
//
// func (repository *SituacaoRepo) UpdateSituacao(c *gin.Context) {
// 	var user models.CriptoEmpresaSituacao
// 	id, _ := c.Params.Get("id")
// 	err := models.GetSituacao(repository.Db, &user, id)
// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			c.AbortWithStatus(http.StatusNotFound)
// 			return
// 		}

// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
// 		return
// 	}
// 	c.BindJSON(&user)
// 	err = models.UpdateSituacao(repository.Db, &user)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
// 		return
// 	}
// 	c.JSON(http.StatusOK, user)
// }

// func (repository *SituacaoRepo) DeleteSituacao(c *gin.Context) {
// 	var user models.CriptoEmpresaSituacao
// 	id, _ := c.Params.Get("id")
// 	err := models.DeleteSituacao(repository.Db, &user, id)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "Situacao deleted successfully"})
// }
