package science.monke.outgoing;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.Id;
import jakarta.persistence.Table;

import java.util.UUID;

@Entity
@Table(name = "tasks")
public class Task {

  @Id
  @Column(name = "task_id")
  public UUID taskId;

  @Column(name = "title")
  public String title;

  @Column(name = "description")
  public String description;
}
