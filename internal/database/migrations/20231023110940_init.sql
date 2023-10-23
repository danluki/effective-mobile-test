-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE "public"."users" (
 "id" serial NOT NULL,
 "name" character varying(255) NOT NULL,
 "age" integer NOT NULL,
 "gender" character varying(255) NOT NULL,
 "country" character varying(255) NOT NULL,
 PRIMARY KEY ("id")
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE IF EXISTS "users" CASCADE; 
-- +goose StatementEnd
