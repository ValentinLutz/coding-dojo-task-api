package science.monke.outgoing;

import java.util.UUID;

import org.springframework.boot.autoconfigure.condition.ConditionalOnProperty;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

@Repository
@ConditionalOnProperty(name = "app.memory.enabled", havingValue = "false")
public interface TaskRepoPostgres extends CrudRepository<Task, UUID>, TaskRepoPort {}
