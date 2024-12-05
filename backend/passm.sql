-- Database: passm

-- DROP DATABASE IF EXISTS passm;

CREATE DATABASE passm
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_IN'
    LC_CTYPE = 'en_IN'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

COMMENT ON DATABASE passm
    IS 'database for password manager application';

CREATE TABLE IF NOT EXISTS public.users
(
    id integer NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    username character varying(255) COLLATE pg_catalog."default" NOT NULL,
    email character varying(255) COLLATE pg_catalog."default" NOT NULL,
    password_enc text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT users_email_key UNIQUE (email),
    CONSTRAINT users_username_key UNIQUE (username)
)

CREATE TABLE IF NOT EXISTS public.passwords
(
    username character varying(255) COLLATE pg_catalog."default" NOT NULL,
    sitename character varying(255) COLLATE pg_catalog."default" NOT NULL,
    encrypted_password character varying(255) COLLATE pg_catalog."default" NOT NULL,
    last_updated timestamp without time zone DEFAULT CURRENT_TIMESTAMP
)