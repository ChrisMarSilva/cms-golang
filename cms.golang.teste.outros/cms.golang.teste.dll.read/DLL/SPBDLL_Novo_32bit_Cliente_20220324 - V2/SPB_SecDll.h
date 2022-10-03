#ifndef SPB_SECDLL_H
#define SPB_SECDLL_H

/////////////////////////////////////////////////////////////////////////////////////////////////////
/** Versões do padrão do SPB */
/** VERSION_1 ou VERSION_CONF: A versão utilizada será definida no registro de configuração.

	\MACHINE\SOFTWARE\Robo\SPB\PROTOCOL_VERSION
	\MACHINE\SOFTWARE\WOW6432Node\Robo\SPB\PROTOCOL_VERSION

	ou em máquinas com registro virtual:

	Computador\HKEY_USERS\S-1-5-21-3146314955-2074416627-2753764838-1217_Classes\VirtualStore\MACHINE\SOFTWARE\WOW6432Node\Robo\SPB\PROTOCOL_VERSION
*/
#define	H_VERSION_CONF			1
/** VERSION_2: Versão 2 com 3DES.*/
#define	H_VERSION_2				2
/** VERSION_3: Versão 3 com AES256-GCM.*/
#define H_VERSION_3				3
/////////////////////////////////////////////////////////////////////////////////////////////////////

// Tipo: 0
#define MSG_NORMAL				0
// Tipo: 2
#define MSG_TYPE_02				2
// Tipo: 3
#define MSG_GEN0004				4
// Tipo: 1
#define MSG_GEN0006				6
// Tipo: 1
#define MSG_GEN0007				7
// Tipo: 1
#define MSG_GEN0008				8
/** Arquivo assinado (tipo 4). */
#define ARQ_CIFRADO				104
/** Arquivo cifrado e assinado (tipo 6). */
#define ARQ_ASSINADO			106
/** Arquivo ZIP cifrado e assinado (tipo 8). */
#define ZIP_CIFRADO				108
/** Arquivo ZIP cifrado e assinado (tipo 10). */
#define ZIP_ASSINADO			110

#ifdef SPB_SECDLL_EXPORTS
	#define DECLSPEC	__declspec( dllexport )
#else
	#define DECLSPEC	__declspec( dllimport )
#endif

extern "C" DECLSPEC int _stdcall InitializeConnSrv(int *servers_ok);

extern "C" DECLSPEC int _stdcall ReConnect(INT *servers_ok);

extern "C" DECLSPEC int _stdcall InitializeConn();

extern "C" DECLSPEC int _stdcall TerminateAllConn();

extern "C" DECLSPEC int _stdcall TerminateConn(int server_id);

extern "C" __declspec(deprecated("Utilizr a versão EncryptMsgV3")) DECLSPEC int _stdcall EncryptMsg(
#ifdef COM_PARAM_DOMINIO
												 LPSTR lpstrDomainName,	// (in) Domain Name
#endif
												 LPSTR lpstrSISPB,		// (in) Source ISPB
												 LPSTR lpstrDISPB,		// (in) Destination ISPB
												 INT   intCodMsg,		// (in) MessageCode
												 LPSTR lpstrPlainText,	// (in) XML Message
												 VARIANT *pvCipherText,	// (out)SPB Envelop
												 VARIANT *pvLog,		// (out)Symmetric Key, Digital Sign, XML (Unicode 16)
												 INT   intCodError);	// (in) ErrorCode (GEN0004)

extern "C" __declspec(deprecated("Utilizr a versão DecryptMsgV3")) DECLSPEC int _stdcall DecryptMsg(
#ifdef COM_PARAM_DOMINIO
												 LPSTR lpstrDomainName,	// (in) Domain Name
#endif
												 LPSTR lpstrSISPB,		// (in) Source ISPB
												 LPSTR lpstrDISPB,		// (in) Destination ISPB
												 VARIANT vCipherText,	// (in) SPB Envelop 
												 VARIANT *pvReturn,		// (out)XML Message
												 VARIANT *pvLog,		// (out)Symmetric Key, Digital Sign, SPB Envelop + XML Uni 16
												 INT *RemoteError);		// (out)RemoteError if GEN0004

extern "C" __declspec(deprecated("Utilizr a versão EncryptMsgV3")) DECLSPEC int _stdcall EncryptMsgB(
												LPCSTR lpstrDomainName, 	// (in) Domain Name
												LPSTR lpstrSISPB,			// (in) Source ISPB
												LPSTR lpstrDISPB,			// (in) Destination ISPB
												INT   intCodMsg,			// (in) MessageCode
												LPSTR lpstrPlainText,		// (in) XML Message
												INT lpstrPlainTextLen,		// (in) XML Message Length
												BYTE **lpCipherText,		// (out)SPB Envelop (Liberar a memória com free())
												INT *intCipherTextLen,		// (out) SPB Envelop Length
												BYTE **lpLog,				// (out)Symmetric Key, Digital Sign, XML (Unicode 16)  (Liberar a memória com FreeMemory())
												INT *intLogLen,				// (out)SPB Log Length
												INT   intCodError);			// (in) ErrorCode (GEN0004)

extern "C" __declspec(deprecated("Utilizr a versão DecryptMsgV3")) DECLSPEC int _stdcall DecryptMsgB(
												LPCSTR lpstrDomainName,	// (in) Domain Name
												LPSTR lpstrSISPB,		// (in) Source ISPB
												LPSTR lpstrDISPB,		// (in) Destination ISPB
												BYTE *lpCipherText,		// (in) SPB Envelop 
												INT intCipherTextLen,	// (in) SPB Envelop Length
												BYTE **lpReturn,		// (out)XML Message  (Liberar a memória com free())
												INT *intReturnLen,		// (out)XML Message Length
												BYTE **lpLog,			// (out)Symmetric Key, Digital Sign, SPB Envelop + XML Uni 16  (Liberar a memória com FreeMemory())
												INT *intLogLen,			// (out)SPB Log Lenth
												INT *RemoteError);		// (out)RemoteError if GEN0004

extern "C" __declspec(deprecated("Utilizr a versão EncryptMsgV3")) DECLSPEC int _stdcall EncryptMsgC(
												LPCSTR lpstrDomainName, 	// (in) Domain Name
												LPSTR lpstrSISPB,			// (in) Source ISPB
												LPSTR lpstrDISPB,			// (in) Destination ISPB
												INT   intCodMsg,			// (in) MessageCode
												LPSTR lpstrPlainText,		// (in) XML Message
												INT lpstrPlainTextLen,		// (in) XML Message Length
												BYTE **lpCipherText,		// (out)SPB Envelop (Liberar a memória com free())
												INT *intCipherTextLen,		// (out) SPB Envelop Length
												BYTE **lpLog,				// (out)Symmetric Key, Digital Sign, XML (Unicode 16)  (Liberar a memória com FreeMemory())
												INT *intLogLen,				// (out)SPB Log Length
												INT   intCodError);			// (in) ErrorCode (GEN0004)

extern "C" __declspec(deprecated("Utilizr a versão DecryptMsgV3")) DECLSPEC int _stdcall DecryptMsgC(
												LPCSTR lpstrDomainName,	// (in) Domain Name
												LPSTR lpstrSISPB,		// (in) Source ISPB
												LPSTR lpstrDISPB,		// (in) Destination ISPB
												BYTE *lpCipherText,		// (in) SPB Envelop 
												INT intCipherTextLen,	// (in) SPB Envelop Length
												BYTE **lpReturn,		// (out)XML Message  (Liberar a memória com free())
												INT *intReturnLen,		// (out)XML Message Length
												BYTE **lpLog,			// (out)Symmetric Key, Digital Sign, SPB Envelop + XML Uni 16  (Liberar a memória com FreeMemory())
												INT *intLogLen,			// (out)SPB Log Lenth
												INT *RemoteError);		// (out)RemoteError if GEN0004

/** Cifra uma mensagem conforme a versão do padrão.

	@param dwProtocolVersion	(IN) Versão do padrão (H_VERSION_CONF | H_VERSION_2 | H_VERSION_3).
								- H_VERSION_CONF: Seleciona a versão conforme a configuração do serviço.
								- H_VERSION_2:	Utiliza a versão 2 (3DES)
								- H_VERSION_3:	Utiliza a versão 3 (AES256).
	@param lpstrDomainName		(IN) Nome do dominio configurado para o certificado.
	@param lpstrSISPB			(IN) ISPB de origem.
	@param lpstrDISPB			(IN) ISPB de destino.
	@param intCodMsg			(IN) Código de mensagem especial.
	@param lpstrPlainText		(IN) Message para envio.
	@param lpstrPlainTextLen	(IN) Tamanho da message.
	@param lpCipherText			(OUT) Mensagem cifrada (Envelope SPB - Liberar a memória com FreeMemory()).
	@param intCipherTextLen		(OUT) Tamanho do envelope SPB.
	@param lpLog				(OUT) Log de auditoria do BACEN (Liberar a memória com FreeMemory()).
	@param intLogLen			(OUT) Tamanho do envelope SPB Log.
	@param intCodError			(IN) Código de erro para GEN0004.
*/
extern "C" DECLSPEC int _stdcall EncryptMsgV3(
												DWORD dwProtocolVersion,	// (in) Versão do protocolo (H_VERSION_CONF | H_VERSION_2 | H_VERSION_3)
												LPCSTR lpstrDomainName, 	// (in) Domain Name
												LPSTR lpstrSISPB,			// (in) Source ISPB
												LPSTR lpstrDISPB,			// (in) Destination ISPB
												INT   intCodMsg,			// (in) MessageCode
												BYTE *lpstrPlainText,		// (in) XML Message
												INT lpstrPlainTextLen,		// (in) XML Message Length
												BYTE **lpCipherText,		// (out)SPB Envelop (Liberar a memória com free())
												INT *intCipherTextLen,		// (out) SPB Envelop Length
												BYTE **lpLog,				// (out)Symmetric Key, Digital Sign, XML (Unicode 16)  (Liberar a memória com FreeMemory())
												INT *intLogLen,				// (out)SPB Log Length
												INT   intCodError);			// (in) ErrorCode (GEN0004)
/** Decifra uma mensagem conforme a versão do padrão.
	@param dwProtocolVersion	(IN) Versão do padrão (H_VERSION_CONF | H_VERSION_2 | H_VERSION_3).
								O sistema definira a versão conforme a mensagem, então utilize H_VERSION_3.
	@param lpstrDomainName		(IN) Nome do dominio configurado para o certificado.
	@param lpstrSISPB			(IN) ISPB origem.
	@param lpstrDISPB			(IN) ISPB destino.
	@param lpCipherText			(IN) Envelope SPB para ser decifrado.
	@param intCipherTextLen		(IN) Tamanho do envelope SPB.
	@param lpReturn				(OUT) Messagem decifrada (Liberar a memória com FreeMemory()).
	@param intReturnLen			(OUT) Tamanho da mensagem decifrada.
	@param lpLog				(OUT) Log de auditoria do BACEN (Liberar a memória com FreeMemory()).
	@param intLogLen,			(OUT) Tamanho do log de auditoria do BACEN.
	@param RemoteError			(OUT) Código de error de GEN0004 (A aplicação deve utilizar o código de erro retornado na mensagem GEN0004).
*/
extern "C" DECLSPEC int _stdcall DecryptMsgV3(
												DWORD dwProtocolVersion,// (in) Versão do protocolo (H_VERSION_CONF | H_VERSION_2 | H_VERSION_3)
												LPCSTR lpstrDomainName,	// (in) Domain Name
												LPSTR lpstrSISPB,		// (in) Source ISPB
												LPSTR lpstrDISPB,		// (in) Destination ISPB
												BYTE *lpCipherText,		// (in) SPB Envelop 
												INT intCipherTextLen,	// (in) SPB Envelop Length
												BYTE **lpReturn,		// (out)XML Message  (Liberar a memória com free())
												INT *intReturnLen,		// (out)XML Message Length
												BYTE **lpLog,			// (out)Symmetric Key, Digital Sign, SPB Envelop + XML Uni 16  (Liberar a memória com FreeMemory())
												INT *intLogLen,			// (out)SPB Log Lenth
												INT *RemoteError);		// (out)RemoteError if GEN0004

extern "C" DECLSPEC int _stdcall StatusConn(
												INT *servers_ok
											   ,LPSTR ip_servers_ok
											   ,INT *ip_servers_ok_len
											   ,INT *fault_servers
											   ,LPSTR ip_fault_servers
											   ,INT *ip_fault_servers_len);

extern "C" DECLSPEC int _stdcall GetServerPoolSize();

extern "C" DECLSPEC int _stdcall GetConnectedServers();

extern "C" DECLSPEC int _stdcall VariantToChar(
												VARIANT var
											    ,INT begin
												,INT end
												,LPSTR msg
												,INT *msgLen);

extern "C" DECLSPEC int _stdcall FreeMemory(BYTE *pbMalloc);

#endif