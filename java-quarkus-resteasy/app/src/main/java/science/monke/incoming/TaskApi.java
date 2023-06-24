package science.monke.incoming;

import jakarta.enterprise.inject.Instance;
import jakarta.inject.Inject;
import jakarta.ws.rs.*;
import jakarta.ws.rs.core.MediaType;
import jakarta.ws.rs.core.Response;
import java.util.UUID;
import java.util.stream.Collectors;
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
  public Response getTasks() {
    return Response.ok(
            taskRepoPort.get().findAll().stream()
                .map(TaskResponse::fromTask)
                .collect(Collectors.toList()))
        .build();
  }

  @POST
  @Consumes(MediaType.APPLICATION_JSON)
  @Produces(MediaType.APPLICATION_JSON)
  public Response postTask(final TaskRequest taskRequest) {
    return Response.status(Response.Status.CREATED).entity(
            TaskResponse.fromTask(taskRepoPort.get().save(TaskRequest.toTask(taskRequest))))
        .build();
  }

  @GET
  @Path("/{taskId}")
  @Produces(MediaType.APPLICATION_JSON)
  public Response getTask(@PathParam("taskId") final UUID taskId) {
    return taskRepoPort
        .get()
        .findByTaskId(taskId)
        .map(task -> Response.ok(TaskResponse.fromTask(task)).build())
        .orElse(Response.status(Response.Status.NOT_FOUND).build());
  }

  @PUT
  @Path("/{taskId}")
  @Consumes(MediaType.APPLICATION_JSON)
  @Produces(MediaType.APPLICATION_JSON)
  public Response putTask(@PathParam("taskId") final UUID taskId, final TaskRequest taskRequest) {
    return taskRepoPort
        .get()
        .update(TaskRequest.toTask(taskId, taskRequest))
        .map(task -> Response.noContent().build())
        .orElse(Response.status(Response.Status.NOT_FOUND).build());
  }

  @DELETE
  @Path("/{taskId}")
  @Produces(MediaType.APPLICATION_JSON)
  public Response deleteTask(@PathParam("taskId") final UUID taskId) {
    return taskRepoPort.get().delete(taskId)
        ? Response.noContent().build()
        : Response.status(Response.Status.NOT_FOUND).build();
  }
}
