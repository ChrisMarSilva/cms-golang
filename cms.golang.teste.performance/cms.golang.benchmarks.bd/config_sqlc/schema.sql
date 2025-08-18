CREATE TABLE if not exists "TbPerson" (
  id uuid NOT NULL PRIMARY KEY,
  name varchar(200) NOT NULL,
  created_at TIMESTAMP NOT NULL
);