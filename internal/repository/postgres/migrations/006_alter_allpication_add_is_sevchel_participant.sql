-- Write your migrate up statements here
ALTER TABLE public.application ADD COLUMN is_sev_chel_participant BOOLEAN DEFAULT FALSE;
---- create above / drop below ----
ALTER TABLE public.application DROP COLUMN is_sev_chel_participant;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
