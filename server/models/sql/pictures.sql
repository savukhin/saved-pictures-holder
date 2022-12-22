CREATE TABLE IF NOT EXISTS "pictures" (
    "id" SERIAL PRIMARY KEY,
    "user_id" INTEGER NOT NULL,
    "title" VARCHAR(255),
    "description" VARCHAR(255),
    "folder_id" INTEGER NOT NULL,
    "file_name" VARCHAR(255) NOT NULL,

    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP,

    FOREIGN KEY ("user_id") REFERENCES "users" ("id"),
    FOREIGN KEY ("folder_id") REFERENCES "folders" ("id")
);
