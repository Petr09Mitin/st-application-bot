-- Write your migrate up statements here
ALTER TABLE public.application ADD COLUMN social_media_link TEXT DEFAULT ''::TEXT;
---- create above / drop below ----
ALTER TABLE public.application DROP COLUMN social_media_link;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
