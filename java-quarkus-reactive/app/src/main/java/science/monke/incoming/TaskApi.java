package science.monke.incoming;

import io.smallrye.mutiny.Multi;
import io.smallrye.mutiny.Uni;
import jakarta.enterprise.inject.Instance;
import jakarta.inject.Inject;
import jakarta.ws.rs.*;
import jakarta.ws.rs.core.MediaType;
import java.util.UUID;
import science.monke.outgoing.TaskRepoPort;

@Path("/tasks")
public class TaskApi {

  private final Instance<TaskRepoPort> taskRepoPort;

  @Inject
  public TaskApi(final Instance<TaskRepoPort> taskRepoPort) {
    this.taskRepoPort = taskRepoPort;
  }

  @GET
  @Produces(MediaType.APPLICATION_JSON)
  public Multi<TaskResponse> getTasks() {
    return taskRepoPort.get().findAll().onItem().transform(TaskResponse::fromTask);
  }

  @POST
  @Consumes(MediaType.APPLICATION_JSON)
  @Produces(MediaType.APPLICATION_JSON)
  public Uni<TaskResponse> postTask(final TaskRequest taskRequest) {
    return taskRepoPort
        .get()
        .save(TaskRequest.toTask(taskRequest))
        .onItem()
        .transform(TaskResponse::fromTask);
  }

  @GET
  @Path("/{taskId}")
  @Produces(MediaType.APPLICATION_JSON)
  public Uni<TaskResponse> getTask(final UUID taskId) {
    return taskRepoPort.get().findByTaskId(taskId).onItem().transform(TaskResponse::fromTask);
  }

  @PUT
  @Path("/{taskId}")
  @Consumes(MediaType.APPLICATION_JSON)
  @Produces(MediaType.APPLICATION_JSON)
  public Uni<TaskResponse> putTask(final UUID taskId, final TaskRequest taskRequest) {
    return taskRepoPort
        .get()
        .update(TaskRequest.toTask(taskId, taskRequest))
        .onItem()
        .transform(TaskResponse::fromTask);
  }

  @DELETE
  @Path("/{taskId}")
  @Produces(MediaType.APPLICATION_JSON)
  public Uni<Void> deleteTask(final UUID taskId) {
    return taskRepoPort.get().delete(taskId);
  }
}
