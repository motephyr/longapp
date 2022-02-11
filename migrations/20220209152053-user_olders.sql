
-- +migrate Up

CREATE TABLE public."user_olders" (
	id serial4 NOT NULL,
	user_id int4 NULL,
	older_id int4 NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	CONSTRAINT user_olders_pkey PRIMARY KEY (id),
	CONSTRAINT user_olders_user_id_foreign FOREIGN KEY (user_id) REFERENCES public.users(id),
	CONSTRAINT user_olders_older_id_foreign FOREIGN KEY (older_id) REFERENCES public.olders(id)

);
CREATE INDEX user_olders_user_id_index ON public.user_olders USING btree (user_id);
CREATE INDEX user_olders_older_id_index ON public.user_olders USING btree (older_id);

-- +migrate Down

DROP TABLE public."groups";
