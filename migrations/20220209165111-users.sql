
-- +migrate Up
ALTER TABLE public.users DROP COLUMN "name";
ALTER TABLE public.users DROP CONSTRAINT users_email_unique;
ALTER TABLE public.users DROP COLUMN email;
ALTER TABLE public.users DROP COLUMN image_url;

-- +migrate Down
