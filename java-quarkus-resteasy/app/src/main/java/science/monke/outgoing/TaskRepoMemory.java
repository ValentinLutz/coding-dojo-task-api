package science.monke.outgoing;

import io.quarkus.arc.lookup.LookupIfProperty;
import jakarta.enterprise.context.ApplicationScoped;
import java.util.*;
import java.util.concurrent.ConcurrentHashMap;

@ApplicationScoped
@LookupIfProperty(name = "app.memory.enabled", stringValue = "true")
public class TaskRepoMemory implements TaskRepoPort {

  private final Map<UUID, Task> tasks;

  public TaskRepoMemory() {
    tasks = new ConcurrentHashMap<>();
  }

  @Override
  public List<Task> findAll() {
    return new ArrayList<>(tasks.values());
  }

  @Override
  public Optional<Task> findByTaskId(final UUID taskId) {
    return Optional.ofNullable(tasks.get(taskId));
  }

  @Override
  public Task save(final Task task) {
    tasks.put(task.taskId, task);
    return task;
  }

  @Override
  public Optional<Task> update(final Task task) {
    final Task previousTask = tasks.replace(task.taskId, task);
    if (previousTask == null) {
      return Optional.empty();
    }
    return Optional.of(task);
  }

  @Override
  public boolean delete(final UUID taskId) {
    return tasks.remove(taskId) != null;
  }
}
