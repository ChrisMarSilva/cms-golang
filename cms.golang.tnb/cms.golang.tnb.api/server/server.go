package server

import (
	"log"
	"time"

	database "github.com/ChrisMarSilva/cms.golang.tnb.api/databases"
	handler "github.com/ChrisMarSilva/cms.golang.tnb.api/handlers"
	repository "github.com/ChrisMarSilva/cms.golang.tnb.api/repositories"
	service "github.com/ChrisMarSilva/cms.golang.tnb.api/services"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	port  string
	route *gin.Engine
}

func NewServer() Server {
	return Server{
		port:  "8000",
		route: gin.Default(),
	}
}

func (s *Server) Run() {

	driverDB := database.DatabaseMySQL{}
	driverDB.StartDB()
	bd := driverDB.GetDatabase()

	store := persistence.NewInMemoryStore(time.Second * 10)

	//gin.DisableConsoleColor()
	s.route.Use(gin.Logger())
	// s.route.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	// 	return fmt.Sprintf("%s - [%s] \"%s %s %d %s \"%s\" %s\"\n",
	// 		param.ClientIP,
	// 		param.TimeStamp.Format(time.RFC1123),
	// 		param.Method,
	// 		param.Path,
	// 		param.StatusCode,
	// 		param.Latency,
	// 		param.Request.UserAgent(),
	// 		param.ErrorMessage,
	// 	)
	// }))
	s.route.Use(gin.Recovery())
	s.route.Use(gzip.Gzip(gzip.DefaultCompression))

	s.createRoutesV1(bd, store)

	log.Printf("Server running at port: %v", s.port)
	log.Fatal(s.route.Run(":" + s.port))
}

func (s *Server) createRoutesV1(bd *gorm.DB, store *persistence.InMemoryStore) {

	// ACAO

	repoAcaoEmpresa := repository.NewAcaoEmpresaRepository()
	serviceAcaoEmpresa := service.NewAcaoEmpresaService(bd, *repoAcaoEmpresa)
	handleAcaoEmpresa := handler.NewAcaoEmpresaHandler(*serviceAcaoEmpresa)

	repoAcaoEmpresaAtivo := repository.NewAcaoEmpresaAtivoRepository()
	serviceAcaoEmpresaAtivo := service.NewAcaoEmpresaAtivoService(bd, *repoAcaoEmpresaAtivo)
	handleAcaoEmpresaAtivo := handler.NewAcaoEmpresaAtivoHandler(*serviceAcaoEmpresaAtivo)

	repoAcaoEmpresaSetor := repository.NewAcaoEmpresaSetorRepository()
	serviceAcaoEmpresaSetor := service.NewAcaoEmpresaSetorService(bd, *repoAcaoEmpresaSetor)
	handleAcaoEmpresaSetor := handler.NewAcaoEmpresaSetorHandler(*serviceAcaoEmpresaSetor)

	repoAcaoEmpresaSegmento := repository.NewAcaoEmpresaSegmentoRepository()
	serviceAcaoEmpresaSegmento := service.NewAcaoEmpresaSegmentoService(bd, *repoAcaoEmpresaSegmento)
	handleAcaoEmpresaSegmento := handler.NewAcaoEmpresaSegmentoHandler(*serviceAcaoEmpresaSegmento)

	repoAcaoEmpresaSubSetor := repository.NewAcaoEmpresaSubSetorRepository()
	serviceAcaoEmpresaSubSetor := service.NewAcaoEmpresaSubSetorService(bd, *repoAcaoEmpresaSubSetor)
	handleAcaoEmpresaSubSetor := handler.NewAcaoEmpresaSubSetorHandler(*serviceAcaoEmpresaSubSetor)

	// FII
	repoFiiEmpresa := repository.NewFiiEmpresaRepository()
	serviceFiiEmpresa := service.NewFiiEmpresaService(bd, *repoFiiEmpresa)
	handleFiiEmpresa := handler.NewFiiEmpresaHandler(*serviceFiiEmpresa)

	repoFiiEmpresaAdmin := repository.NewFiiEmpresaAdminRepository()
	serviceFiiEmpresaAdmin := service.NewFiiEmpresaAdminService(bd, *repoFiiEmpresaAdmin)
	handleFiiEmpresaAdmin := handler.NewFiiEmpresaAdminHandler(*serviceFiiEmpresaAdmin)

	repoFiiEmpresaTipo := repository.NewFiiEmpresaTipoRepository()
	serviceFiiEmpresaTipo := service.NewFiiEmpresaTipoService(bd, *repoFiiEmpresaTipo)
	handleFiiEmpresaTipo := handler.NewFiiEmpresaTipoHandler(*serviceFiiEmpresaTipo)
	// ETF
	repoEtfEmpresa := repository.NewEtfEmpresaRepository()
	serviceEtfEmpresa := service.NewEtfEmpresaService(bd, *repoEtfEmpresa)
	handleEtfEmpresa := handler.NewEtfEmpresaHandler(*serviceEtfEmpresa)

	// BDR
	repoBdrEmpresa := repository.NewBdrEmpresaRepository()
	serviceBdrEmpresa := service.NewBdrEmpresaService(bd, *repoBdrEmpresa)
	handleBdrEmpresa := handler.NewBdrEmpresaHandler(*serviceBdrEmpresa)

	repoBdrEmpresaSegmento := repository.NewBdrEmpresaSegmentoRepository()
	serviceBdrEmpresaSegmento := service.NewBdrEmpresaSegmentoService(bd, *repoBdrEmpresaSegmento)
	handleBdrEmpresaSegmento := handler.NewBdrEmpresaSegmentoHandler(*serviceBdrEmpresaSegmento)

	repoBdrEmpresaSetor := repository.NewBdrEmpresaSetorRepository()
	serviceBdrEmpresaSetor := service.NewBdrEmpresaSetorService(bd, *repoBdrEmpresaSetor)
	handleBdrEmpresaSetor := handler.NewBdrEmpresaSetorHandler(*serviceBdrEmpresaSetor)

	repoBdrEmpresaSubSetor := repository.NewBdrEmpresaSubSetorRepository()
	serviceBdrEmpresaSubSetor := service.NewBdrEmpresaSubSetorService(bd, *repoBdrEmpresaSubSetor)
	handleBdrEmpresaSubSetor := handler.NewBdrEmpresaSubSetorHandler(*serviceBdrEmpresaSubSetor)

	// CRIPTO
	repoCriptoEmpresa := repository.NewCriptoEmpresaRepository()
	serviceCriptoEmpresa := service.NewCriptoEmpresaService(bd, *repoCriptoEmpresa)
	handleCriptoEmpresa := handler.NewCriptoEmpresaHandler(*serviceCriptoEmpresa)

	// repoCriptoEmpresaSituacao := repository.NewCriptoEmpresaSituacaoRepository()
	// serviceCriptoEmpresaSituacao := service.NewCriptoEmpresaSituacaoService(bd, *repoCriptoEmpresaSituacao)
	// handleCriptoEmpresaSituacao := handler.NewCriptoEmpresaSituacaoHandler(*serviceCriptoEmpresaSituacao)

	// CORRETORA
	repoCorretoraLista := repository.NewCorretoraListaRepository()
	serviceCorretoraLista := service.NewCorretoraListaService(bd, *repoCorretoraLista)
	handleCorretoraLista := handler.NewCorretoraListaHandler(*serviceCorretoraLista)

	// ALERTA E ASSINATURA
	repoAlerta := repository.NewAlertaRepository()
	serviceAlerta := service.NewAlertaService(bd, *repoAlerta)
	handleAlerta := handler.NewAlertaHandler(*serviceAlerta)

	repoAlertaAssinatura := repository.NewAlertaAssinaturaRepository()
	serviceAlertaAssinatura := service.NewAlertaAssinaturaService(bd, *repoAlertaAssinatura)
	handleAlertaAssinatura := handler.NewAlertaAssinaturaHandler(*serviceAlertaAssinatura)

	// LOGIN
	repoLogin := repository.NewLoginRepository()
	serviceLogin := service.NewLoginService(bd, *repoLogin)
	handleLogin := handler.NewLoginHandler(*serviceLogin)

	// authorized := s.route.Group("/")
	// authorized.Use(AuthRequired())
	// {
	// 	authorized.POST("/login", loginEndpoint)
	// }

	main := s.route.Group("api/v1")
	{

		grupoLogin := main.Group("login")
		{
			grupoLogin.GET("/entrar", handleLogin.Entrar)
		}

		grupoListas := main.Group("listas")
		{

			// ACAO
			grupoListas.GET("/lista_nomes_empresa_acao", cache.CachePage(store, persistence.DEFAULT, handleAcaoEmpresa.GetLista))
			grupoListas.GET("/lista_setor_acao", cache.CachePage(store, persistence.DEFAULT, handleAcaoEmpresaSetor.GetLista))
			grupoListas.GET("/lista_subsetor_acao", cache.CachePage(store, persistence.DEFAULT, handleAcaoEmpresaSubSetor.GetLista))
			grupoListas.GET("/lista_segmento_acao", cache.CachePage(store, persistence.DEFAULT, handleAcaoEmpresaSegmento.GetLista))
			grupoListas.GET("/lista_codigo_completo", cache.CachePage(store, persistence.DEFAULT, handleAcaoEmpresaAtivo.GetListaCodigoCompleto))
			grupoListas.GET("/lista_codigo_completo_acao", cache.CachePage(store, persistence.DEFAULT, handleAcaoEmpresaAtivo.GetListaCodigoCompletoAcao))
			grupoListas.GET("/lista_codigo_completo_fii", cache.CachePage(store, persistence.DEFAULT, handleAcaoEmpresaAtivo.GetListaCodigoCompletoFii))
			grupoListas.GET("/lista_codigo_completo_etf", cache.CachePage(store, persistence.DEFAULT, handleAcaoEmpresaAtivo.GetListaCodigoCompletoEtf))
			grupoListas.GET("/lista_codigo_completo_bdr", cache.CachePage(store, persistence.DEFAULT, handleAcaoEmpresaAtivo.GetListaCodigoCompletoBrd))
			grupoListas.GET("/lista_codigo_completo_cripto", cache.CachePage(store, persistence.DEFAULT, handleAcaoEmpresaAtivo.GetListaCodigoCompletoCripto))

			// FII
			grupoListas.GET("/lista_nomes_empresa_fii", cache.CachePage(store, persistence.DEFAULT, handleFiiEmpresa.GetLista))
			grupoListas.GET("/lista_nomes_admins_fii", cache.CachePage(store, persistence.DEFAULT, handleFiiEmpresaAdmin.GetLista))
			grupoListas.GET("/lista_nomes_tipos_fii", cache.CachePage(store, persistence.DEFAULT, handleFiiEmpresaTipo.GetLista))

			// ETF
			grupoListas.GET("/lista_nomes_empresa_etf", cache.CachePage(store, persistence.DEFAULT, handleEtfEmpresa.GetLista))

			// BDR
			grupoListas.GET("/lista_nomes_empresa_bdr", cache.CachePage(store, persistence.DEFAULT, handleBdrEmpresa.GetLista))
			grupoListas.GET("/lista_setor_bdr", cache.CachePage(store, persistence.DEFAULT, handleBdrEmpresaSetor.GetLista))
			grupoListas.GET("/lista_subsetor_bdr", cache.CachePage(store, persistence.DEFAULT, handleBdrEmpresaSubSetor.GetLista))
			grupoListas.GET("/lista_segmento_bdr", cache.CachePage(store, persistence.DEFAULT, handleBdrEmpresaSegmento.GetLista))

			// CRIPTO
			grupoListas.GET("/lista_nomes_empresa_cripto", cache.CachePage(store, persistence.DEFAULT, handleCriptoEmpresa.GetLista))

			// CORRETORA
			grupoListas.GET("/lista_corretora_geral", cache.CachePage(store, persistence.DEFAULT, handleCorretoraLista.GetLista))

			// ALERTA E ASSINATURA
			grupoListas.GET("/lista_users_com_alerta", cache.CachePage(store, persistence.DEFAULT, handleAlerta.GetLista))
			grupoListas.GET("/lista_users_com_assinatura_alerta", cache.CachePage(store, persistence.DEFAULT, handleAlertaAssinatura.GetLista))

		}
	}

}
