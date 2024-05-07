ALTER TABLE "user_logs" DROP CONSTRAINT user_logs_user_id_fkey;

ALTER TABLE "user_tokens" DROP CONSTRAINT user_tokens_user_id_fkey;

DROP TABLE users;
DROP TABLE user_logs;
DROP TABLE user_tokens;
