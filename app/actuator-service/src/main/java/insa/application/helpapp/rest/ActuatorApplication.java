package insa.application.helpapp.rest;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.client.RestTemplate;

import java.util.Arrays;

@SpringBootApplication
public class ActuatorApplication {

    public static void main(String[] args) {
        SpringApplication.run(ActuatorApplication.class, args);

        // Fetch actuators from external API and print to console
        fetchActuators();
    }

    public static void fetchActuators() {
        // URL of the external API
        String apiUrl = "http://localhost:8085/actuators/";

        // RestTemplate to make HTTP requests
        RestTemplate restTemplate = new RestTemplate();

        try {
            // Fetch actuators as an array
            Actuator[] actuators = restTemplate.getForObject(apiUrl, Actuator[].class);

            if (actuators != null) {
                // Print each actuator to the console
                Arrays.stream(actuators).forEach(actuator -> {
                    System.out.println("Name: " + actuator.getName());
                    System.out.println("Type: " + actuator.getType());
                    System.out.println("ID: " + actuator.getId());
                    System.out.println("Value: " + actuator.getValue());
                    System.out.println("Room: " + actuator.getRoom());
                    System.out.println("--------------------------------");
                });
            } else {
                System.out.println("No actuators found.");
            }
        } catch (Exception e) {
            System.err.println("Error fetching actuators: " + e.getMessage());
        }
    }
}
