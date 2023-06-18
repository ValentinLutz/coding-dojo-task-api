DROP TABLE IF EXISTS public.tasks;

CREATE TABLE IF NOT EXISTS public.tasks (
    task_id UUID PRIMARY KEY NOT NULL,
    title TEXT,
    description TEXT
);