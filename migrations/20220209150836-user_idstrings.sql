
-- +migrate Up

CREATE TABLE public."user_idstrings" (
	id serial4 NOT NULL,
	user_id int4 NULL,
	idstring varchar(255) NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	CONSTRAINT user_idstrings_pkey PRIMARY KEY (id),
	CONSTRAINT user_idstrings_user_id_foreign FOREIGN KEY (user_id) REFERENCES public.users(id)

);
CREATE INDEX user_idstrings_user_id_index ON public.user_idstrings USING btree (user_id);
CREATE INDEX user_idstrings_idstring_index ON public.user_idstrings USING btree (idstring);

-- +migrate Down

DROP TABLE public."groups";
