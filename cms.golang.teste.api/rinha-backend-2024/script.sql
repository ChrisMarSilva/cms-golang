
CREATE TABLE cliente (
  id SERIAL PRIMARY KEY,
  nome VARCHAR(50) NOT NULL,
  limite INTEGER NOT NULL,
  saldo INTEGER NOT NULL DEFAULT 0
);

CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_cliente ON cliente (id) include (saldo);

CREATE TABLE cliente_transacao (
  id SERIAL PRIMARY KEY,
  cliente_id INTEGER NOT NULL,
  valor INTEGER NOT NULL,
  tipo CHAR(1) NOT NULL,
  descricao VARCHAR(10) NOT NULL,
  dthrregistro TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_transacao_to_cliente FOREIGN KEY(cliente_id) REFERENCES cliente(id)
);

CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_transacoes_id_desc ON cliente_transacao (id desc);

CREATE TABLE cliente_saldo  (
	id SERIAL PRIMARY KEY,
	cliente_id INTEGER NOT NULL,
	total INTEGER NOT NULL,
  CONSTRAINT fk_transacao_to_saldo FOREIGN KEY(cliente_id) REFERENCES cliente(id)
);

CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_saldo_cliente ON cliente_saldo (cliente_id) include (total);

--SET TRANSACTION ISOLATION LEVEL READ UNCOMMITTED;
--SET TRANSACTION ISOLATION LEVEL READ_COMMITED;
--SET GLOBAL TRANSACTION ISOLATION LEVEL READ UNCOMMITTED ;
--SET default_transaction_isolation TO 'repeatable read'

DO $$
BEGIN
	INSERT INTO cliente (nome, limite)
	VALUES ('cliente 01',   1000 * 100), ('cliente 02',    800 * 100), ('cliente 03',  10000 * 100), ('cliente 04', 100000 * 100), ('cliente 05',   5000 * 100);
	INSERT INTO cliente_saldo (cliente_id, total) SELECT id, 0 FROM cliente;
END;
$$;


