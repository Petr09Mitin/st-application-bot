-- Write your migrate up statements here
ALTER TABLE public.application ADD COLUMN tg_username TEXT DEFAULT ''::TEXT;
---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
