CREATE TABLE courses (
    code        VARCHAR(10) PRIMARY KEY,
    title       TEXT        NOT NULL,
    credits     INTEGER     NOT NULL,
    enrolled    INTEGER     NOT NULL DEFAULT 0,
    instructors TEXT[]      NOT NULL 
);

CREATE INDEX idx_courses_title ON courses (title);