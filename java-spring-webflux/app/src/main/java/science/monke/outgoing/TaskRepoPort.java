package science.monke.outgoing;

import java.util.UUID;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

public interface TaskRepoPort {
  Flux<Task> findAll();

  Mono<Task> findByTaskId(final UUID taskId);

  Mono<Task> save(final Task task);

  Mono<Task> update(final Task task);

  Mono<Void> deleteById(final UUID taskId);

  Mono<Boolean> existsById(final UUID taskId);
}
