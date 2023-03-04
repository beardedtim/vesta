/*
  Users are the end-users that actively use the application
*/
CREATE TABLE IF NOT EXISTS users (
  _id UUID NOT NULL DEFAULT gen_random_uuid(),
  email TEXT UNIQUE NOT NULL,
  password TEXT NOT NULL,
  profile JSONB NOT NULL DEFAULT '{}',
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY (_id)
);

/*
  Communities are the grouping of Users in a specific context
*/
CREATE TABLE IF NOT EXISTS communities (
  _id UUID NOT NULL DEFAULT gen_random_uuid(),
  name TEXT NOT NULL,
  slug TEXT UNIQUE NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  profile JSONB NOT NULL DEFAULT '{}',
  PRIMARY KEY (_id)
);

/*
  A User can be in many Communities and a Community
  can have many Users 
*/
CREATE TABLE IF NOT EXISTS user_communities (
  user_id UUID NOT NULL,
  community_id UUID NOT NULL,
  FOREIGN KEY (user_id) 
    REFERENCES users(_id) 
    ON DELETE CASCADE,
  FOREIGN KEY (community_id)
    REFERENCES communities(_id)
    ON DELETE CASCADE,
  UNIQUE(user_id, community_id)
);

/*
  Users can have Roles within the system in order to
  attach different permissions
*/
CREATE TABLE IF NOT EXISTS roles (
  _id UUID NOT NUL DEFAULT gen_random_uuid(),
  name TEXT UNIQUE NOT NULL,
  description TEXT,
  PRIMARY KEY (_id)
);

/*
  A User can have 0->Many Roles within 0->Many Communities
  Roles can be attached to 0->Many Users withing 0->Many Communities
*/
CREATE TABLE IF NOT EXISTS user_roles (
  user_id UUID NOT NULL,
  role_id UUID NOT NULL,
  community_id UUID NOT NULL,
  FOREIGN KEY (user_id) 
    REFERENCES users(_id) 
    ON DELETE CASCADE,
  FOREIGN KEY (role_id)
    REFERENCES roles(_id)
    ON DELETE CASCADE,
  FOREIGN KEY (community_id)
    REFERENCES communities(_id)
    ON DELETE CASCADE,
  UNIQUE(user_id, role_id, community_id)
);
