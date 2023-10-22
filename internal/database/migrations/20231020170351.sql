-- Create "users" table
CREATE TABLE "public"."users" (
 "id" serial NOT NULL,
 "name" character varying(255) NOT NULL,
 "age" integer NOT NULL,
 "gender" character varying(255) NOT NULL,
 "country" character varying(255) NOT NULL,
 PRIMARY KEY ("id")
);
