-- Write your migrate up statements here
ALTER TABLE public.application ADD COLUMN resolved BOOLEAN DEFAULT FALSE;
ALTER TABLE public.application ADD COLUMN approved BOOLEAN DEFAULT FALSE;
ALTER TABLE public.application ADD CONSTRAINT application_resolved_approved_check CHECK (resolved = TRUE OR approved = FALSE); -- !resolved -> !approved, implication
ALTER TABLE public.application ADD CONSTRAINT application_tg_user_id_unique UNIQUE (tg_user_id);
---- create above / drop below ----
ALTER TABLE public.application DROP CONSTRAINT application_tg_user_id_unique;
ALTER TABLE public.application DROP CONSTRAINT application_resolved_approved_check;
ALTER TABLE public.application DROP COLUMN resolved;
ALTER TABLE public.application DROP COLUMN approved;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
