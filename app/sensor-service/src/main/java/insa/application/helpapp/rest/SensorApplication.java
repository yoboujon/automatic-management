package insa.application.helpapp.rest;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.*;

import java.time.LocalDateTime;
import java.util.List;

@SpringBootApplication
@RestController
@RequestMapping("/sensors")
public class SensorApplication {

    @Autowired
    private SensorRepository sensorRepository;

    public static void main(String[] args) {
        SpringApplication.run(SensorApplication.class, args);
    }

    @PostMapping
    public Sensor createSensor(@RequestBody Sensor sensor) {
        sensor.setTimestamp(LocalDateTime.now());
        return sensorRepository.save(sensor);
    }

    @GetMapping
    public List<Sensor> getAllSensors() {
        return sensorRepository.findAll();
    }

    @GetMapping("/{id}")
    public Sensor getSensorById(@PathVariable Long id) {
        return sensorRepository.findById(id).orElseThrow(() -> new RuntimeException("Sensor not found"));
    }

    @DeleteMapping("/{id}")
    public void deleteSensor(@PathVariable Long id) {
        sensorRepository.deleteById(id);
    }
}
