package science.monke.outgoing;

import io.smallrye.mutiny.Multi;
import io.smallrye.mutiny.Uni;
import java.util.UUID;

public interface TaskRepoPort {
  Multi<Task> findAll();

  Uni<Task> findByTaskId(final UUID taskId);

  Uni<Task> save(final Task task);

  Uni<Task> update(final Task task);

  Uni<Void> delete(final UUID taskId);
}
