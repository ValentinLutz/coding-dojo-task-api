package science.monke.outgoing;

import java.util.List;
import java.util.Map;
import java.util.Optional;
import java.util.UUID;
import java.util.concurrent.ConcurrentHashMap;

public class TaskRepoMemory implements TaskRepoPort {

  private final Map<UUID, Task> tasks;

  public TaskRepoMemory() {
    tasks = new ConcurrentHashMap<>();
  }

  @Override
  public List<Task> findAll() {
    return List.copyOf(tasks.values());
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
  public Task update(final Task task) {
    return save(task);
  }

  @Override
  public void deleteById(final UUID taskId) {
    tasks.remove(taskId);
  }
}
