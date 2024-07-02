-- Write your migrate up statements here
ALTER TABLE public.application DROP COLUMN phone;
ALTER TABLE public.application DROP COLUMN source_of_information;
---- create above / drop below ----
ALTER TABLE public.application ADD COLUMN phone TEXT CHECK (char_length(phone) <= 20);
ALTER TABLE public.application ADD COLUMN source_of_information TEXT;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
