CREATE TABLE IF NOT EXISTS "pictures" (
    "id" SERIAL PRIMARY KEY,
    "user_id" INTEGER NOT NULL,
    "title" VARCHAR(255),
    "description" VARCHAR(255),
    "url" VARCHAR(255) NOT NULL,
    "folder_id" INTEGER,

    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP

    FOREIGN KEY ("user_id") REFERENCES "user" ("id")
    FOREIGN KEY ("folder_id") REFERENCES "folder" ("id")
    PRIMARY KEY ("id")
);
