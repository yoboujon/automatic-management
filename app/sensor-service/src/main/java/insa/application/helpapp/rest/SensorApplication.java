package insa.application.helpapp.rest;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.*;

import java.time.LocalDateTime;
import java.util.List;

@SpringBootApplication
@RestController
@RequestMapping("/")
public class SensorApplication {

    @Autowired
    private SensorRepository sensorRepository;

    public static void main(String[] args) {
        SpringApplication.run(SensorApplication.class, args);
    }

    @GetMapping("/test")
    public String test() {
        return "test";
    }

    @PostMapping("/init-db")
    public String initDatabase() {
        Sensor sensor = new Sensor("temperature", 22.5, LocalDateTime.now());
        sensorRepository.save(sensor);
        return "Sensor saved: " + sensor;
    }

    @GetMapping("/sensors")
    public List<Sensor> getAllSensors() {
        return sensorRepository.findAll();
    }

    @PostMapping("/test-repo")
    public String testRepository() {
    Sensor testSensor = new Sensor("test", 0.0, LocalDateTime.now());
    sensorRepository.save(testSensor);
    return "Repository is working!";
    }

}
