CREATE TABLE getstronger.auth
(
    id            UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    email         VARCHAR(128)     NOT NULL UNIQUE,
    password      BYTEA            NOT NULL,
    refresh_token VARCHAR(256)     NULL,
    created_at    TIMESTAMP        NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC')
);

CREATE TABLE getstronger.users
(
    id         UUID PRIMARY KEY NOT NULL REFERENCES getstronger.auth (id),
    first_name VARCHAR          NOT NULL,
    last_name  VARCHAR          NOT NULL,
    created_at TIMESTAMP        NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC')
);

CREATE TABLE getstronger.routines
(
    id         UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    user_id    UUID             NOT NULL REFERENCES getstronger.users (id),
    title      VARCHAR          NOT NULL,
    created_at TIMESTAMP        NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    deleted_at TIMESTAMP        NULL
);

CREATE TABLE getstronger.exercises
(
    id                UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    user_id           UUID             NOT NULL REFERENCES getstronger.users (id),
    title             VARCHAR          NOT NULL,
    sub_title         VARCHAR,
    rest_between_sets SMALLINT         NULL,
    created_at        TIMESTAMP        NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    deleted_at        TIMESTAMP        NULL
);

CREATE TABLE getstronger.routine_exercises
(
    routine_id  UUID NOT NULL REFERENCES getstronger.routines (id),
    exercise_id UUID NOT NULL REFERENCES getstronger.exercises (id),
    PRIMARY KEY (routine_id, exercise_id)
);

CREATE TABLE getstronger.workouts
(
    id         UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    user_id    UUID             NOT NULL REFERENCES getstronger.users (id),
    routine_id UUID             NOT NULL REFERENCES getstronger.routines (id),
    date       DATE             NOT NULL,
    created_at TIMESTAMP        NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC')
);

CREATE TABLE getstronger.sets
(
    id          UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    workout_id  UUID             NOT NULL REFERENCES getstronger.workouts (id),
    exercise_id UUID             NOT NULL REFERENCES getstronger.exercises (id),
    weight      DECIMAL(8, 2)    NOT NULL,
    reps        INT              NOT NULL,
    created_at  TIMESTAMP        NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC')
);
