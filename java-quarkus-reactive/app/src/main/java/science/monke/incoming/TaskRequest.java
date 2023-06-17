package science.monke.incoming;

import science.monke.outgoing.Task;

import java.util.UUID;

public record TaskRequest(String title, String description) {
  public static Task toTask(final TaskRequest taskRequest) {
    return toTask(UUID.randomUUID(), taskRequest);
  }

  public static Task toTask(final UUID taskId, final TaskRequest taskRequest) {
    final Task task = new Task();
    task.taskId = taskId;
    task.title = taskRequest.title;
    task.description = taskRequest.description;
    return task;
  }
}
