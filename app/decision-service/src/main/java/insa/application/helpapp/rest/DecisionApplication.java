package insa.application.helpapp.rest;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import java.time.LocalDateTime;
import java.util.List;

@SpringBootApplication
@RestController
public class DecisionApplication implements CommandLineRunner {

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

    // Main loop for the decision application
    @Override
    public void run(String... args) {
        decisionRepository.save(new Decision("Temperature > 25°C", "Turn on cooling", LocalDateTime.now()));
        decisionRepository.save(new Decision("Humidity < 30%", "Turn on humidifier", LocalDateTime.now()));
        decisionRepository.save(new Decision("No presence detected", "Turn off lights", LocalDateTime.now()));
    }
}
