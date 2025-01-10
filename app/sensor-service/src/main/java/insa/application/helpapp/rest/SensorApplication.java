package insa.application.helpapp.rest;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.RestTemplate;

@SpringBootApplication
@RestController
public class SensorApplication {

    public static void main(String[] args) {
        SpringApplication.run(SensorApplication.class, args);
    }

    @GetMapping("/hello")
    public String hello() {
        return "Hello from SensorApplication!";
    }

    @GetMapping("/sensors")
    public String getSensors() {
        RestTemplate restTemplate = new RestTemplate();
        String sensorsData = restTemplate.getForObject("http://localhost:8085/sensors", String.class);
        return sensorsData;
    }
}
