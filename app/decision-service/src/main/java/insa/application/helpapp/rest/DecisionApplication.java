package insa.application.helpapp.rest;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
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

    private float getRequest(String url) {
        RestTemplate restTemplate = new RestTemplate();
        SensorResponse response = restTemplate.getForObject(url, SensorResponse.class);
        return response != null ? response.getValue() : 0;
    }
    

    // Main loop for the decision application
    @Scheduled(fixedRate = 200)
    public void decisionLoop() {
        float temperature = getRequest("http://localhost:8085/sensors/1");
        float ppm = getRequest("http://localhost:8085/sensors/0");

        if(temperature < 20) {
            // startHeating()
            DecisionList d = new DecisionList(ActionEnum.HEATING_START, SensorEnum.TEMPERATURE);
            decisionRepository.save(new Decision(d));
            // closeWindow()
            d = new DecisionList(ActionEnum.HEATING_START, SensorEnum.TEMPERATURE);
            decisionRepository.save(new Decision(d));
        }

        if(temperature > 22) {
            // stopHeating()
            DecisionList d = new DecisionList(ActionEnum.HEATING_STOP, SensorEnum.TEMPERATURE);
            decisionRepository.save(new Decision(d));
        }

        if(temperature > 25) {
            // openWindow()
            DecisionList d = new DecisionList(ActionEnum.WINDOWS_OPEN, SensorEnum.TEMPERATURE);
            decisionRepository.save(new Decision(d));
        }

        if(ppm >= 800) {
            // openWindow()
            DecisionList d = new DecisionList(ActionEnum.WINDOWS_OPEN, SensorEnum.CARBON_DIOXIDE);
            decisionRepository.save(new Decision(d));
        }
    }
}
