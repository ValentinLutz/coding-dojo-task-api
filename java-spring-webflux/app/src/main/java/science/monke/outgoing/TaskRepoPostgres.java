package science.monke.outgoing;

import java.util.UUID;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.autoconfigure.condition.ConditionalOnProperty;
import org.springframework.data.r2dbc.core.R2dbcEntityTemplate;
import org.springframework.data.relational.core.query.Criteria;
import org.springframework.data.relational.core.query.Query;
import org.springframework.stereotype.Repository;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

@Repository
@ConditionalOnProperty(name = "app.memory.enabled", havingValue = "false")
public class TaskRepoPostgres implements TaskRepoPort {

  private final R2dbcEntityTemplate r2dbcEntityTemplate;

  @Autowired
  public TaskRepoPostgres(R2dbcEntityTemplate r2dbcEntityTemplate) {
    this.r2dbcEntityTemplate = r2dbcEntityTemplate;
  }

  @Override
  public Flux<Task> findAll() {
    return r2dbcEntityTemplate.select(Task.class).all();
  }

  @Override
  public Mono<Task> findByTaskId(UUID taskId) {
    return r2dbcEntityTemplate.selectOne(
        Query.query(Criteria.where("task_id").is(taskId)), Task.class);
  }

  @Override
  public Mono<Task> save(Task task) {
    return r2dbcEntityTemplate.insert(task);
  }

  @Override
  public Mono<Task> update(Task task) {
    return r2dbcEntityTemplate.update(task);
  }

  @Override
  public Mono<Void> deleteById(UUID taskId) {
    return r2dbcEntityTemplate
        .delete(Task.class)
        .matching(Query.query(Criteria.where("task_id").is(taskId)))
        .all()
        .then();
  }

  @Override
  public Mono<Boolean> existsById(UUID taskId) {
    return r2dbcEntityTemplate.exists(
        Query.query(Criteria.where("task_id").is(taskId)), Task.class);
  }
}
