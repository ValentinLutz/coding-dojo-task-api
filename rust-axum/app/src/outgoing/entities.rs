use sqlx::FromRow;
use uuid::Uuid;

#[derive(Clone, FromRow)]
pub struct TaskEntity {
    pub task_id: Uuid,
    pub title: String,
    pub description: Option<String>,
}
