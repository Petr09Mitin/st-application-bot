-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS target_chat (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    tg_chat_id BIGINT UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_target_chat_tg_chat_id ON target_chat(tg_chat_id);
---- create above / drop below ----
DROP INDEX IF EXISTS idx_target_chat_tg_chat_id;
DROP TABLE IF EXISTS target_chat;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
