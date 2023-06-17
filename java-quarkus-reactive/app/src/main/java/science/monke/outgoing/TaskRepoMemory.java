package science.monke.outgoing;

import io.quarkus.arc.lookup.LookupIfProperty;
import io.smallrye.mutiny.Multi;
import io.smallrye.mutiny.Uni;
import jakarta.enterprise.context.ApplicationScoped;

import java.util.Map;
import java.util.UUID;
import java.util.concurrent.ConcurrentHashMap;

@ApplicationScoped
@LookupIfProperty(name = "app.memory.enabled", stringValue = "true")
public class TaskRepoMemory implements TaskRepoPort {

  private final Map<UUID, Task> tasks;

  public TaskRepoMemory() {
    tasks = new ConcurrentHashMap<>();
  }

  public Multi<Task> findAll() {
    return Multi.createFrom().iterable(tasks.values());
  }

  public Uni<Task> findByTaskId(final UUID taskId) {
    return Uni.createFrom().item(tasks.get(taskId));
  }

  public Uni<Task> save(final Task task) {
    tasks.put(task.taskId, task);
    return Uni.createFrom().item(task);
  }

  public Uni<Task> update(final Task task) {
    return save(task);
  }

  public Uni<Void> delete(final UUID taskId) {
    return Uni.createFrom().item(tasks.remove(taskId)).onItem().transform(task -> null);
  }
}
