package science.monke.outgoing;

import com.zaxxer.hikari.HikariConfig;
import com.zaxxer.hikari.HikariDataSource;

import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.SQLException;
import java.sql.Statement;
import java.util.ArrayList;
import java.util.List;
import java.util.Optional;
import java.util.UUID;

public class TaskRepoPostgres implements TaskRepoPort {

  private final HikariDataSource dataSource;

  public TaskRepoPostgres(final HikariDataSource dataSource) {
    this.dataSource = dataSource;
  }

  @Override
  public List<Task> findAll() {
    try (Connection connection = dataSource.getConnection()) {
      try (Statement statement = connection.createStatement()) {
        var resultSet = statement.executeQuery("SELECT task_id, title, description FROM public.tasks;");

        var tasks = new ArrayList<Task>();
        while (resultSet.next()) {
          tasks.add(Task.fromResultSet(resultSet));
        }

        return tasks;
      }
    } catch (SQLException exception) {
      throw new RuntimeException(exception);
    }
  }

  @Override
  public Optional<Task> findByTaskId(final UUID taskId) {
    try (Connection connection = dataSource.getConnection()) {
      try (PreparedStatement statement = connection.prepareStatement("SELECT task_id, title, description FROM public.tasks WHERE task_id = ?;")) {
        statement.setObject(1, taskId);
        var resultSet = statement.executeQuery();

        if (resultSet.next()) {
          return Optional.of(Task.fromResultSet(resultSet));
        } else {
          return Optional.empty();
        }
      }
    } catch (SQLException exception) {
      throw new RuntimeException(exception);
    }
  }

  @Override
  public Task save(final Task task) {
    try (Connection connection = dataSource.getConnection()) {
      try (PreparedStatement statement = connection.prepareStatement("INSERT INTO public.tasks (task_id, title, description) VALUES (?, ?, ?);")) {
        statement.setObject(1, task.taskId);
        statement.setString(2, task.title);
        statement.setString(3, task.description);
        statement.executeUpdate();
        return task;
      }
    } catch (SQLException exception) {
      throw new RuntimeException(exception);
    }
  }

  @Override
  public Task update(Task task) {
    try (Connection connection = dataSource.getConnection()) {
      try (PreparedStatement statement =
          connection.prepareStatement(
              "UPDATE public.tasks SET title = ?, description = ? WHERE task_id = ?;")) {
        statement.setString(1, task.title);
        statement.setString(2, task.description);
        statement.setObject(3, task.taskId);
        int rowsUpdated = statement.executeUpdate();
        if (rowsUpdated == 0) {
          return save(task);
        }
        return task;
      }
    } catch (SQLException exception) {
      throw new RuntimeException(exception);
    }
  }


  @Override
  public void deleteById(final UUID taskId) {
    try (Connection connection = dataSource.getConnection()) {
      try (PreparedStatement statement = connection.prepareStatement("DELETE FROM public.tasks WHERE task_id = ?;")) {
        statement.setObject(1, taskId);
        statement.executeUpdate();
      }
    } catch (SQLException exception) {
      throw new RuntimeException(exception);
    }
  }
}