package science.monke.outgoing;


import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.UUID;

public class Task {
  public UUID taskId;
  public String title;
  public String description;

  public static Task fromResultSet(final ResultSet resultSet) throws SQLException {
    Task task = new Task();
    task.taskId = UUID.fromString(resultSet.getString("task_id"));
    task.title = resultSet.getString("title");
    task.description = resultSet.getString("description");
    return task;
  }
}
