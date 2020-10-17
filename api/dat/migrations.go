package dat

var migration string = `
DO $$ DECLARE
BEGIN

--
-- stage: if exist do not create migrations table
--

IF EXISTS(SELECT * FROM pg_tables WHERE tablename = 'migrations') THEN
	RAISE NOTICE 'migrations table exist, skipping...';
	RETURN;
END IF;

--
-- stage: migraitons table creation
--

CREATE TABLE migrations (
	name TEXT PRIMARY KEY,
	time TIMESTAMP DEFAULT NOW()
);

END $$;

--
-- stage: users table creation unless exist
--

DO $$ BEGIN
IF EXISTS(SELECT * FROM migrations WHERE name = 'users') THEN RETURN;
END IF;

CREATE TABLE users (
	id SERIAL PRIMARY KEY NOT NULL,
	created TIMESTAMP DEFAULT NOW(), 
	name TEXT NOT NULL, 
	email TEXT NOT NULL UNIQUE, 
	password TEXT NOT NULL,
	auth_method	TEXT
);

INSERT INTO migrations (name) VALUES ('users');

END $$;

--
-- stage: sessions table creation unless exsit
--

DO $$ BEGIN
IF EXISTS(SELECT * FROM migrations WHERE name = 'sessions') THEN RETURN;
END IF;

CREATE TABLE sessions (
	userid INTEGER,
	created  TIMESTAMP DEFAULT NOW(), 
	token	text,
	csrf	text,
	expires TIMESTAMP,
	CONSTRAINT fk_user
	   FOREIGN KEY(userid) 
	   REFERENCES users(id)
);

-- by alter table: ALTER TABLE sessions ADD CONSTRAINT fk_user FOREIGN KEY (userid) REFERENCES users (id);

INSERT INTO migrations (name) VALUES ('sessions');

END $$;
`
