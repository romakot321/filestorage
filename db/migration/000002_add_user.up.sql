CREATE SEQUENCE users_id_seq;

CREATE TABLE "users"(
    "id" INT PRIMARY KEY DEFAULT nextval('users_id_seq'),
    "password_hash" VARCHAR NOT NULL,
    "name" VARCHAR NOT NULL UNIQUE,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER SEQUENCE users_id_seq
OWNED BY users.id;
