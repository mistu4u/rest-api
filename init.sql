CREATE TABLE public.message (
	id uuid PRIMARY KEY,
	message_text text NULL
);

INSERT INTO message VALUES ('11e41e52-0333-47c2-af3b-791a68ba6ad0','Hello from docker compose db');