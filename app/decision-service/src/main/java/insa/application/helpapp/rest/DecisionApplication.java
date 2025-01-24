package insa.application.helpapp.rest;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;
import org.springframework.scheduling.annotation.EnableScheduling;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.RestTemplate;

import insa.application.helpapp.rest.DecisionList.ActionEnum;
import insa.application.helpapp.rest.DecisionList.SensorEnum;

import java.util.List;

@SpringBootApplication
@RestController
@EnableScheduling
public class DecisionApplication {
    private static final String SENSOR_PATH = "http://localhost:8085";
    private static final String ACTUATOR_PATH = "http://localhost:8085";

    public static void main(String[] args) {
        SpringApplication.run(DecisionApplication.class, args);
    }

    @Autowired
    private DecisionRepository decisionRepository;

    // Logs endpoint
    @GetMapping("/logs")
    public List<Decision> getAllLogs() {
        return decisionRepository.findAll();
    }

    // Bean to do Rest requests
    @Bean
    public RestTemplate restTemplate() {
        return new RestTemplate();
    }

    private float getSensor(int id) {
        RestTemplate restTemplate = new RestTemplate();
        String url = SENSOR_PATH + "/sensors/" + String.valueOf(id);
        SensorResponse response = restTemplate.getForObject(url, SensorResponse.class);
        return response != null ? response.getValue() : 0;
    }

    private int getActuator(int id) {
        RestTemplate restTemplate = new RestTemplate();
        String url = ACTUATOR_PATH + "/actuators/" + String.valueOf(id);
        ActuatorResponse response = restTemplate.getForObject(url, ActuatorResponse.class);
        return response != null ? response.getValue() : 0;
    }

    private void putActuator(int id, int state) {
        RestTemplate restTemplate = new RestTemplate();
        String url = ACTUATOR_PATH + "/actuators/" + String.valueOf(id);
        ActuatorRequest request = new ActuatorRequest(state);
        restTemplate.put(url, request);
    }

    // Main loop for the decision application
    @Scheduled(fixedRate = 200)
    public void decisionLoop() {
        float temperature = getSensor(1);
        float ppm = getSensor(0);
        int heatingValue = getActuator(0);
        int windowStatus = getActuator(1);
        // System.out.println("window: "+windowStatus+"\nheating: "+heatingValue+"\nppm:
        // "+ppm+"\ntemp: "+temperature+"\n");

        // temperature too low has a higher priority
        if (temperature < 20) {
            if ((heatingValue != 22) || (windowStatus != 0)) {
                System.out.println("Temperature too low");
            }

            if (heatingValue != 22) {
                // Start heating
                putActuator(0, 22);
                DecisionList d = new DecisionList(ActionEnum.HEATING_START, SensorEnum.TEMPERATURE);
                decisionRepository.save(new Decision(d));
            }
            if (windowStatus != 0) {
                // Closing window
                putActuator(1, 0);
                DecisionList d = new DecisionList(ActionEnum.WINDOWS_CLOSE, SensorEnum.TEMPERATURE);
                decisionRepository.save(new Decision(d));
            }
            return;
        }

        if (ppm >= 800 && windowStatus != 1) {
            System.out.println("PPM too high");
            // Opening window
            putActuator(1, 1);
            DecisionList d = new DecisionList(ActionEnum.WINDOWS_OPEN, SensorEnum.CARBON_DIOXIDE);
            decisionRepository.save(new Decision(d));
        }

        if ((temperature > 25) && (windowStatus != 1)) {
            System.out.println("Temperature too high");
            // Opening window
            putActuator(1, 1);
            DecisionList d = new DecisionList(ActionEnum.WINDOWS_OPEN, SensorEnum.TEMPERATURE);
            decisionRepository.save(new Decision(d));
            // Stop heating
            putActuator(0, 0);
            d = new DecisionList(ActionEnum.HEATING_STOP, SensorEnum.TEMPERATURE);
            decisionRepository.save(new Decision(d));
        }
    }
}
