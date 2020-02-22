CREATE TABLE public.news (
	id int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
	author varchar(64) NOT NULL,
	body text NOT NULL,
	created timestamp NOT NULL DEFAULT '2000-01-01 00:00:00'::timestamp without time zone,
	CONSTRAINT news_pk PRIMARY KEY (id)
);