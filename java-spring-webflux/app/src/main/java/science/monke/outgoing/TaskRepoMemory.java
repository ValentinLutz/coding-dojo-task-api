package science.monke.outgoing;

import java.util.Map;
import java.util.UUID;
import java.util.concurrent.ConcurrentHashMap;
import org.springframework.boot.autoconfigure.condition.ConditionalOnProperty;
import org.springframework.stereotype.Component;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

@Component
@ConditionalOnProperty(name = "app.memory.enabled", havingValue = "true")
public class TaskRepoMemory implements TaskRepoPort {

  private final Map<UUID, Task> tasks;

  public TaskRepoMemory() {
    tasks = new ConcurrentHashMap<>();
  }

  @Override
  public Flux<Task> findAll() {
    return Flux.fromIterable(tasks.values());
  }

  @Override
  public Mono<Task> findByTaskId(final UUID taskId) {
    return Mono.justOrEmpty(tasks.get(taskId));
  }

  @Override
  public Mono<Task> save(final Task task) {
    tasks.put(task.taskId, task);
    return Mono.just(task);
  }

  @Override
  public Mono<Task> update(Task task) {
    tasks.replace(task.taskId, task);
    return Mono.just(task);
  }

  @Override
  public Mono<Void> deleteById(final UUID taskId) {
    tasks.remove(taskId);
    return Mono.empty();
  }

  @Override
  public Mono<Boolean> existsById(UUID taskId) {
    return Mono.just(tasks.containsKey(taskId));
  }
}
