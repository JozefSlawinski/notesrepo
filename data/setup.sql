-- Table: public.notes
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

ALTER TABLE public.notes
    OWNER to postgres;