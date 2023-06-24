package science.monke.incoming;

import java.util.UUID;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;
import science.monke.outgoing.TaskRepoPort;

@RestController
@RequestMapping(path = "/tasks")
public class TaskApi {

  private final TaskRepoPort taskRepoPort;

  @Autowired
  public TaskApi(final TaskRepoPort taskRepoPort) {
    this.taskRepoPort = taskRepoPort;
  }

  @GetMapping(produces = "application/json")
  public Flux<TaskResponse> getTasks() {
    return taskRepoPort.findAll().map(TaskResponse::fromTask);
  }

  @PostMapping(consumes = "application/json", produces = "application/json")
  public Mono<ResponseEntity<TaskResponse>> postTask(@RequestBody final TaskRequest taskRequest) {
    return taskRepoPort
        .save(TaskRequest.toTask(taskRequest))
        .map(TaskResponse::fromTask)
        .flatMap(taskResponse -> Mono.just(ResponseEntity.status(201).body(taskResponse)));
  }

  @GetMapping(path = "/{taskId}", produces = "application/json")
  public Mono<ResponseEntity<TaskResponse>> getTask(@PathVariable final UUID taskId) {
    return taskRepoPort
        .findByTaskId(taskId)
        .map(TaskResponse::fromTask)
        .flatMap(taskResponse -> Mono.just(ResponseEntity.ok().body(taskResponse)))
        .switchIfEmpty(Mono.just(ResponseEntity.notFound().build()));
  }

  @PutMapping(path = "/{taskId}", consumes = "application/json", produces = "application/json")
  public Mono<ResponseEntity<TaskResponse>> putTask(
      @PathVariable final UUID taskId, @RequestBody final TaskRequest taskRequest) {
    return taskRepoPort
        .existsById(taskId)
        .flatMap(
            exists -> {
              if (Boolean.FALSE.equals(exists)) {
                return Mono.just(ResponseEntity.notFound().build());
              }
              return taskRepoPort
                  .update(TaskRequest.toTask(taskId, taskRequest))
                  .map(TaskResponse::fromTask)
                  .flatMap(taskResponse -> Mono.just(ResponseEntity.noContent().build()));
            });
  }

  @DeleteMapping(path = "/{taskId}", produces = "application/json")
  public Mono<ResponseEntity<TaskResponse>> deleteTask(@PathVariable final UUID taskId) {
    return taskRepoPort
        .existsById(taskId)
        .flatMap(
            exists -> {
              if (Boolean.FALSE.equals(exists)) {
                return Mono.just(ResponseEntity.notFound().build());
              }
              return taskRepoPort
                  .deleteById(taskId)
                  .then(Mono.defer(() -> Mono.just(ResponseEntity.noContent().build())));
            });
  }
}
