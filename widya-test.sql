-- public."user" definition

-- Drop table

-- DROP TABLE public."user";

CREATE TABLE public."user" (
	id text NOT NULL,
	username varchar NULL,
	email varchar NULL,
	"password" varchar NULL,
	gender varchar NULL,
	CONSTRAINT user_pk PRIMARY KEY (id)
);


-- public.stored_token definition

-- Drop table

-- DROP TABLE public.stored_token;

CREATE TABLE public.stored_token (
	id_user text NOT NULL,
	"token" text NOT NULL,
	CONSTRAINT stored_token_pk PRIMARY KEY (id_user)
);