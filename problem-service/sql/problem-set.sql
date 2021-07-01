
create table problem_set (id CHAR(36) primary key, title text not null default 'New Problem Set', access_level text not null default '', included_problems text not null default '', author text not null default '');
