package science.monke.outgoing;

import io.agroal.api.AgroalDataSource;
import io.quarkus.arc.lookup.LookupIfProperty;
import io.quarkus.runtime.StartupEvent;
import jakarta.enterprise.context.ApplicationScoped;
import jakarta.enterprise.event.Observes;
import jakarta.inject.Inject;
import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.ArrayList;
import java.util.List;
import java.util.Optional;
import java.util.UUID;
import org.jboss.logging.Logger;

@ApplicationScoped
@LookupIfProperty(name = "app.memory.enabled", stringValue = "false")
public class TaskRepoPostgres implements TaskRepoPort {

  private static final Logger LOGGER = Logger.getLogger(TaskRepoPostgres.class);

  private final AgroalDataSource dataSource;

  @Inject
  public TaskRepoPostgres(final AgroalDataSource dataSource) {
    this.dataSource = dataSource;
  }

  void config(@Observes final StartupEvent startupEvent) throws SQLException {
    try (final Connection connection = dataSource.getConnection()) {
      try (final PreparedStatement preparedStatement =
          connection.prepareStatement("DROP TABLE IF EXISTS public.tasks")) {
        preparedStatement.execute();
        LOGGER.info("Dropped table tasks");
      }
      try (final PreparedStatement preparedStatement =
          connection.prepareStatement(
              "CREATE TABLE IF NOT EXISTS public.tasks (task_id UUID PRIMARY KEY NOT NULL, title TEXT, description TEXT)")) {
        preparedStatement.execute();
        LOGGER.info("Created table tasks");
      }
    }
  }

  @Override
  public List<Task> findAll() {
    try (final Connection connection = dataSource.getConnection()) {
      try (final PreparedStatement preparedStatement =
          connection.prepareStatement("SELECT task_id, title, description FROM public.tasks")) {
        final ResultSet resultSet = preparedStatement.executeQuery();
        final List<Task> tasks = new ArrayList<>();
        while (resultSet.next()) {
          tasks.add(Task.fromResultSet(resultSet));
        }
        return tasks;
      }
    } catch (SQLException sqlException) {
      throw new RuntimeException(sqlException);
    }
  }

  @Override
  public Optional<Task> findByTaskId(final UUID taskId) {
    try (final Connection connection = dataSource.getConnection()) {
      try (final PreparedStatement preparedStatement =
          connection.prepareStatement(
              "SELECT task_id, title, description FROM public.tasks WHERE task_id = ?")) {
        preparedStatement.setObject(1, taskId);
        final ResultSet resultSet = preparedStatement.executeQuery();
        if (!resultSet.next()) {
          return Optional.empty();
        }
        return Optional.of(Task.fromResultSet(resultSet));
      }
    } catch (SQLException sqlException) {
      throw new RuntimeException(sqlException);
    }
  }

  @Override
  public Task save(final Task task) {
    try (final Connection connection = dataSource.getConnection()) {
      try (final PreparedStatement preparedStatement =
          connection.prepareStatement(
              "INSERT INTO public.tasks (task_id, title, description) VALUES (?, ?, ?)")) {
        preparedStatement.setObject(1, task.taskId);
        preparedStatement.setString(2, task.title);
        preparedStatement.setString(3, task.description);
        preparedStatement.execute();
        return task;
      }
    } catch (SQLException sqlException) {
      throw new RuntimeException(sqlException);
    }
  }

  @Override
  public Optional<Task> update(final Task task) {
    try (final Connection connection = dataSource.getConnection()) {
      try (final PreparedStatement preparedStatement =
          connection.prepareStatement(
              "UPDATE public.tasks SET title = ?, description = ? WHERE task_id = ?")) {
        preparedStatement.setString(1, task.title);
        preparedStatement.setString(2, task.description);
        preparedStatement.setObject(3, task.taskId);
        final int updatedRows = preparedStatement.executeUpdate();
        if (updatedRows == 0) {
          return Optional.empty();
        }
        return Optional.of(task);
      }
    } catch (SQLException sqlException) {
      throw new RuntimeException(sqlException);
    }
  }

  @Override
  public boolean delete(final UUID taskId) {
    try (final Connection connection = dataSource.getConnection()) {
      try (final PreparedStatement preparedStatement =
          connection.prepareStatement("DELETE FROM public.tasks WHERE task_id = ?")) {
        preparedStatement.setObject(1, taskId);
        final int deletedRows = preparedStatement.executeUpdate();
        return deletedRows != 0;
      }
    } catch (SQLException sqlException) {
      throw new RuntimeException(sqlException);
    }
  }
}
