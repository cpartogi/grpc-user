CREATE TABLE users (
	id uuid NOT NULL,
	full_name varchar(60) NOT NULL,
    email varchar(100) NOT NULL,
    phone_number varchar(13) NOT NULL,
    user_password varchar(255) NOT NULL,
    created_by varchar(255) NOT NULL,
	created_at timestamptz NOT NULL,
	updated_by varchar(255) NULL,
	updated_at timestamptz NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id)
);

CREATE TABLE user_logs (
	id uuid NOT NULL,
    user_id uuid NOT NULL,
    is_success bool DEFAULT FALSE,
    login_message varchar(500),
	created_at timestamptz NULL,
    CONSTRAINT user_logs_pkey PRIMARY KEY (id)
);

create table user_tokens (
	id uuid NOT NULL,
    user_id uuid NOT NULL,
    token varchar(500),
    token_expired_at timestamptz NOT NULL,
    refresh_token varchar(500),
    refresh_token_expired_at timestamptz NOT NULL,
	created_at timestamptz NOT NULL,
    updated_at timestamptz NULL,
    CONSTRAINT user_tokens_pkey PRIMARY KEY (id)
);

ALTER TABLE "user_logs" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_tokens" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");