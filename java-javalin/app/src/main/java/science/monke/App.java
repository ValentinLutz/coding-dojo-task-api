package science.monke;

import com.zaxxer.hikari.HikariConfig;
import com.zaxxer.hikari.HikariDataSource;
import io.javalin.Javalin;
import org.jetbrains.annotations.NotNull;
import science.monke.incoming.TaskApi;
import science.monke.outgoing.TaskRepoMemory;
import science.monke.outgoing.TaskRepoPort;
import science.monke.outgoing.TaskRepoPostgres;

import java.sql.SQLException;
import java.sql.Statement;

public class App {
  public static void main(String[] args) {
    final int port = Integer.parseInt(System.getenv("PORT"));
    final boolean useInMemory = Boolean.parseBoolean(System.getenv("USE_IN_MEMORY"));

    TaskRepoPort taskRepo;
    if (useInMemory) {
      taskRepo = new TaskRepoMemory();
    } else {
      HikariDataSource dataSource = createDatabase();
      initDatabase(dataSource);
      taskRepo = new TaskRepoPostgres(dataSource);
    }

    final var app = Javalin.create();

    TaskApi taskApi = new TaskApi(taskRepo);
    taskApi.registerRoutes(app);

    app.start(port);
  }

  private static HikariDataSource createDatabase() {
    HikariConfig config = new HikariConfig();
    config.setJdbcUrl(
        String.format(
            "jdbc:postgresql://%s:%s/%s",
            System.getenv("POSTGRES_HOST"),
            System.getenv("POSTGRES_PORT"),
            System.getenv("POSTGRES_DATABASE")));
    config.setUsername(System.getenv("POSTGRES_USER"));
    config.setPassword(System.getenv("POSTGRES_PASSWORD"));
    return new HikariDataSource(config);
  }

  private static void initDatabase(@NotNull final HikariDataSource dataSource) {
    try (var connection = dataSource.getConnection()) {

      try (Statement statement = connection.createStatement()) {
        statement.execute("DROP TABLE IF EXISTS public.tasks;");
        statement.execute(
            "CREATE TABLE IF NOT EXISTS public.tasks (" +
                    "task_id UUID PRIMARY KEY NOT NULL," +
                    "title TEXT," +
                    "description TEXT" +
                    ");");
      }
    } catch (SQLException exception) {
      throw new RuntimeException(exception);
    }
  }
}