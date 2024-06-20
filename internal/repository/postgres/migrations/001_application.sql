-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS application (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    tg_user_id BIGINT,
    full_name TEXT,
    phone TEXT CHECK (char_length(phone) <= 20),
    district TEXT CHECK (char_length(district) <= 100),
    purpose_of_application TEXT,
    source_of_information TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_application_tg_user_id ON application(tg_user_id);

CREATE OR REPLACE FUNCTION trigger_set_timestamp() 
RETURNS TRIGGER AS $$ 
BEGIN 
  NEW.updated_at = NOW(); 
  RETURN NEW; 
  END; 
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER set_timestamp
BEFORE UPDATE ON public.application 
FOR EACH ROW 
EXECUTE PROCEDURE trigger_set_timestamp();
---- create above / drop below ----
DROP TRIGGER IF EXISTS set_timestamp ON application;
DROP FUNCTION IF EXISTS trigger_set_timestamp();
DROP INDEX IF EXISTS idx_application_tg_user_id;
DROP TABLE IF EXISTS application;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
