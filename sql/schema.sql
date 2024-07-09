CREATE TABLE todo (
  id          uuid       DEFAULT uuid_generate_v4()  PRIMARY KEY NOT NULL,
  created_at  timestamp  DEFAULT NOW()               NOT NULL,
  updated_at  timestamp,
  title       varchar                                NOT NULL,
  content     varchar    DEFAULT ''                  NOT NULL
);
