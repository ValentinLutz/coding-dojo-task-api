CREATE TABLE task_service.task (
    task_id UUID NOT NULL UNIQUE,
    title VARCHAR NOT NULL,
    description VARCHAR,
    PRIMARY KEY (task_id)
);