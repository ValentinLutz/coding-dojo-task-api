use serde::{Deserialize, Serialize};
use uuid::Uuid;

use crate::outgoing::entities::TaskEntity;

#[derive(Deserialize)]
pub struct TaskRequest {
    pub title: String,
    pub description: Option<String>,
}

impl TaskRequest {
    pub fn to_new_task_entity(&self) -> TaskEntity {
        return TaskEntity {
            task_id: Uuid::new_v4(),
            title: self.title.clone(),
            description: self.description.clone(),
        };
    }

    pub fn to_updated_task_entity(&self, task_id: uuid::Uuid) -> TaskEntity {
        return TaskEntity {
            task_id: task_id,
            title: self.title.clone(),
            description: self.description.clone(),
        };
    }
}

#[derive(Serialize)]
pub struct TaskResponse {
    pub task_id: Uuid,
    pub title: String,
    pub description: Option<String>,
}

pub fn new_task_response_from_task_entity(task: TaskEntity) -> TaskResponse {
    return TaskResponse {
        task_id: task.task_id,
        title: task.title,
        description: task.description,
    };
}
