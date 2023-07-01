use crate::outgoing::task_repo::DynTaskRepository;
use axum::{
    extract::{Path, State},
    http::StatusCode,
    Json,
};

use super::models::{new_task_response_from_task_entity, TaskRequest, TaskResponse};

pub async fn get_tasks(
    State(task_repository): State<DynTaskRepository>,
) -> (StatusCode, Json<Vec<TaskResponse>>) {
    let task_entities = task_repository.lock().await.find_all().await;
    return (
        StatusCode::OK,
        Json(
            task_entities
                .into_iter()
                .map(new_task_response_from_task_entity)
                .collect(),
        ),
    );
}

pub async fn post_tasks(
    State(task_repository): State<DynTaskRepository>,
    Json(task_request): Json<TaskRequest>,
) -> (StatusCode, Json<TaskResponse>) {
    let task_entity = task_repository
        .lock()
        .await
        .save(task_request.to_new_task_entity())
        .await;
    return (
        StatusCode::CREATED,
        Json(new_task_response_from_task_entity(task_entity)),
    );
}

pub async fn get_task(
    State(task_repository): State<DynTaskRepository>,
    Path(task_id): Path<uuid::Uuid>,
) -> (StatusCode, Json<Option<TaskResponse>>) {
    let task_entity_option = task_repository.lock().await.find_by_task_id(task_id).await;
    return match task_entity_option {
        Some(task_entity) => (
            StatusCode::OK,
            Json(Some(new_task_response_from_task_entity(task_entity))),
        ),
        None => (StatusCode::NOT_FOUND, Json(None)),
    };
}

pub async fn put_task(
    State(task_repository): State<DynTaskRepository>,
    Path(task_id): Path<uuid::Uuid>,
    Json(task_request): Json<TaskRequest>,
) -> StatusCode {
    let task_entity_option = task_repository
        .lock()
        .await
        .update(task_request.to_updated_task_entity(task_id))
        .await;
    return match task_entity_option {
        Some(_) => StatusCode::NO_CONTENT,
        None => StatusCode::NOT_FOUND,
    };
}

pub async fn delete_task(
    State(task_repository): State<DynTaskRepository>,
    Path(task_id): Path<uuid::Uuid>,
) -> StatusCode {
    let task_entity_option = task_repository
        .lock()
        .await
        .delete_by_task_id(task_id)
        .await;
    return match task_entity_option {
        Some(_) => StatusCode::NO_CONTENT,
        None => StatusCode::NOT_FOUND,
    };
}
