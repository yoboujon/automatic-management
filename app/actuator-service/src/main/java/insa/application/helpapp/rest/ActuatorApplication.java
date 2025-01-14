package insa.application.helpapp.rest;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@SpringBootApplication
@RestController
@RequestMapping("/actuators")
public class ActuatorApplication {

    private final ActuatorService actuatorService;

    public ActuatorApplication(ActuatorService actuatorService) {
        this.actuatorService = actuatorService;
    }

    public static void main(String[] args) {
        SpringApplication.run(ActuatorApplication.class, args);
    }

    @GetMapping
    public List<Actuator> getAllActuators() {
        return actuatorService.getAllActuators();
    }

    @PutMapping("/{room}/{type}")
    public Actuator updateActuatorState(
            @PathVariable int room,
            @PathVariable String type,
            @RequestBody ActuatorRequest actuatorRequest) {
        return actuatorService.updateActuatorState(room, type, actuatorRequest.getState());
    }
}
