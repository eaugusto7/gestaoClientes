create table clientes(
    id serial primary key,
    nome varchar,
    cpf varchar,
    rg varchar,
    email varchar,
    celular2 varchar,
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
    idatendente int,
    status varchar,
    statusfixo boolean,
    formapagamento varchar
);

create table atendentes(
    id serial primary key,
    nome varchar,
    celular varchar,
    idquadrohorario int
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

create table logins(
    id serial primary key,
    username varchar,
    password varchar
);

create table quadroshorarios(
    id serial primary key,
    domingo int[],
    segunda int[],
    terca int[],
    quarta int[],
    quinta int[],
    sexta int[],
    sabado int[]
);
