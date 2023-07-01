use crate::outgoing::entities::TaskEntity;
use crate::outgoing::task_repo::TaskRepository;
use async_trait::async_trait;
use std::collections::HashMap;
use uuid::Uuid;

pub struct MemoryTaskRepository {
    tasks: HashMap<Uuid, TaskEntity>,
}

impl MemoryTaskRepository {
    pub fn new() -> MemoryTaskRepository {
        return MemoryTaskRepository {
            tasks: HashMap::new(),
        };
    }
}

#[async_trait]
impl TaskRepository for MemoryTaskRepository {
    async fn find_all(&self) -> Vec<TaskEntity> {
        return self.tasks.values().cloned().collect();
    }

    async fn find_by_task_id(&self, task_id: Uuid) -> Option<TaskEntity> {
        return self.tasks.get(&task_id).cloned();
    }

    async fn save(&mut self, task: TaskEntity) -> TaskEntity {
        self.tasks.insert(task.task_id, task.clone());
        return task;
    }

    async fn update(&mut self, task: TaskEntity) -> Option<TaskEntity> {
        let contains_task = self.tasks.contains_key(&task.task_id);
        if !contains_task {
            return None;
        }
        self.tasks.insert(task.task_id, task.clone());
        return Some(task);
    }

    async fn delete_by_task_id(&mut self, task_id: Uuid) -> Option<TaskEntity> {
        let removed_task = self.tasks.remove(&task_id);
        return removed_task;
    }
}
