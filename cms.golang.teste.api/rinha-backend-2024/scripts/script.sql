
CREATE TABLE cliente (
  id SERIAL PRIMARY KEY,
  nome VARCHAR(50) NOT NULL,
  limite INTEGER NOT NULL,
  saldo INTEGER NOT NULL DEFAULT 0
);

CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_cliente ON cliente(id) INCLUDE (limite, saldo);

CREATE TABLE cliente_transacao (
  id SERIAL NOT NULL, 
  cliente_id INTEGER NOT NULL, 
  valor INTEGER NOT NULL,
  -- saldo INTEGER NOT NULL,
  tipo CHAR(1) NOT NULL CHECK (tipo IN ('c', 'd')),
  descricao VARCHAR(10) NOT NULL,
  dthrregistro TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT pk_cliente_transacao PRIMARY KEY(id, cliente_id) ,
  CONSTRAINT fk_transacao_to_cliente FOREIGN KEY(cliente_id) REFERENCES cliente(id) ON DELETE CASCADE
) PARTITION by LIST(cliente_id); 

CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_transacoes_cliente_id ON cliente_transacao (cliente_id);

CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_transacoes_id_desc ON cliente_transacao (cliente_id, id DESC) INCLUDE (valor, tipo, descricao, dthrregistro);

CREATE TABLE cliente_transacao_1 PARTITION OF cliente_transacao FOR VALUES IN (1);
CREATE TABLE cliente_transacao_2 PARTITION OF cliente_transacao FOR VALUES IN (2);
CREATE TABLE cliente_transacao_3 PARTITION OF cliente_transacao FOR VALUES IN (3);
CREATE TABLE cliente_transacao_4 PARTITION OF cliente_transacao FOR VALUES IN (4);
CREATE TABLE cliente_transacao_5 PARTITION OF cliente_transacao FOR VALUES IN (5);

--ALTER TABLE cliente_transacao NO INHERIT cliente_transacao_1, cliente_transacao_2, cliente_transacao_3, cliente_transacao_4, cliente_transacao_5;

-- CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_transacoes_cliente_id_part1 ON cliente_transacao_1 (cliente_id);
-- CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_transacoes_cliente_id_part2 ON cliente_transacao_2 (cliente_id);
-- CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_transacoes_cliente_id_part3 ON cliente_transacao_3 (cliente_id);
-- CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_transacoes_cliente_id_part4 ON cliente_transacao_4 (cliente_id);
-- CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_transacoes_cliente_id_part5 ON cliente_transacao_5 (cliente_id);

-- CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_transacoes_id_desc_part1 ON cliente_transacao_1 (cliente_id, id DESC) INCLUDE (valor, tipo, descricao, dthrregistro);
-- CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_transacoes_id_desc_part2 ON cliente_transacao_2 (cliente_id, id DESC) INCLUDE (valor, tipo, descricao, dthrregistro);
-- CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_transacoes_id_desc_part3 ON cliente_transacao_3 (cliente_id, id DESC) INCLUDE (valor, tipo, descricao, dthrregistro);
-- CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_transacoes_id_desc_part4 ON cliente_transacao_4 (cliente_id, id DESC) INCLUDE (valor, tipo, descricao, dthrregistro);
-- CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_transacoes_id_desc_part5 ON cliente_transacao_5 (cliente_id, id DESC) INCLUDE (valor, tipo, descricao, dthrregistro);

--ALTER TABLE cliente_transacao INHERIT cliente_transacao_1, cliente_transacao_2, cliente_transacao_3, cliente_transacao_4, cliente_transacao_5;

--PREPARE consulta_cliente_por_id (INTEGER) AS SELECT id, limite, saldo  FROM cliente  WHERE id = $1;
--PREPARE consulta_transacoes_por_cliente_id (INTEGER) AS SELECT valor, tipo, descricao, dthrregistro FROM cliente_transacao WHERE cliente_id = $1 ORDER BY id DESC LIMIT 10;
--PREPARE update_saldo_cliente_por_id (INTEGER, INTEGER) AS UPDATE cliente SET saldo = $2 WHERE id = $1;
--PREPARE insert_transacao_por_cliente (INTEGER, INTEGER, CHAR, VARCHAR, TIMESTAMP) AS INSERT INTO cliente_transacao (cliente_id, valor, tipo, descricao, dthrregistro) VALUES ($1, $2, $3, $4, $5);

--SET TRANSACTION ISOLATION LEVEL READ UNCOMMITTED;
--SET TRANSACTION ISOLATION LEVEL READ_COMMITED;
--SET GLOBAL TRANSACTION ISOLATION LEVEL READ UNCOMMITTED ;
--SET default_transaction_isolation TO 'repeatable read'

DO $$
BEGIN
	INSERT INTO cliente (nome, limite)
	VALUES ('cliente 01',   1000 * 100), 
         ('cliente 02',    800 * 100), 
         ('cliente 03',  10000 * 100), 
         ('cliente 04', 100000 * 100), 
         ('cliente 05',   5000 * 100);
END;
$$;
