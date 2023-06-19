package science.monke.incoming;

import io.javalin.Javalin;
import io.javalin.http.Context;
import java.util.List;
import java.util.UUID;
import org.jetbrains.annotations.NotNull;
import science.monke.outgoing.TaskRepoPort;

public class TaskApi {

  private final TaskRepoPort taskRepoPort;

  public TaskApi(final TaskRepoPort taskRepoPort) {
    this.taskRepoPort = taskRepoPort;
  }

  public Javalin registerRoutes(Javalin app) {
    app.get("/tasks", this::getTasks);
    app.post("/tasks", this::postTasks);
    app.get("/tasks/{task_id}", this::getTask);
    app.put("/tasks/{task_id}", this::putTask);
    app.delete("/tasks/{task_id}", this::deleteTask);
    return app;
  }

  public void getTasks(@NotNull Context ctx) {
    List<TaskResponse> taskResponses =
        taskRepoPort.findAll().stream().map(TaskResponse::fromTask).toList();

    ctx.json(taskResponses).status(200);
  }

  public void postTasks(@NotNull Context ctx) {
    TaskRequest taskRequest = ctx.bodyAsClass(TaskRequest.class);

    TaskResponse taskResponse =
        TaskResponse.fromTask(taskRepoPort.save(TaskRequest.toTask(taskRequest)));

    ctx.json(taskResponse).status(201);
  }

  public void getTask(@NotNull Context ctx) {
    UUID taskId = UUID.fromString(ctx.pathParam("task_id"));

    taskRepoPort
        .findByTaskId(taskId)
        .map(TaskResponse::fromTask)
        .ifPresentOrElse(taskResponse -> ctx.json(taskResponse).status(200), () -> ctx.status(404));
  }

  public void putTask(@NotNull Context ctx) {
    UUID taskId = UUID.fromString(ctx.pathParam("task_id"));
    TaskRequest taskRequest = ctx.bodyAsClass(TaskRequest.class);

    TaskResponse taskResponse =
        TaskResponse.fromTask(taskRepoPort.update(TaskRequest.toTask(taskId, taskRequest)));

    ctx.json(taskResponse).status(204);
  }

  public void deleteTask(@NotNull Context ctx) {
    UUID taskId = UUID.fromString(ctx.pathParam("task_id"));

    taskRepoPort.deleteById(taskId);

    ctx.status(204);
  }
}
