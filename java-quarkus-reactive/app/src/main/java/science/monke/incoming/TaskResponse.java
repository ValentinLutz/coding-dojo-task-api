package science.monke.incoming;

import com.fasterxml.jackson.annotation.JsonProperty;
import science.monke.outgoing.Task;

import java.util.UUID;

public record TaskResponse(
    @JsonProperty("task_id") UUID taskId,
    @JsonProperty("title") String title,
    @JsonProperty("description") String description) {

  public static TaskResponse fromTask(final Task task) {
    return new TaskResponse(task.taskId, task.title, task.description);
  }
}
