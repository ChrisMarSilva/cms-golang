
DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS users (
	id TEXT NOT NULL PRIMARY KEY,
    nome TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    is_active INTEGER NOT NULL DEFAULT 1,
    created_at DATETIME NOT NULL DEFAULT(STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW', 'localtime'))
    
);

-- created_at2 TIMESTAMP NOT NULL DEFAULT(STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW'))
-- CREATE UNIQUE INDEX IF NOT EXISTS idx_users_email ON users (email);

DELETE FROM users;
INSERT INTO users(id, nome, email, password, is_active) VALUES ('39b190b2-061b-47ac-90fc-f59997866a92', 'ChrisMarSil', 'chris.mar.silva@gmail.com', '123', 1);
INSERT INTO users(id, nome, email, password, is_active) VALUES ('119b16ba-2c15-416c-a733-57645643e0ff', 'Pessoa01', 'pessoal.01@gmail.com', '123', 1);
