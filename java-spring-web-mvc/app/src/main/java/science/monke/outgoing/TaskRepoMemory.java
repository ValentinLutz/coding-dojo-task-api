package science.monke.outgoing;

import org.springframework.boot.autoconfigure.condition.ConditionalOnProperty;
import org.springframework.stereotype.Component;

import java.util.List;
import java.util.Map;
import java.util.Optional;
import java.util.UUID;
import java.util.concurrent.ConcurrentHashMap;

@Component
@ConditionalOnProperty(name = "app.memory.enabled", havingValue = "true")
public class TaskRepoMemory implements TaskRepoPort {

  private final Map<UUID, Task> tasks;

  public TaskRepoMemory() {
    tasks = new ConcurrentHashMap<>();
  }

  public List<Task> findAll() {
    return List.copyOf(tasks.values());
  }

  public Optional<Task> findByTaskId(final UUID taskId) {
    return Optional.ofNullable(tasks.get(taskId));
  }

  public Task save(final Task task) {
    tasks.put(task.taskId, task);
    return task;
  }

  public void deleteById(final UUID taskId) {
    tasks.remove(taskId);
  }
}
