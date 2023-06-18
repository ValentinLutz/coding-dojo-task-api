package science.monke.incoming;

import java.util.List;
import java.util.UUID;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
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
  public ResponseEntity<List<TaskResponse>> getTasks() {
    List<TaskResponse> taskResponses = taskRepoPort.findAll().stream().map(TaskResponse::fromTask).toList();
    return ResponseEntity.ok(taskResponses);
  }

  @PostMapping(consumes = "application/json", produces = "application/json")
  public ResponseEntity<TaskResponse> postTask(@RequestBody final TaskRequest taskRequest) {
    TaskResponse taskResponse = TaskResponse.fromTask(taskRepoPort.save(TaskRequest.toTask(taskRequest)));
    return ResponseEntity.ok(taskResponse);
  }

  @GetMapping(path = "/{taskId}", produces = "application/json")
  public ResponseEntity<TaskResponse> getTask(@PathVariable final UUID taskId) {
    return taskRepoPort
        .findByTaskId(taskId)
        .map(TaskResponse::fromTask)
        .map(ResponseEntity::ok)
        .orElse(ResponseEntity.notFound().build());
  }

  @PutMapping(path = "/{taskId}", consumes = "application/json", produces = "application/json")
  public ResponseEntity<TaskResponse> putTask(@PathVariable final UUID taskId, @RequestBody final TaskRequest taskRequest) {
    TaskResponse taskResponse = TaskResponse.fromTask(taskRepoPort.save(TaskRequest.toTask(taskId, taskRequest)));
    return ResponseEntity.ok(taskResponse);
  }

  @DeleteMapping(path = "/{taskId}", produces = "application/json")
  public ResponseEntity<Void> deleteTask(@PathVariable final UUID taskId) {
    taskRepoPort.deleteById(taskId);
    return ResponseEntity.noContent().build();
  }
}
