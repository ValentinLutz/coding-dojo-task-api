package science.monke.outgoing;

import io.vertx.mutiny.sqlclient.Row;

import java.util.UUID;

public class Task {
  public UUID taskId;
  public String title;
  public String description;

  public static Task fromRow(Row row) {
    Task task = new Task();
    task.taskId = row.getUUID("task_id");
    task.title = row.getString("title");
    task.description = row.getString("description");
    return task;
  }
}
