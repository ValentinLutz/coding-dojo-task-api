package science.monke.outgoing;


import java.util.List;
import java.util.Optional;
import java.util.UUID;

public interface TaskRepoPort {
  List<Task> findAll();

  Optional<Task> findByTaskId(final UUID taskId);

  Task save(final Task task);

  void deleteById(final UUID taskId);
}
