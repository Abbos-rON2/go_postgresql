CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE "roles" AS ENUM (
  'super_admin',
  'branch_admin'
);

CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "guid" uuid UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
  "username" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "phone" varchar UNIQUE NOT NULL,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "is_active" bool NOT NULL DEFAULT true,
  "birth_date" date NOT NULL,
  "address" varchar NOT NULL,
  "role" roles NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now(),
  "deleted_at" timestamp DEFAULT null
);

CREATE TABLE "branches" (
  "id" SERIAL PRIMARY KEY,
  "guid" uuid UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
  "address" varchar NOT NULL,
  "phone" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now(),
  "deleted_at" timestamp DEFAULT null
);

CREATE TABLE "groups" (
  "id" SERIAL PRIMARY KEY,
  "guid" uuid UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
  "name" varchar NOT NULL,
  "subject_id" int,
  "schedule" jsonb NOT NULL,
  "teacher_id" int,
  "branch_id" int,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now(),
  "deleted_at" timestamp DEFAULT null
);

CREATE TABLE "teachers" (
  "id" SERIAL PRIMARY KEY,
  "guid" uuid UNIQUE NOT NULL,
  "phone" varchar UNIQUE NOT NULL,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "birth_date" date NOT NULL,
  "address" varchar NOT NULL,
  "salary" float NOT NULL,
  "branch_id" int,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now(),
  "deleted_at" timestamp DEFAULT null
);

CREATE TABLE "students" (
  "id" SERIAL PRIMARY KEY,
  "guid" uuid UNIQUE NOT NULL,
  "phone" varchar UNIQUE NOT NULL,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "birth_date" date NOT NULL,
  "address" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now(),
  "deleted_at" timestamp DEFAULT null
);

CREATE TABLE "subjects" (
  "id" SERIAL PRIMARY KEY,
  "guid" uuid UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
  "name" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now(),
  "deleted_at" timestamp DEFAULT null
);

CREATE TABLE "payments" (
  "id" SERIAL PRIMARY KEY,
  "guid" uuid UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
  "student_id" int,
  "group_id" int,
  "amount" float NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now(),
  "deleted_at" timestamp DEFAULT null
);

CREATE TABLE "attendances" (
  "id" SERIAL PRIMARY KEY,
  "guid" uuid UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
  "group_id" int,
  "student_id" int,
  "is_attend" bool DEFAULT false,
  "comment" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now(),
  "deleted_at" timestamp DEFAULT null
);

CREATE TABLE "group_students" (
  "group_id" int,
  "student_id" int
);

CREATE TABLE "subject_teachers" (
  "subject_id" int,
  "teacher_id" int
);

ALTER TABLE "groups" ADD FOREIGN KEY ("subject_id") REFERENCES "subjects" ("id");

ALTER TABLE "groups" ADD FOREIGN KEY ("teacher_id") REFERENCES "teachers" ("id");

ALTER TABLE "groups" ADD FOREIGN KEY ("branch_id") REFERENCES "branches" ("id");

ALTER TABLE "teachers" ADD FOREIGN KEY ("branch_id") REFERENCES "branches" ("id");

ALTER TABLE "payments" ADD FOREIGN KEY ("student_id") REFERENCES "students" ("id");

ALTER TABLE "payments" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("id");

ALTER TABLE "attendances" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("id");

ALTER TABLE "attendances" ADD FOREIGN KEY ("student_id") REFERENCES "students" ("id");

ALTER TABLE "group_students" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("id");

ALTER TABLE "group_students" ADD FOREIGN KEY ("student_id") REFERENCES "students" ("id");

ALTER TABLE "subject_teachers" ADD FOREIGN KEY ("subject_id") REFERENCES "subjects" ("id");

ALTER TABLE "subject_teachers" ADD FOREIGN KEY ("teacher_id") REFERENCES "teachers" ("id");