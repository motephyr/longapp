
-- +migrate Up
drop index user_olders_user_id_index;
drop index user_olders_older_id_index;
CREATE UNIQUE INDEX user_olders_user_id_idx ON user_olders (user_id, older_id);

drop index user_idstrings_user_id_index;
drop index user_idstrings_idstring_index;
CREATE UNIQUE INDEX user_idstrings_user_id_idx ON user_idstrings (user_id, idstring);



-- +migrate Down
CREATE INDEX user_idstrings_user_id_index ON public.user_idstrings USING btree (user_id);
CREATE INDEX user_idstrings_idstring_index ON public.user_idstrings USING btree (idstring);

CREATE INDEX user_olders_user_id_index ON public.user_olders USING btree (user_id);
CREATE INDEX user_olders_older_id_index ON public.user_olders USING btree (older_id);

drop UNIQUE index user_olders_user_id_idx;
drop UNIQUE index user_idstrings_user_id_idx;