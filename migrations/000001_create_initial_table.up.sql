CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS auth (
  id          uuid       DEFAULT uuid_generate_v4()   PRIMARY KEY NOT NULL,
  created_at  timestamp  DEFAULT NOW()                NOT NULL,
  updated_at  timestamp,
  email       varchar                                 UNIQUE NOT NULL,
  password    varchar                                 NOT NULL,
  username    varchar                                 NOT NULL
);

CREATE TABLE IF NOT EXISTS todo (
  id          uuid       DEFAULT uuid_generate_v4()  PRIMARY KEY NOT NULL,
  created_at  timestamp  DEFAULT NOW()               NOT NULL,
  updated_at  timestamp  DEFAULT NOW(),
  auth_id     uuid                                   REFERENCES auth (id),
  title       varchar                                NOT NULL,
  content     varchar
);

-- for update_at
CREATE OR REPLACE FUNCTION update_time()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_time_trigger
BEFORE UPDATE ON todo
FOR EACH ROW
EXECUTE PROCEDURE update_time();
