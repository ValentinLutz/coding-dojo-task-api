package science.monke.outgoing;

import org.springframework.data.annotation.Id;
import org.springframework.data.relational.core.mapping.Column;
import org.springframework.data.relational.core.mapping.Table;

import java.util.UUID;

@Table(name = "tasks")
public class Task {

  @Id
  @Column(value = "task_id")
  public UUID taskId;

  @Column(value = "title")
  public String title;

  @Column(value = "description")
  public String description;
}
