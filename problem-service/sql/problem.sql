
create table problem (id CHAR(36) primary key, version INTEGER not null default 1, statement text not null default '', type text not null default '', answers text not null default '');
