
create table if not exists "TbPerson"(
  id uuid not null primary key,
  name varchar(200) not null,
  created_at TIMESTAMP WITHOUT TIME  not null
);

create index if not exists idx_person_created_at on "TbPerson" (created_at);

insert into "TbPerson" (id, name, created_at) values
  (gen_random_uuid(), 'João da Silva', CURRENT_TIMESTAMP),
  (gen_random_uuid(), 'Maria Oliveira', CURRENT_TIMESTAMP),
  (gen_random_uuid(), 'Pedro Santos', CURRENT_TIMESTAMP);
