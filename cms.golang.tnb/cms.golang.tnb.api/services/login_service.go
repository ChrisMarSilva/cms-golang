package service

import (
	"errors"

	dto "github.com/ChrisMarSilva/cms.golang.tnb.api/dtos"
	entity "github.com/ChrisMarSilva/cms.golang.tnb.api/entities"
	repository "github.com/ChrisMarSilva/cms.golang.tnb.api/repositories"
	"gorm.io/gorm"
)

type LoginService struct {
	bd   *gorm.DB
	repo repository.LoginRepository
}

func NewLoginService(bd *gorm.DB, repo repository.LoginRepository) *LoginService {
	return &LoginService{
		bd:   bd,
		repo: repo,
	}
}

func (s *LoginService) Entrar(txtEmail string, txtSenha string) (dto.UsuarioLocal, error) {

	var row entity.Usuario
	err := s.repo.GetEntrar(s.bd, &row)
	if err != nil {
		return dto.UsuarioLocal{}, err
	}

	if row.Situacao == "X" { // X-Bloqueado
		return dto.UsuarioLocal{}, errors.New("Usuário Bloqueado.")
	}

	if row.Situacao != "A" { // A-Ativo
		return dto.UsuarioLocal{}, errors.New("Usuário não está Ativo.")
	}

	// log.Println("SenhaBD: ", row.Senha)

	// // sum := sha256.Sum256([]byte(txtSenha))
	// // hexstring := fmt.Sprintf("%x", sum) //  sum[:]
	// // log.Println("hexstring: ", hexstring)

	// h := sha256.New()
	// h.Write([]byte(txtSenha))
	// senhaDigitada := hex.EncodeToString(h.Sum(nil))
	// log.Println("senhaDigitada: ", senhaDigitada)

	// if row.Senha != senhaDigitada { //  return self.senha == hashlib.sha256(senha.encode()).hexdigest()
	// 	log.Println("error: ", "Usuário/Senha Inválido.")
	// 	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Usuário/Senha Inválido."})
	// 	return
	// }

	user := dto.UsuarioLocal{
		ID:       row.ID,
		Nome:     row.Nome,
		Email:    row.Email,
		Foto:     row.Foto,
		ChatId:   row.ChatId,
		Tipo:     row.Tipo,
		Situacao: row.Situacao,
	}

	return user, nil
}
