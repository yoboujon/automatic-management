package insa.application.helpapp.rest;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.client.RestTemplate;

import java.util.List;

@SpringBootApplication
@RestController
@RequestMapping("/actuators")
public class ActuatorApplication {

    @Autowired
    private ActuatorService actuatorService;

    public static void main(String[] args) {
        SpringApplication.run(ActuatorApplication.class, args);
    }

    @GetMapping
    public List<Actuator> getActuators() {
        return actuatorService.getAllActuators();
    }

    @PutMapping("/{id}")
    public Actuator updateActuatorState(@PathVariable Long id, @RequestBody ActuatorStateRequest stateRequest) {
        return actuatorService.updateActuatorState(id, stateRequest.getState());
    }
}
