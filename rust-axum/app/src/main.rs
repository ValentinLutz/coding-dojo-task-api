use crate::outgoing::task_repo::DynTaskRepository;
use crate::outgoing::{memory::MemoryTaskRepository, postgres::PostgresTaskRepository};
use axum::{
    routing::{delete, get, post, put},
    Router,
};
use incoming::task_api;
use sqlx::postgres::PgPoolOptions;
use tokio::sync::Mutex;

use std::{env, net::SocketAddr, sync::Arc};

mod incoming;
mod outgoing;

#[tokio::main]
async fn main() {
    let use_in_memory = env::var("USE_IN_MEMORY").unwrap();
    let port = env::var("PORT").unwrap().parse().unwrap();

    let task_repo: DynTaskRepository;
    if use_in_memory == "true" {
        println!("Using memory repository");
        task_repo = Arc::new(Mutex::new(MemoryTaskRepository::new()));
    } else {
        println!("Using postgres repository");
        let postgres_pool = create_postgres_pool().await;
        init_task_table(postgres_pool.to_owned()).await;
        task_repo = Arc::new(Mutex::new(PostgresTaskRepository::new(postgres_pool)));
    }

    let app = Router::new()
        .route("/tasks", get(task_api::get_tasks))
        .route("/tasks", post(task_api::post_tasks))
        .route("/tasks/:task_id", get(task_api::get_task))
        .route("/tasks/:task_id", put(task_api::put_task))
        .route("/tasks/:task_id", delete(task_api::delete_task))
        .with_state(task_repo);

    let addr = SocketAddr::from(([0, 0, 0, 0], port));
    axum::Server::bind(&addr)
        .serve(app.into_make_service())
        .await
        .unwrap();
}

async fn create_postgres_pool() -> sqlx::PgPool {
    let host = env::var("POSTGRES_HOST").unwrap();
    let port = env::var("POSTGRES_PORT").unwrap();
    let user = env::var("POSTGRES_USER").unwrap();
    let password = env::var("POSTGRES_PASSWORD").unwrap();
    let database = env::var("POSTGRES_DATABASE").unwrap();

    let postgres_url = format!(
        "postgres://{}:{}@{}:{}/{}",
        user, password, host, port, database
    );

    return PgPoolOptions::new()
        .max_connections(100)
        .idle_timeout(std::time::Duration::from_secs(60))
        .connect(postgres_url.as_str())
        .await
        .unwrap();
}

async fn init_task_table(postgres_pool: sqlx::PgPool) {
    sqlx::query("DROP TABLE IF EXISTS public.tasks")
        .execute(&postgres_pool)
        .await
        .unwrap();
    print!("deleted table public.tasks");
    sqlx::query(
        r#"
            CREATE TABLE IF NOT EXISTS public.tasks (
                task_id UUID PRIMARY KEY,
                title VARCHAR(255) NOT NULL,
                description VARCHAR(255) NOT NULL
            )
            "#,
    )
    .execute(&postgres_pool)
    .await
    .unwrap();
    print!("created table public.tasks");
}
