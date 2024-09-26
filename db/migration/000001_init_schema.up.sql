CREATE TABLE "files"(
    "filename" UUID NOT NULL DEFAULT (uuid_generate_v4()),
    "owner_id" INT NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT "files_pkey" PRIMARY KEY("filename")
);
