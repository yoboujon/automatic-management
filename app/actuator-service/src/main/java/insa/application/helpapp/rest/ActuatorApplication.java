package insa.application.helpapp.rest;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
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

    @RestController
    public static class ActuatorController {

        @Autowired
        private ActuatorRepository actuatorRepository;

        @GetMapping("/actuators")
        public List<Actuator> getAllActuators() {
            return actuatorRepository.findAll();
        }

        @PutMapping("/actuators/{id}")
        public Actuator updateActuatorState(@PathVariable Long id, @RequestBody StateRequest stateRequest) {
            Actuator actuator = actuatorRepository.findById(id).orElseThrow(() -> new RuntimeException("Actuator not found"));
            actuator.setValue(stateRequest.getState());
            return actuatorRepository.save(actuator);
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
            Actuator[] actuators = restTemplate.getForObject(url, Actuator[].class);

            if (actuators != null) {
                for (Actuator actuator : actuators) {
                    actuatorRepository.save(actuator);
                }
            }

            // Ensure actuator with id 0 is included
            Actuator actuatorWithIdZero = actuatorRepository.findById(0L).orElse(null);
            if (actuatorWithIdZero == null) {
                actuatorWithIdZero = new Actuator(0L, "Heating", "HEATING4000", 0, 1);
                actuatorRepository.save(actuatorWithIdZero);
            }

            // Retrieve all actuators and print to console
            List<Actuator> storedActuators = actuatorRepository.findAll();
            System.out.println("Stored Actuators:");
            storedActuators.forEach(System.out::println);
        }
    }
}

@Entity
class Actuator {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    private String name;
    private String type;
    private Integer svalue;
    private Integer room;

    public Actuator() {
    }

    public Actuator(Long id, String name, String type, Integer svalue, Integer room) {
        this.id = id;
        this.name = name;
        this.type = type;
        this.svalue = svalue;
        this.room = room;
    }

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
        return svalue;
    }

    public void setValue(Integer svalue) {
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
                ", name='" + name + '\'' +
                ", type='" + type + '\'' +
                ", svalue=" + svalue +
                ", room=" + room +
                '}';
    }
}

interface ActuatorRepository extends JpaRepository<Actuator, Long> {
}
