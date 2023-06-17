package science.monke.outgoing;

import io.quarkus.arc.lookup.LookupIfProperty;
import io.quarkus.runtime.StartupEvent;
import io.smallrye.mutiny.Multi;
import io.smallrye.mutiny.Uni;
import io.vertx.mutiny.pgclient.PgPool;
import org.jboss.logging.Logger;
import io.vertx.mutiny.sqlclient.RowSet;
import io.vertx.mutiny.sqlclient.Tuple;
import jakarta.enterprise.context.ApplicationScoped;
import jakarta.enterprise.event.Observes;
import jakarta.inject.Inject;

import java.util.UUID;

@ApplicationScoped
@LookupIfProperty(name = "app.memory.enabled", stringValue = "false")
public class TaskRepoPostgres implements TaskRepoPort {

  private static final Logger LOGGER = Logger.getLogger(TaskRepoPostgres.class);

  private final PgPool pgPool;

  @Inject
  public TaskRepoPostgres(final PgPool pgPool) {
    this.pgPool = pgPool;
  }

  void config(@Observes final StartupEvent startupEvent) {
    pgPool.query("DROP TABLE IF EXISTS public.tasks").execute().await().indefinitely();
    LOGGER.info("Dropped table tasks");
    pgPool
        .query(
            "CREATE TABLE IF NOT EXISTS public.tasks (task_id UUID PRIMARY KEY NOT NULL, title TEXT, description TEXT)")
        .execute()
        .await()
        .indefinitely();
    LOGGER.info("Created table tasks");
  }

  public Multi<Task> findAll() {
    return pgPool
        .query("SELECT task_id, title, description FROM public.tasks")
        .execute()
        .onItem()
        .transformToMulti(rows -> Multi.createFrom().iterable(rows))
        .onItem()
        .transform(Task::fromRow);
  }

  public Uni<Task> findByTaskId(final UUID taskId) {
    return pgPool
        .preparedQuery("SELECT task_id, title, description FROM public.tasks WHERE task_id = $1")
        .execute(Tuple.of(taskId))
        .onItem()
        .transform(RowSet::iterator)
        .onItem()
        .transform(iterator -> iterator.hasNext() ? Task.fromRow(iterator.next()) : null);
  }

  public Uni<Task> save(final Task task) {
    return pgPool
        .preparedQuery("INSERT INTO public.tasks (task_id, title, description) VALUES ($1, $2, $3)")
        .execute(Tuple.of(task.taskId, task.title, task.description))
        .onItem()
        .transform(rowSet -> task);
  }

  public Uni<Task> update(final Task task) {
    return pgPool
        .preparedQuery("UPDATE public.tasks SET title = $1, description = $2 WHERE task_id = $3")
        .execute(Tuple.of(task.title, task.description, task.taskId))
        .onItem()
        .transform(rowSet -> task);
  }

  public Uni<Void> delete(final UUID taskId) {
    return pgPool
        .preparedQuery("DELETE FROM public.tasks WHERE task_id = $1")
        .execute(Tuple.of(taskId))
        .onItem()
        .transform(rowSet -> null);
  }
}
