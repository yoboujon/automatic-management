package insa.application.helpapp.rest;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/actuators")
public class ActuatorController {

    @Autowired
    private ActuatorRepository actuatorRepository;

    @GetMapping("/")
    public List<Actuator> getAllActuators() {
        return actuatorRepository.findAll();
    }

    @PutMapping("/{id}")
    public Actuator updateActuatorState(@PathVariable Long id, @RequestBody Actuator actuatorRequest) {
        return actuatorRepository.findById(id).map(actuator -> {
            actuator.setValue(actuatorRequest.getValue());
            return actuatorRepository.save(actuator);
        }).orElseThrow(() -> new RuntimeException("Actuator not found with id " + id));
    }
}
