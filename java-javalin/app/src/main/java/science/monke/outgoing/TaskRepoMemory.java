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
  public Optional<Task> update(final Task task) {
    final Task removedTask = tasks.replace(task.taskId, task);
    return Optional.ofNullable(removedTask);
  }

  @Override
  public boolean deleteById(final UUID taskId) {
    final Task removedTask = tasks.remove(taskId);
    return removedTask != null;
  }
}
