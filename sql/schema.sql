CREATE TABLE auth (
  id          uuid       DEFAULT uuid_generate_v4()   PRIMARY KEY NOT NULL,
  created_at  timestamp  DEFAULT NOW()                NOT NULL,
  updated_at  timestamp,
  email       varchar                                 UNIQUE NOT NULL,
  password    varchar                                 NOT NULL,
  username    varchar                                 NOT NULL
);

CREATE TABLE todo (
  id          uuid       DEFAULT uuid_generate_v4()  PRIMARY KEY NOT NULL,
  created_at  timestamp  DEFAULT NOW()               NOT NULL,
  updated_at  timestamp,
  auth_id     uuid                                   REFERENCES auth (id),
  title       varchar                                NOT NULL,
  content     varchar    DEFAULT ''                  NOT NULL
);
