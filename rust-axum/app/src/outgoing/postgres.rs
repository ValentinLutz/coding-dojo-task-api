use crate::outgoing::entities::TaskEntity;
use crate::outgoing::task_repo::TaskRepository;
use async_trait::async_trait;

use uuid::Uuid;

pub struct PostgresTaskRepository {
    postgres_pool: sqlx::PgPool,
}

impl PostgresTaskRepository {
    pub fn new(postgres_pool: sqlx::PgPool) -> PostgresTaskRepository {
        return PostgresTaskRepository {
            postgres_pool: postgres_pool,
        };
    }
}

#[async_trait]
impl TaskRepository for PostgresTaskRepository {
    async fn find_all(&self) -> Vec<TaskEntity> {
        return sqlx::query_as::<_, TaskEntity>("SELECT task_id, title, description FROM public.tasks")
            .fetch_all(&self.postgres_pool)
            .await
            .unwrap();
    }

    async fn find_by_task_id(&self, task_id: Uuid) -> Option<TaskEntity> {
        return sqlx::query_as::<_, TaskEntity>(
            "SELECT task_id, title, description FROM public.tasks WHERE task_id = $1",
        )
        .bind(task_id)
        .fetch_optional(&self.postgres_pool)
        .await
        .unwrap();
    }

    async fn save(&mut self, task: TaskEntity) -> TaskEntity {
        return sqlx::query_as::<_, TaskEntity>(
            "INSERT INTO public.tasks (task_id, title, description) VALUES ($1, $2, $3) RETURNING task_id, title, description",
        )
        .bind(task.task_id)
        .bind(task.title)
        .bind(task.description)
        .fetch_one(&self.postgres_pool)
        .await
        .unwrap();
    }

    async fn update(&mut self, task: TaskEntity) -> Option<TaskEntity> {
        return sqlx::query_as::<_, TaskEntity>(
            "UPDATE public.tasks SET title = $1, description = $2 WHERE task_id = $3 RETURNING task_id, title, description",
        )
        .bind(task.title)
        .bind(task.description)
        .bind(task.task_id)
        .fetch_one(&self.postgres_pool)
        .await
        .ok();
    }

    async fn delete_by_task_id(&mut self, task_id: Uuid) -> Option<TaskEntity> {
        return sqlx::query_as::<_, TaskEntity>(
            "DELETE FROM public.tasks WHERE task_id = $1 RETURNING task_id, title, description",
        )
        .bind(task_id)
        .fetch_optional(&self.postgres_pool)
        .await
        .unwrap();
    }
}
