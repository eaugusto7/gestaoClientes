create table clientes(
    id serial primary key,
    nome varchar,
    cpf varchar,
    rg varchar,
    email varchar,
    telefone varchar,
    celular varchar,
    dataNascimento date,
    sexo varchar,
    profissao varchar
);