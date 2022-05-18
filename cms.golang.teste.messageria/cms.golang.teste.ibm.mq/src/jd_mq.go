package src

import (
	"encoding/hex"
	ibmmq "github.com/ibm-messaging/mq-golang/v5/ibmmq"
	_ "log"
	"strconv"
)

// https://github.com/ibm-messaging/mq-golang/blob/master/ibmmq/cmqc_linux_amd64.go

type JDMQ struct {
	conn    ibmmq.MQQueueManager
	fila    ibmmq.MQObject
	mqod    *ibmmq.MQOD
	pmo     *ibmmq.MQPMO
	mqmd    *ibmmq.MQMD
	gmo     *ibmmq.MQGMO
	buffer  []byte
	datalen int
	err     error // ibmmq.MQReturn
}

func (m *JDMQ) Conectar(mgrName string) bool {
	m.conn, m.err = ibmmq.Conn(mgrName)
	return m.err == nil
}

func (m *JDMQ) Desconectar() bool {
	empty := ibmmq.MQQueueManager{}
	if m.conn == empty {
		return true
	}
	m.Rollback()
	m.FecharFila()
	m.err = m.conn.Disc()
	return m.err == nil
}

func (m *JDMQ) Commit() bool {
	m.err = m.conn.Cmit()
	return m.err == nil
}

func (m *JDMQ) Rollback() bool {
	m.err = m.conn.Back()
	return m.err == nil
}

func (m *JDMQ) abrirFila(objName string, openOptions int32) bool {

	m.FecharFila()

	m.mqod = ibmmq.NewMQOD()
	m.mqod.ObjectType = ibmmq.MQOT_Q
	m.mqod.ObjectName = objName

	m.fila, m.err = m.conn.Open(m.mqod, openOptions)
	return m.err == nil
}

func (m *JDMQ) AbrirFilaPut(objName string) bool {
	var openOptions int32
	openOptions = ibmmq.MQOO_OUTPUT + ibmmq.MQOO_FAIL_IF_QUIESCING
	openOptions |= ibmmq.MQOO_INPUT_AS_Q_DEF
	return m.abrirFila(objName, openOptions)
}

func (m *JDMQ) AbrirFilaGet(objName string) bool {
	var openOptions int32
	openOptions = ibmmq.MQOO_INPUT_EXCLUSIVE
	return m.abrirFila(objName, openOptions)
}

func (m *JDMQ) FecharFila() bool {
	m.Rollback()
	empty := ibmmq.MQObject{}
	if m.fila == empty {
		return true
	}
	m.err = m.fila.Close(0)
	return m.err == nil
}

func (m *JDMQ) EnviarMensagem(message string) bool {

	m.buffer = []byte(message)

	m.mqmd = ibmmq.NewMQMD()
	m.mqmd.Format = ibmmq.MQFMT_STRING

	// empty := ibmmq.MQPMO{}
	if m.pmo != new(ibmmq.MQPMO) {
		m.pmo = ibmmq.NewMQPMO()
		m.pmo.Options = ibmmq.MQPMO_SYNCPOINT
	}

	m.err = m.fila.Put(m.mqmd, m.pmo, m.buffer)
	return m.err == nil
}

func (m *JDMQ) ReceberMensagem() bool {

	m.buffer = make([]byte, 1024)

	m.mqmd = ibmmq.NewMQMD()
	m.mqmd.Format = ibmmq.MQFMT_STRING

	// empty := ibmmq.MQGMO{}
	//if m.pmo != new(ibmmq.MQGMO) {
	m.gmo = ibmmq.NewMQGMO()
	m.gmo.Options = ibmmq.MQGMO_SYNCPOINT
	m.gmo.Options |= ibmmq.MQGMO_NO_WAIT
	// m.gmo.Options |= ibmmq.MQGMO_WAIT
	// m.gmo.WaitInterval = 100 // 3 * 1000 // The WaitInterval is in milliseconds
	//}

	m.datalen, m.err = m.fila.Get(m.mqmd, m.gmo, m.buffer)

	if m.err != nil {
		if m.err.(*ibmmq.MQReturn).MQRC == ibmmq.MQRC_NO_MSG_AVAILABLE {
			m.err = nil
		}
		return false
	}

	return true
}

func (m *JDMQ) GetMessage() string {
	return string(m.buffer[:m.datalen]) // strings.TrimSpace(string(m.buffer[:m.datalen]))
}

func (m *JDMQ) GetMsgId() string {
	return hex.EncodeToString(m.mqmd.MsgId) // putmqmd
}

func (m *JDMQ) GetErr() error {
	return m.err
}

func (m *JDMQ) GetCode() string {
	if m.err != nil {
		rcInt := int(m.err.(*ibmmq.MQReturn).MQRC)
		return strconv.Itoa(rcInt) // int to string
	}
	return ""
}

func (m *JDMQ) GetReason() string {
	if m.err != nil {
		rcInt := int(m.err.(*ibmmq.MQReturn).MQRC)
		return ibmmq.MQItoString("RC", rcInt)
	}
	return "" // m.err.MQCC != ibmmq.MQCC_OK
}

/*





----------------------------------------------------------------------------------------------------------------
----------------------------------------------------------------------------------------------------------------


unit uTJDMQFila;

  tpAberturaFila = (afSomentePut, afSomenteGet, afPutGet);

  TJDMQFila = class(TComponent)
  private
    FNomeFila: string;
    FCodigoErro: Integer;
    FNomeErro: string;
    FDetalheErro: string;
    FInterfaceMQ: tpInterfaceMQ;

    FMQAX_MQSession: MQSession;
    FMQAX_QMgr: MQQueueManager;
    FMQAX_QueueFila: MQQueue;

    FMQI_Hconn: MQHCONN;
    FMQI_ObjDesc: MQOD;
    FMQI_Hobj: MQHOBJ;

    FJDMQ_ConnID: AnsiString;

    FQueueManager: string;
    FQueueManagerOpen: string;
    FQueueManagerRemoto: string;
    FUserIDOpen: string;
    FInTransacao: Boolean;
    FFilaAberta: Boolean;
    FConectado: Boolean;
    FTipoAbertura: tpAberturaFila;
    FJDTraceApp: TJDTraceApp;
    FCabecalhoMsgEstendido: Boolean;
    FGETIntervaloMSeg: Integer;
    FChannel: String;
    FPort: Word;
    FServer: String;

    // SSL
    FSSLPeerName: AnsiString;
    FSSLFipsRequired: Boolean;
    FUtilizarSSL: Boolean;
    FCipherSuite: AnsiString;
    FChannelName: AnsiString;
    FCipherSpec: AnsiString;
    FKeyResetCount: Integer;
    FSSLCertRevocationCheck: Boolean;
    FSSLKeyRepository: AnsiString;


    Mensagem: TJDMQMensagem;
    function Conectar(QMgr: AnsiString): Boolean;
    function Desconectar: Boolean;
    function AbrirFila(const NomeFila: AnsiString; TipoAbertura: tpAberturaFila; QueueManager: AnsiString = ''; QueueManagerRemoto: AnsiString = ''; UserID: AnsiString = ''): Boolean; overload;
    function AbrirFila(const NomeFila: AnsiString; OpcoesAbertura: Integer; QueueManager: AnsiString = ''; QueueManagerRemoto: AnsiString = ''; UserID: AnsiString = ''): Boolean; overload;
    function FecharFila(): Boolean;
    function RollBack(GeraException: Boolean = True): Boolean;
    function Commit(GeraException: Boolean = True): Boolean;
    function EnviarMensagem(EnvMensagemID: string = ''; EnvMsgRelacionadaID: string = ''): Boolean; overload;
    function EnviarMensagem(OpcoesEnvio: Integer; EnvMensagemID: string = ''; EnvMsgRelacionadaID: string = ''): Boolean; overload;
    function ReceberMensagem(RecMensagemID: string = ''; RecMsgRelacionadaID: string = ''): Boolean; overload;
    function ReceberMensagem(OpcoesRecebimento: Integer; RecMensagemID: string = ''; RecMsgRelacionadaID: string = ''): Boolean; overload;
    function inTransaction: Boolean;
    function DescrReasonCode(ReasonCode: Integer): string;

    property TraceApp: TJDTraceApp read FJDTraceApp write FJDTraceApp;
    property QueueManager: string read GetQueueManager;
    property CodigoErro: Integer read FCodigoErro;
    property NomeErro: string read FNomeErro;
    property DetalheErro: string read FDetalheErro;
    property Fila: string read FNomeFila;
    property Conectado: Boolean read FConectado;
    property FilaAberta: Boolean read FFilaAberta;
    property TipoAbertura: tpAberturaFila read FTipoAbertura;
    property CabecalhoMsgEstendido: Boolean read FCabecalhoMsgEstendido write FCabecalhoMsgEstendido;
    property InterfaceMQ: tpInterfaceMQ read FInterfaceMQ write FInterfaceMQ;
    property WaitGet: Integer read FGETIntervaloMSeg write FGETIntervaloMSeg;
    property UtilizarSSL: Boolean read FUtilizarSSL write FUtilizarSSL;
    property CipherSpec: AnsiString read FCipherSpec write FCipherSpec;
    property CipherSuite: AnsiString read FCipherSuite write FCipherSuite;
    property KeyResetCount: Integer read FKeyResetCount write FKeyResetCount;
    property SSLCertRevocationCheck: Boolean read FSSLCertRevocationCheck write FSSLCertRevocationCheck;
    property SSLFipsRequired: Boolean read FSSLFipsRequired write FSSLFipsRequired;
    property SSLKeyRepository: AnsiString read FSSLKeyRepository write FSSLKeyRepository;
    property SSLPeerName: AnsiString read FSSLPeerName write FSSLPeerName;
    property Server: String read FServer write FServer;
    property Port: Word read FPort write FPort default 1414;
    property Channel: String read FChannel write FChannel;


constructor TJDMQFila.Create(AOwner: TComponent);
  FServer := 'localhost';
  FPort := 1414;
  FChannel := '';
  FConectado := False;
  FInTransacao := False;
  FFilaAberta := False;
  FQueueManagerOpen := '';
  FQueueManagerRemoto := '';
  FUserIDOpen := '';
  FInterfaceMQ := imMQAX;
  FGETIntervaloMSeg := 0;
  FJDMQ_ConnID := '';
  FUtilizarSSL := False;
  FSSLKeyRepository := 'C:\fontes\mq\certificado\client';
  FCipherSuite := 'SSL_RSA_WITH_AES_256_CBC_SHA256';
  FCipherSpec := 'TLS_RSA_WITH_AES_256_CBC_SHA256';
  FSSLPeerName := '';
  FKeyResetCount := 0;
  FSSLFipsRequired := False;
  FSSLCertRevocationCheck := False;


function TJDMQFila.Desconectar: Boolean;
      Result := False;
      try
 if Assigned(FMQAX_QMgr) and (FMQAX_QMgr.ConnectionStatus) then
              begin
                  FMQAX_QMgr.Backout;
                  FInTransacao := False;
                finally
                  FMQAX_QMgr.Disconnect;
                end;
              end;
      except
      end;

      FQueueManager := '';
      FNomeFila := '';
      FFilaAberta := False;
      FConectado := False;
      Result := True;


function TJDMQFila.inTransaction: Boolean;
  Result := FInTransacao;

function TJDMQFila.Conectar(QMgr: AnsiString): Boolean;
var
  MQI_QMgrName: MQCHAR48;
  MQI_CompCode, MQI_Reason: MQLONG;
  JDMQ_ReasonCode: Integer;
begin
      Result := False;
      try
        FCodigoErro := -12000;
        if Conectado then
          Desconectar;
        FCodigoErro := -14000;

          FCodigoErro := -14100;
          if not Assigned(FMQAX_MQSession) then
          begin
            TJDTraceApp.Log(FJDTraceApp, TempoProc, Proc, 'CoMQSession.Create');
            FMQAX_MQSession := CoMQSession.Create;
          end;

          FCodigoErro := -14110;
          TJDTraceApp.Log(FJDTraceApp, TempoProc, Proc, 'AccessQueueManager(QMgr)');
          FMQAX_QMgr := (FMQAX_MQSession.AccessQueueManager(QMgr) as MQQueueManager);

          FCodigoErro := -14120;
          TJDTraceApp.Log(FJDTraceApp, TempoProc, Proc, 'Connect');
          FMQAX_QMgr.Connect;

        FCodigoErro := -15000;

        FQueueManager := QMgr;
        FConectado := True;
        FFilaAberta := False;

        FCodigoErro := -16000;

        if Assigned(Mensagem) then
          Mensagem.Free;

        case FInterfaceMQ of
          imMQAX:
            Mensagem := TJDMQMensagem.MQAX_Create(FMQAX_MQSession);
          imMQI:
            Mensagem := TJDMQMensagem.MQI_Create();
          imJDDotNet:
            Mensagem := TJDMQMensagem.JDMQ_Create();
        end;

        FCodigoErro := 0;
        FNomeErro := '';

        Result := True;

      except
        on E: Exception do
        begin

          TJDTraceApp.Log(FJDTraceApp, TempoProc, Proc, 'Erro [' + InttoStr(FCodigoErro) + ']: ' + E.Message);

          try

            if FInterfaceMQ = imMQAX then
            begin

              if (FMQAX_MQSession.ReasonCode <> 0) then
              begin
                FCodigoErro := FMQAX_MQSession.ReasonCode;
                FNomeErro := FMQAX_MQSession.ReasonName + '(MQSession)';
              end
              else
              begin
                if (FMQAX_QMgr.ReasonCode <> 0) then
                begin
                  FCodigoErro := FMQAX_QMgr.ReasonCode;
                  FNomeErro := FMQAX_QMgr.ReasonName + '(MQQueueManager)';
                end
                else
                  FNomeErro := 'Problemas com a instala��o(mem�ria) do Client do MQSeries';
              end;
            end;

            if FInterfaceMQ = imMQI then
            begin
              FCodigoErro := MQI_Reason;
              FNomeErro := '';
            end;

            if FInterfaceMQ = imJDDotNet then
            begin
              FCodigoErro := JDMQ_ReasonCode;
              FNomeErro := JDMQ_NomeErro(FCodigoErro);

              if FCodigoErro = 2012 then
              begin

                FDetalheErro :=
                  'A funcionalidade de reconex�o ao MQ somente � suportada a partir do MQ SERIES vers�o 7.1.0. ' +
                  'Verifique qual a vers�o instalada do MQ SERVER e MQ CLIENT est� sendo utilizada pelo servi�o.';
              end;

            end;

          except
            FCodigoErro := -200000 + FCodigoErro;
            FNomeErro := 'Problemas com a instala��o do Client do MQSeries';
          end;

        end;
      end;





function TJDMQFila.RollBack(GeraException: Boolean = True): Boolean;
              FMQAX_QMgr.Backout;
              FInTransacao := False;
        Result := True;

function TJDMQFila.AbrirFila(const NomeFila: AnsiString; OpcoesAbertura: Integer; QueueManager: AnsiString = ''; QueueManagerRemoto: AnsiString = ''; UserID: AnsiString = ''): Boolean;
      Result := AbrirFilaMQ(NomeFila, OpcoesAbertura, QueueManager, QueueManagerRemoto, UserID);

function TJDMQFila.AbrirFila(const NomeFila: AnsiString; TipoAbertura: tpAberturaFila; QueueManager: AnsiString = ''; QueueManagerRemoto: AnsiString = ''; UserID: AnsiString = ''): Boolean;
      OpcsAbertura := 0;
      if (TipoAbertura = afSomentePut) or (TipoAbertura = afPutGet) then
      begin
        OpcsAbertura := OpcsAbertura + MQOO_OUTPUT;
        if FCabecalhoMsgEstendido then
          OpcsAbertura := OpcsAbertura + MQOO_SET_ALL_CONTEXT;
      end;

      if (TipoAbertura = afSomenteGet) or (TipoAbertura = afPutGet) then
        OpcsAbertura := OpcsAbertura + MQOO_INPUT_AS_Q_DEF;
      Result := AbrirFilaMQ(NomeFila, OpcsAbertura, QueueManager, QueueManagerRemoto, UserID);
      FTipoAbertura := TipoAbertura;


function TJDMQFila.AbrirFilaMQ(const NomeFila: AnsiString; OpcoesAbertura: Integer; QueueManager, QueueManagerRemoto, UserID: AnsiString): Boolean;
var
  MQI_CompCode, MQI_Reason: MQLONG;
  JDMQ_ReasonCode: Integer;
  OpcsAbertura: Integer;
      Result := False;
      FFilaAberta := False;
      FCodigoErro := 0;
      FNomeErro := '';
      try
        OpcsAbertura := OpcoesAbertura;
         FMQAX_QueueFila := (FMQAX_QMgr.AccessQueue(NomeFila, OpcsAbertura, QueueManager, QueueManagerRemoto, UserID) as MQQueue);
      except
        // Erro na Abertura
        on E: Exception do
        begin
          FConectado := False;
            FDetalheErro := '[OpcsAbertura:' + InttoStr(OpcsAbertura) + ']' + E.Message;
        end;
      end;

      // Abrindo a Fila
      try
         FMQAX_QueueFila.Open;
        FQueueManagerOpen := QueueManager;
        FQueueManagerRemoto := QueueManagerRemoto;
        FUserIDOpen := UserID;
        FTipoAbertura := TipoAbertura;
        FNomeFila := NomeFila;
        FFilaAberta := True;
        Result := True;
      except
        on e1: Exception do
        begin
          FConectado := False;
            try
              FCodigoErro := FMQAX_QMgr.ReasonCode;
              FNomeErro := FMQAX_QMgr.ReasonName;
            except
              FCodigoErro := -20000;
              FNomeErro := 'MQAX: Erro na abertura da Fila: ' + NomeFila + ' [' + e1.Message + ']';
            end;

        end;
      end;



function TJDMQFila.EnviarMensagem(OpcoesEnvio: Integer; EnvMensagemID: string = ''; EnvMsgRelacionadaID: string = ''): Boolean;
      Result := EnviarMensagemMQ(OpcoesEnvio, EnvMensagemID, EnvMsgRelacionadaID);

function TJDMQFila.EnviarMensagem(EnvMensagemID: string = ''; EnvMsgRelacionadaID: string = ''): Boolean;
      if FCabecalhoMsgEstendido then
        Result := EnviarMensagemMQ(MQPMO_SYNCPOINT + MQPMO_SET_ALL_CONTEXT, EnvMensagemID, EnvMsgRelacionadaID)
      else
        Result := EnviarMensagemMQ(MQPMO_SYNCPOINT, EnvMensagemID, EnvMsgRelacionadaID);


function TJDMQFila.EnviarMensagemMQ(OpcoesEnvio: Integer; EnvMensagemID, EnvMsgRelacionadaID: AnsiString): Boolean;

      Result := False;

      if (Mensagem.Tamanho > 0) then
      begin
        Mensagem.ClearID;
      end
      else
      begin
        FCodigoErro := -10001;
        FNomeErro := 'Mensagem n�o inst�nciada para realizar o Envio';
        Result := False;
        Exit;
      end;

      if EnvMensagemID <> '' then
        Mensagem.MensagemID := EnvMensagemID;

      if EnvMsgRelacionadaID <> '' then
        Mensagem.RelacionadoID := EnvMsgRelacionadaID;

      if FInterfaceMQ = imMQAX then
      begin
        MQPutOpcoes := (FMQAX_MQSession.AccessPutMessageOptions as MQPutMessageOptions);
        MQPutOpcoes.Options := OpcoesEnvio;
      end;

      try
              FMQAX_QueueFila.Put(Mensagem.MQMsg_MQAX, MQPutOpcoes);
              FInTransacao := True;
              EnvMensagemID := Mensagem.MensagemID;
              EnvMsgRelacionadaID := Mensagem.RelacionadoID;
        Result := True;
      except
        Result := False;
         try
                  FCodigoErro := FMQAX_QueueFila.ReasonCode;
                  FNomeErro := FMQAX_QueueFila.ReasonName;
                except
                  FCodigoErro := -10000;
                  FNomeErro := 'Erro no envio da Mensagem';
                end;
      end;


function TJDMQFila.ReceberMensagem(OpcoesRecebimento: Integer; RecMensagemID: string = ''; RecMsgRelacionadaID: string = ''): Boolean;
      Result := ReceberMensagemMQ(OpcoesRecebimento, RecMensagemID, RecMsgRelacionadaID);

function TJDMQFila.ReceberMensagem(RecMensagemID: string = ''; RecMsgRelacionadaID: string = ''): Boolean;
      Result := ReceberMensagemMQ(MQGMO_SYNCPOINT, RecMensagemID, RecMsgRelacionadaID);

function TJDMQFila.ReceberMensagemMQ(OpcoesRecebimento: Integer; RecMensagemID, RecMsgRelacionadaID: AnsiString): Boolean;

      Result := False;

      if not FilaAberta then
        raise Exception.Create('JDMQFila.ReceberMensagem: Fila n�o est� aberta');

      Mensagem.Clear;

      if RecMensagemID <> '' then
        Mensagem.MensagemID := RecMensagemID;

      if RecMsgRelacionadaID <> '' then
        Mensagem.RelacionadoID := RecMsgRelacionadaID;

      MQGetOpcoes := (FMQAX_MQSession.AccessGetMessageOptions as MQGetMessageOptions);
      MQGetOpcoes.Options := OpcoesRecebimento;
      if (FGETIntervaloMSeg > 0) then
      begin
        MQGetOpcoes.Options := MQGetOpcoes.Options + MQGMO_WAIT;
        MQGetOpcoes.WaitInterval := FGETIntervaloMSeg;
      end;

      try
              FMQAX_QueueFila.Get(Mensagem.MQMsg_MQAX, MQGetOpcoes);
              FInTransacao := True; // Controle de Transa��o
              RecMensagemID := Mensagem.MensagemID;
              RecMsgRelacionadaID := Mensagem.RelacionadoID;
        Mensagem.CarregaGet();
        Result := True;
      except
        on e1: Exception do
        begin
          Result := False;
            try
              FCodigoErro := FMQAX_QueueFila.ReasonCode;
              if (FCodigoErro = 2033) then
                FNomeErro := 'N�o existem mensagens na fila ' + FNomeFila
              else
              begin
                FNomeErro := FMQAX_QueueFila.ReasonName;
                FFilaAberta := False;
                Desconectar();
              end;
            except
              FCodigoErro := -10000;
              FNomeErro := 'Erros no Get da Mensagem. [' + e1.Message + ']';
              Desconectar();
            end;
        end;
      end;



function TJDMQFila.FecharFila(): Boolean;
      Result := FecharFilaMQ();

function TJDMQFila.FecharFilaMQ: Boolean;
      Result := False;
      FFilaAberta := False;
      Result := True

function TJDMQFila.Commit(GeraException: Boolean = True): Boolean;
var
  MQI_CompCode, MQI_Reason: MQLONG;
  JDMQ_ReasonCode: Integer;
  MQAX_ReasonCode: Integer;
  MQAX_ReasonName: String;

      try

          try
            FMQAX_QMgr.Commit;
            FInTransacao := False;
          except
              MQAX_ReasonCode := FMQAX_QMgr.ReasonCode;
              MQAX_ReasonName := FMQAX_QMgr.ReasonName;
          end;
          if (not FMQAX_QMgr.IsConnected) then
          begin
            MQAX_ReasonCode := -99995;
            MQAX_ReasonName := 'MQAX: QueueManager Disconnected';
            raise Exception.Create('MQAX: Falha no Commit. QueueManager disconnected');
          end;

        Result := True;
      except
        on E: Exception do
        begin
          Result := False;
                FCodigoErro := MQAX_ReasonCode;
                FNomeErro := MQAX_ReasonName;
      end;



----------------------------------------------------------------------------------------------------------------
----------------------------------------------------------------------------------------------------------------

unit uTJDMQMensagem;

interface

uses MQAX200_TLB, JDInterfaceMQI;

const
  JDMQI_TAM_BUFFER = 4 * 1024 * 1024;
  JDMQAX_TAM_BUFFER = 4 * 1024 * 1024;

type
  tpFormatoMsg = (fm_MQFMT_NONE, fm_MQFMT_STRING);

  tpMQMsgPersistencia = (mpNaoPersistente, mpPersistente, mpDefaultFila);
  tpInterfaceMQ = (imMQAX, imMQI, imJDDotNet);

  TJDMQMensagem = class
  private
    FInterfaceMQ: tpInterfaceMQ;
    FDataEnvio: TDateTime;

    FEnviaCOA_RecebidoCOA: Boolean;
    FEnviaCOD_RecebidoCOD: Boolean;

    FEnviaCOAFullData: Boolean;
    FEnviaCODFullData: Boolean;

    FVersaoMQ: Integer;
    FFormatoMsg: tpFormatoMsg;

    FMQAX_MQMsg: MQMessage;

    FMQI_MsgDesc: MQMD;
    FMQI_Conteudo: AnsiString;

    FJDMQ_Conteudo: AnsiString;
    FJDMQ_MessagaID: AnsiString;
    FJDMQ_MessagaIDRelac: AnsiString;
    FJDMQ_ReplyQM: AnsiString;
    FJDMQ_ReplyFila: AnsiString;
    FJDMQ_Prioridade: Integer;
    FJDMQ_Persistente: tpMQMsgPersistencia;
    FJDMQ_PutDateTime: AnsiString;
  public
    function isCoaCod: Boolean;
    function isCoa: Boolean;
    function isCod: Boolean;
    procedure ClearID;
    procedure Clear;
    procedure ClearConteudo;
    procedure ClearReport;
    procedure CarregaGet;

    property MQI_Conteudo: AnsiString read FMQI_Conteudo write FMQI_Conteudo;
    property JDMQ_PutDateTime: AnsiString read FJDMQ_PutDateTime write FJDMQ_PutDateTime;
    property Conteudo: AnsiString read GetConteudo write SetConteudo;
    property Tamanho: Integer read GetTamanho;
    property MQMsg_MQAX: MQMessage read FMQAX_MQMsg write FMQAX_MQMsg;
    property MensagemID: AnsiString read GetMensagemID write SetMensagemID;
    property RelacionadoID: string read GetRelacionadoID write SetRelacionadoID;
    property DataEnvio: TDateTime read GetDataEnvio;
    property COA: Boolean read FEnviaCOA_RecebidoCOA write SetEnviaCOA_RecebidoCOA;
    property COD: Boolean read FEnviaCOD_RecebidoCOD write SetEnviaCOD_RecebidoCOD;
    property COA_FullData: Boolean read FEnviaCOAFullData write SetEnviaCOAFullData;
    property COD_FullData: Boolean read FEnviaCODFullData write SetEnviaCODFullData;
    property ReplyQM: AnsiString read GetReplyQM write SetReplyQM;
    property ReplyFila: AnsiString read GetReplyFila write SetReplyFila;
    property Formato: tpFormatoMsg read FFormatoMsg write SetFormatoMsg default fm_MQFMT_NONE;
    property FormatoStr: string read GetFormatoMsgStr write SetFormatoMsgStr;
    property ApplicationIdData: AnsiString read GetApplicationIdData write SetApplicationIdData;
    property ContabilidadeHex: AnsiString read GetContabilidadeHex;
    property Contabilidade: AnsiString read GetContabilidade;
    property Expiry_MSeg: Integer read GetExpiryMSeg write SetExpiryMSeg;
    property CharacterSet: Integer read GetCharacterSet write SetCharacterSet;
    property Prioridade: Integer read GetPrioridade write SetPrioridade;
    property Persistente: tpMQMsgPersistencia read GetPersistente write SetPersistente;
    property MsgDesc: MQMD read FMQI_MsgDesc write FMQI_MsgDesc;
    property PutDateTime: AnsiString read GetPutDateTime;
    property PutApplicationName: AnsiString read GetPutApplicationName;
    property Expiry: Integer read GetExpiry;
    property UserId: AnsiString read GetUserId;


function StringtoHex(Texto: AnsiString): AnsiString;
  Hexadecimais = '0123456789ABCDEF';
  result := '';
  for iPosLetra := 1 to Length(Texto) div 2 do
  begin
    sValorHex := Copy(Texto, (iPosLetra * 2) - 1, 2);
    Digito1 := pos(sValorHex[1], Hexadecimais) - 1;
    Digito2 := pos(sValorHex[2], Hexadecimais) - 1;
    result := result + AnsiChar(Digito1 * 16 + Digito2);
  end;


constructor TJDMQMensagem.MQAX_Create(MQSessao: MQSession);
  FEnviaCOA_RecebidoCOA := False;
  FEnviaCOD_RecebidoCOD := False;
  FMQAX_MQMsg := (MQSessao.AccessMessage as MQMessage);
  FFormatoMsg := fm_MQFMT_NONE;
  FMQAX_MQMsg.Format := FormatoMsgtoMQAXFormatStr(FFormatoMsg);
  FMQAX_MQMsg.MessageType := MQMT_DATAGRAM;
  FMQAX_MQMsg.Report := MQRO_NONE;
  FMQAX_MQMsg.ReplyToQueueManagerName := '';
  FMQAX_MQMsg.ReplyToQueueName := '';
  FVersaoMQ := 6;
  Clear;

constructor TJDMQMensagem.JDMQ_Create();
  FInterfaceMQ := imJDDotNet;
  FFormatoMsg := fm_MQFMT_NONE;
  FJDMQ_Prioridade := -1;
  FJDMQ_Persistente := mpDefaultFila;
  Clear;



procedure TJDMQMensagem.CarregaGet;
  FDataEnvio := now();
  FFormatoMsg := MQFormatStrtoFormatoMsg(MQMsg_MQAX.Format);
  try
    FDataEnvio := MQPutDateTimetoDateTime(MQMsg_MQAX.PutDateTime);
  except
  end;

procedure TJDMQMensagem.ClearConteudo;
    FMQAX_MQMsg.ClearMessage;
    FMQAX_MQMsg.ResizeBuffer(JDMQAX_TAM_BUFFER);

procedure TJDMQMensagem.ClearReport;
    FMQAX_MQMsg.Report := MQRO_NONE;

procedure TJDMQMensagem.Clear;
  ClearID;
  ClearReport;
  ClearConteudo;

procedure TJDMQMensagem.ClearID;
    if FVersaoMQ < 6 then
    begin
      FMQAX_MQMsg.MessageId := '';
      FMQAX_MQMsg.MessageIdHex := '';
      FMQAX_MQMsg.CorrelationId := '';
    end;
    FMQAX_MQMsg.MessageId := #0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0;
    FMQAX_MQMsg.CorrelationId := #0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0;
    FMQAX_MQMsg.Feedback := 0;
    FMQAX_MQMsg.ClearErrorCodes;


procedure TJDMQMensagem.SetApplicationIdData(const Value: AnsiString);
      FMQAX_MQMsg.ApplicationIdData := Value;

procedure TJDMQMensagem.SetCharacterSet(const Value: Integer);
      FMQAX_MQMsg.CharacterSet := Value;

procedure TJDMQMensagem.SetConteudo(NewVal: AnsiString);
  ClearID;
  ClearConteudo;
    if Length(NewVal) > 0 then
    begin
        for intLetra := 1 to Length(NewVal) do
        begin
          Letra := Ord(NewVal[intLetra]);
          FMQAX_MQMsg.WriteByte(Letra);
        end;
    end;




procedure TJDMQMensagem.SetEnviaCOA_RecebidoCOA(const Value: Boolean);
  FEnviaCOA_RecebidoCOA := Value;
  AplicaConfigReport;

procedure TJDMQMensagem.AplicaConfigReport;
    FMQAX_MQMsg.Report := MQRO_NONE;
    if FEnviaCOAFullData then
        FMQAX_MQMsg.Report := FMQAX_MQMsg.Report + MQRO_COA_WITH_FULL_DATA
    else
      if FEnviaCOA_RecebidoCOA then
          FMQAX_MQMsg.Report := FMQAX_MQMsg.Report + MQRO_COA;
    if FEnviaCODFullData then
        FMQAX_MQMsg.Report := FMQAX_MQMsg.Report + MQRO_COD_WITH_FULL_DATA
    else
      if FEnviaCOD_RecebidoCOD then
          FMQAX_MQMsg.Report := FMQAX_MQMsg.Report + MQRO_COD;

procedure TJDMQMensagem.SetEnviaCOAFullData(const Value: Boolean);
  FEnviaCOAFullData := Value;
  AplicaConfigReport;

procedure TJDMQMensagem.SetEnviaCOD_RecebidoCOD(const Value: Boolean);
  FEnviaCOD_RecebidoCOD := Value;
  AplicaConfigReport;

procedure TJDMQMensagem.SetEnviaCODFullData(const Value: Boolean);
  FEnviaCODFullData := Value;
  AplicaConfigReport;

procedure TJDMQMensagem.SetExpiryMSeg(const Value: Integer);
  FMQAX_MQMsg.Expiry := (Value div 100);

procedure TJDMQMensagem.SetFormatoMsg(const Value: tpFormatoMsg);
begin
  FFormatoMsg := Value;
      FMQAX_MQMsg.Format := FormatoMsgtoMQAXFormatStr(FFormatoMsg);

procedure TJDMQMensagem.SetFormatoMsgStr(const Value: string);
  FFormatoMsg := MQFormatStrtoFormatoMsg(Value);
FMQAX_MQMsg.Format := FormatoMsgtoMQAXFormatStr(FFormatoMsg);

procedure TJDMQMensagem.SetMensagemID(const Value: AnsiString);
    MsgId := StringtoHex(Value);
    FMQAX_MQMsg.MessageId := Copy(MsgId + StringofChar(#0, 24), 1, 24);

procedure TJDMQMensagem.SetPersistente(const Value: tpMQMsgPersistencia);
    case Value of
      mpNaoPersistente:
        FMQAX_MQMsg.Persistence := 0;
      mpPersistente:
        FMQAX_MQMsg.Persistence := 1;
      mpDefaultFila:
        FMQAX_MQMsg.Persistence := 2;
    end;

procedure TJDMQMensagem.SetPrioridade(const Value: Integer);
      FMQAX_MQMsg.Priority := Value;

procedure TJDMQMensagem.SetRelacionadoID(const Value: string);
    MsgId := StringtoHex(Value);
    FMQAX_MQMsg.CorrelationId := Copy(MsgId + StringofChar(#0, 24), 1, 24);

procedure TJDMQMensagem.SetReplyFila(const Value: AnsiString);
FMQAX_MQMsg.ReplyToQueueName := Value;

procedure TJDMQMensagem.SetReplyQM(const Value: AnsiString);
FMQAX_MQMsg.ReplyToQueueManagerName := Value;


function TJDMQMensagem.isCoaCod: Boolean;
result := (MQMsg_MQAX.Feedback = MQFB_COA) or (MQMsg_MQAX.Feedback = MQFB_COD);


function TJDMQMensagem.isCoa: Boolean;
      result := (MQMsg_MQAX.Feedback = MQFB_COA);

function TJDMQMensagem.isCod: Boolean;
      result := (MQMsg_MQAX.Feedback = MQFB_COD);


function TJDMQMensagem.MQPutDateTimetoDateTime(PutDateTime: String): TDateTime;
var
  DHExtraida: TDateTime;
  sDHExtraida: String;
begin
  DHExtraida := StrToDateTimeDef(PutDateTime, 0);

  if DHExtraida = 0 then
  begin
    if (Length(PutDateTime) = 19)
      and (PutDateTime[3] = '/')
      and (PutDateTime[6] = '/')
      and (PutDateTime[11] = ' ')
      and (PutDateTime[14] = ':')
      and (PutDateTime[17] = ':')
    then
        DHExtraida :=
        EncodeDate(StrtoInt(Copy(PutDateTime, 7, 4)), StrtoInt(Copy(PutDateTime, 4, 2)), StrtoInt(Copy(PutDateTime, 1, 2))) +
        EncodeTime(StrtoInt(Copy(PutDateTime, 12, 2)), StrtoInt(Copy(PutDateTime, 15, 2)), StrtoInt(Copy(PutDateTime, 18, 2)), 0)
    else
        DHExtraida := now();
  end;

  result := TJDData.UTCtoDateTime(DHExtraida);
end;

function TJDMQMensagem.MQFormatStrtoFormatoMsg(Formato: string): tpFormatoMsg;
begin
  Formato := Trim(Formato);

  result := fm_MQFMT_NONE;

  if SameText(Formato, 'MQSTR') then
      result := fm_MQFMT_STRING;
end;

function TJDMQMensagem.FormatoMsgtoFormatMsgStr(Formato: tpFormatoMsg = fm_MQFMT_NONE): AnsiString;
begin
  case Formato of
    fm_MQFMT_STRING: result := 'MQSTR';
  else
    result := '';
  end;
end;

function InterfaceMQtoStr(InterfaceMQ: tpInterfaceMQ) : String;
begin
  case InterfaceMQ of
    imMQAX: Result := 'MQAX';
    imMQI: Result := 'MQI';
    imJDDotNet: Result := 'JDMQ.Net';
  end;
end;

*/
