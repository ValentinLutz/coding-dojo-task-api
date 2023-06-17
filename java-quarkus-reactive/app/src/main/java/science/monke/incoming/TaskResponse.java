package science.monke.incoming;

import science.monke.outgoing.Task;

import java.util.UUID;

public record TaskResponse(UUID taskId, String title, String description) {

  public static TaskResponse fromTask(final Task task) {
    return new TaskResponse(task.taskId, task.title, task.description);
  }
}
