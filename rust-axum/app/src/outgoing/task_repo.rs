use crate::outgoing::entities::TaskEntity;
use async_trait::async_trait;
use std::sync::Arc;
use tokio::sync::Mutex;

pub type DynTaskRepository = Arc<Mutex<dyn TaskRepository + Send + Sync>>;

#[async_trait]
pub trait TaskRepository {
    async fn find_all(&self) -> Vec<TaskEntity>;
    async fn find_by_task_id(&self, task_id: uuid::Uuid) -> Option<TaskEntity>;
    async fn save(&mut self, task: TaskEntity) -> TaskEntity;
    async fn update(&mut self, task: TaskEntity) -> Option<TaskEntity>;
    async fn delete_by_task_id(&mut self, task_id: uuid::Uuid) -> Option<TaskEntity>;
}
