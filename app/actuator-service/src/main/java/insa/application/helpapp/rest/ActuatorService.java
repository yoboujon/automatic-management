package insa.application.helpapp.rest;

import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class ActuatorService {

    private final ActuatorRepository actuatorRepository;

    public ActuatorService(ActuatorRepository actuatorRepository) {
        this.actuatorRepository = actuatorRepository;
    }

    public List<Actuator> getAllActuators() {
        return actuatorRepository.findAll();
    }

    public Actuator updateActuatorState(int room, String type, int state) {
        Actuator actuator = actuatorRepository.findByRoomAndType(room, type)
                .orElseThrow(() -> new RuntimeException("Actuator not found for room " + room + " and type " + type));
        actuator.setValue(state);
        return actuatorRepository.save(actuator);
    }
}
