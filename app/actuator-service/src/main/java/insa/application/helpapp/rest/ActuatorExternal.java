package insa.application.helpapp.rest;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.Arrays;
import java.util.List;

@Service
public class ActuatorExternal {

    @Autowired
    private ActuatorRepository actuatorRepository;

    public void fetchAndSaveActuators() {
        // Simulate fetching from an external API
        List<Actuator> actuators = Arrays.asList(
            new Actuator("Heating", "HEATING4000", 0, 1),
            new Actuator("Windows", "AUTOW1", 1, 1),
            new Actuator("Doors", "DOOR2032X", 0, 1)
        );
        actuatorRepository.saveAll(actuators);
    }
}
