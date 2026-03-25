CREATE TABLE students (
    id        BIGSERIAL    PRIMARY KEY,
    name      VARCHAR(100) NOT NULL,
    programme TEXT         NOT NULL,
    year      SMALLINT     NOT NULL CHECK (year BETWEEN 1 AND 4),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_students_name ON students (name);