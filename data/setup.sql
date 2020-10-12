-- Table: public.notes
 -- Database: notes

-- DROP DATABASE notes;

CREATE DATABASE notes
    WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'Polish_Poland.1250'
    LC_CTYPE = 'Polish_Poland.1250'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;
 
 DROP TABLE public.notes;
CREATE table public.notes
(
    id serial primary key,
    uuid varchar(64) not null,
    title text,
    content text,
	version integer,
    created timestamp not null,
    modified timestamp not null
)

TABLESPACE pg_default;
