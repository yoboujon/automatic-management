package insa.application.helpapp.rest;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpMethod;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;

import java.util.List;

@Service
public class ActuatorService {

    @Autowired
    private ActuatorRepository actuatorRepository;

    private final String EXTERNAL_API_URL = "http://localhost:8085/actuators/";

    public List<Actuator> getAllActuators() {
        // Fetch actuators from the database
        return actuatorRepository.findAll();
    }

    public Actuator updateActuatorState(Long id, Integer state) {
        // Create a RestTemplate object
        RestTemplate restTemplate = new RestTemplate();

        // Create the request body
        ActuatorStateRequest stateRequest = new ActuatorStateRequest(state);

        // Prepare the HTTP entity with the request body
        HttpEntity<ActuatorStateRequest> requestEntity = new HttpEntity<>(stateRequest);

        // Send the PUT request using RestTemplate.exchange
        ResponseEntity<Actuator> response = restTemplate.exchange(
                EXTERNAL_API_URL + id,
                HttpMethod.PUT,
                requestEntity,
                Actuator.class
        );

        // Get the response body
        Actuator updatedActuator = response.getBody();

        // Update actuator in the database
        if (updatedActuator != null) {
            actuatorRepository.save(updatedActuator);
        }

        return updatedActuator;
    }

    public void fetchAndStoreActuators() {
        RestTemplate restTemplate = new RestTemplate();
        Actuator[] actuators = restTemplate.getForObject(EXTERNAL_API_URL, Actuator[].class);

        if (actuators != null) {
            actuatorRepository.saveAll(List.of(actuators));
        }
    }
}
