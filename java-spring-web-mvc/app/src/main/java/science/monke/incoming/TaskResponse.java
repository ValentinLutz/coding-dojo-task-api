package science.monke.incoming;

import java.util.UUID;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonValue;
import science.monke.outgoing.Task;

public record TaskResponse(
    @JsonProperty("task_id") UUID taskId,
    @JsonProperty("title") String title,
    @JsonProperty("description") String description) {
  public static TaskResponse fromTask(final Task task) {
    return new TaskResponse(task.taskId, task.title, task.description);
  }
}
