
package insa.application.helpapp.rest;

import org.springframework.data.jpa.repository.JpaRepository;

import java.util.Optional;

public interface ActuatorRepository extends JpaRepository<Actuator, Long> {
    Optional<Actuator> findByRoomAndType(int room, String type);
}