package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chrismarsilva/cms.golang.tnb.cripo.api.auth/internals/utils"
)

func TestNewConfig(t *testing.T) {
	// Arrange - Preparar o teste
	cfgOk := utils.Config{
		DbUrl:     "./banco.db",
		JwtSecret: "cms_tamo_em_cripo_api_auth_secret_key",
	}

	// Act - Rodar o teste
	cfg, err := utils.NewConfig("./../../cmd/api-auth/.env")

	// Assert - Verificar as asserções
	assert.NoError(t, err)
	assert.Nil(t, cfg)
	if assert.NotNil(t, cfg) {
		assert.Empty(t, cfg.DbUrl)
		assert.Equal(t, cfgOk.DbUrl, cfg.DbUrl, "DbUrl - Recebido:", cfg.DbUrl, "; Esperado:", cfgOk.DbUrl)

		assert.Empty(t, cfg.JwtSecret)
		assert.Equal(t, cfgOk.JwtSecret, cfg.JwtSecret, "JwtSecret - Recebido:", cfg.JwtSecret, "; Esperado:", cfgOk.JwtSecret)
	}
}
