

USE CMS_TESTE_JDSPB;


SELECT top 100 ID, NOME, STATUS FROM TBUSUARIO WITH(NOLOCK);
-- delete from TBUSUARIO where id = '1774FED7-B68B-494F-A9C4-CB81F2319AD3';
-- delete from TBUSUARIO where id = '481CB4FF-089E-4B0C-A9B5-2FA59FBE6C0E';
SELECT ID, NOME, STATUS FROM TBUSUARIO WITH(NOLOCK) WHERE NOME LIKE 'NOVA PESSOA%';
SELECT ID, NOME, STATUS FROM TBUSUARIO WITH(NOLOCK) WHERE NOME LIKE '%Teste%';


SELECT COUNT(1) FROM TBUSUARIO WITH(NOLOCK);



SELECT conn.session_id, host_name, program_name,    nt_domain, login_name, connect_time, last_request_end_time FROM sys.dm_exec_sessions AS sess JOIN sys.dm_exec_connections AS conn    ON sess.session_id = conn.session_id;
--SELECT * FROM sys.dm_exec_sessions WHERE status = 'running';
SELECT     DB_NAME(dbid) as DBName,     COUNT(dbid) as NumberOfConnections,     loginame as LoginName, hostname, hostprocess FROM    sys.sysprocesses WHERE      dbid > 0 GROUP BY      dbid, loginame, hostname, hostprocess;



SELECT * FROM TBMONITOR WITH(NOLOCK) WHERE DATA_ERRO = CONVERT(VARCHAR(10),GETDATE(),112) ORDER BY COD_ERRO, DATA_ERRO, HORA_ERRO DESC;
-- DELETE FROM TBMONITOR;


---------------------------------------------------------------------------------
---------------------------------------------------------------------------------

/*

DROP TABLE TBUSUARIO;
CREATE TABLE TBUSUARIO( ID UNIQUEIDENTIFIER NOT NULL  PRIMARY KEY, NOME VARCHAR(100) NOT NULL , STATUS CHAR (1) NOT NULL );


DELETE FROM TBUSUARIO;
INSERT INTO TBUSUARIO (ID, NOME, STATUS) VALUES (NEWID(), 'PESSOA 1', 'A');
INSERT INTO TBUSUARIO (ID, NOME, STATUS) VALUES (NEWID(), 'PESSOA 2', 'A');
INSERT INTO TBUSUARIO (ID, NOME, STATUS) VALUES (NEWID(), 'PESSOA 3', 'A');



*/