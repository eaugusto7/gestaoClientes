create table clientes(
    id serial primary key,
    nome varchar,
    cpf varchar,
    rg varchar,
    email varchar,
    telefone varchar,
    celular varchar,
    dataNascimento varchar,
    sexo varchar,
    profissao varchar
);

create table servicos(
    id serial primary key,
    nome varchar,
    valor numeric,
    tempo numeric
);

create table atendimentos(
    id serial primary key,
    nome varchar,
    horario numeric,
    idservico int,
    idatendente int
);

create table atendentes(
    id serial primary key,
    nome varchar,
    telefone varchar
);

create table produtos(
    id serial primary key,
    nome varchar,
    quantidade int,
    descricao varchar,
    fabricante varchar,
    valorcusto numeric,
    valorvenda numeric
);