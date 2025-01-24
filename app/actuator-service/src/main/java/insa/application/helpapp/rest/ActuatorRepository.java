package insa.application.helpapp.rest;

import org.springframework.data.jpa.repository.JpaRepository;

public interface ActuatorRepository extends JpaRepository<Actuator, Long> {
}
