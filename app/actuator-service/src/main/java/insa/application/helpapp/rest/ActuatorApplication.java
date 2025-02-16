package insa.application.helpapp.rest;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Component;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.RestTemplate;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;

import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;

import java.util.List;

@SpringBootApplication
public class ActuatorApplication {

    public static void main(String[] args) {
        SpringApplication.run(ActuatorApplication.class, args);
    }

    @Bean
    public RestTemplate restTemplate() {
        return new RestTemplate();
    }

    @RestController
    public static class ActuatorController {

        @Autowired
        private ActuatorRepository actuatorRepository;

        @Autowired
        private RestTemplate restTemplate;

        @GetMapping("/actuators")
        public List<Actuator> getAllActuators() {
            return actuatorRepository.findAll();
        }

        @GetMapping("/actuators/{id}")
        public Actuator getActuatorById(@PathVariable Long id) {
            return actuatorRepository.findById(id)
                .orElseThrow(() -> new RuntimeException("Actuator not found"));
        }

        @PutMapping("/actuators/{id}")
        public Actuator updateActuatorState(@PathVariable Long id, @RequestBody StateRequest stateRequest) {
            Actuator actuator = actuatorRepository.findById(id)
                .orElseThrow(() -> new RuntimeException("Actuator not found"));
            actuator.setSvalue(stateRequest.getState());
            actuatorRepository.save(actuator);

            // Send PUT request to external API
            String url = "http://localhost:8085/actuators/" + actuator.getExternalId();
            restTemplate.put(url, stateRequest);

            return actuator;
        }
    }

    public static class StateRequest {
        private Integer state;

        public Integer getState() {
            return state;
        }

        public void setState(Integer state) {
            this.state = state;
        }
    }

    @Component
    public static class DataLoader implements CommandLineRunner {

        @Autowired
        private ActuatorRepository actuatorRepository;

        @Override
        public void run(String... args) {
            RestTemplate restTemplate = new RestTemplate();
            String url = "http://localhost:8085/actuators/";
            ExternalActuator[] externalActuators = restTemplate.getForObject(url, ExternalActuator[].class);

            if (externalActuators != null) {
                for (ExternalActuator externalActuator : externalActuators) {
                    Actuator actuator = actuatorRepository.findByExternalId(externalActuator.getId());
                    if (actuator == null) { // Avoid duplicates
                        actuator = new Actuator(
                            externalActuator.getId(), // External ID
                            externalActuator.getName(),
                            externalActuator.getType(),
                            externalActuator.getValue(),
                            externalActuator.getRoom()
                        );
                        actuatorRepository.save(actuator);
                    }
                }
            }

            // Print stored actuators
            List<Actuator> storedActuators = actuatorRepository.findAll();
            System.out.println("Stored Actuators:");
            storedActuators.forEach(System.out::println);
        }
    }
}

@Entity
class Actuator {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY) // Database-generated ID
    private Long id;

    private Long externalId; // ID from external API
    private String name;
    private String type;
    private Integer svalue;
    private Integer room;

    public Actuator() {
    }

    public Actuator(Long externalId, String name, String type, Integer svalue, Integer room) {
        this.externalId = externalId;
        this.name = name;
        this.type = type;
        this.svalue = svalue;
        this.room = room;
    }

    public Long getId() {
        return id;
    }

    public Long getExternalId() {
        return externalId;
    }

    public void setExternalId(Long externalId) {
        this.externalId = externalId;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getType() {
        return type;
    }

    public void setType(String type) {
        this.type = type;
    }

    public Integer getSvalue() {
        return svalue;
    }

    public void setSvalue(Integer svalue) {
        this.svalue = svalue;
    }

    public Integer getRoom() {
        return room;
    }

    public void setRoom(Integer room) {
        this.room = room;
    }

    @Override
    public String toString() {
        return "Actuator{" +
                "id=" + id +
                ", externalId=" + externalId +
                ", name='" + name + '\'' +
                ", type='" + type + '\'' +
                ", svalue=" + svalue +
                ", room=" + room +
                '}';
    }
}

class ExternalActuator {
    private Long id; // External API ID
    private String name;
    private String type;
    private Integer value;
    private Integer room;

    // Getters and setters
    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getType() {
        return type;
    }

    public void setType(String type) {
        this.type = type;
    }

    public Integer getValue() {
        return value;
    }

    public void setValue(Integer value) {
        this.value = value;
    }

    public Integer getRoom() {
        return room;
    }

    public void setRoom(Integer room) {
        this.room = room;
    }
}

interface ActuatorRepository extends JpaRepository<Actuator, Long> {
    Actuator findByExternalId(Long externalId); // Custom method to find by external ID
}
