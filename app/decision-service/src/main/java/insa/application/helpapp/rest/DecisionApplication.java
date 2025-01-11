package insa.application.helpapp.rest;

import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.time.LocalDateTime;
import java.util.List;

@SpringBootApplication
@RestController
@RequestMapping("/decisions")
public class DecisionApplication implements CommandLineRunner {

    public static void main(String[] args) {
        SpringApplication.run(DecisionApplication.class, args);
    }

    @Autowired
    private DecisionLogRepository decisionLogRepository;

    // Endpoint pour tester si le service fonctionne
    @GetMapping("/logs")
    public List<DecisionLog> getAllLogs() {
        return decisionLogRepository.findAll();
    }

    // Ajouter des données fictives lors du démarrage
    @Override
    public void run(String... args) {
        decisionLogRepository.save(new DecisionLog("Temperature > 25°C", "Turn on cooling", LocalDateTime.now()));
        decisionLogRepository.save(new DecisionLog("Humidity < 30%", "Turn on humidifier", LocalDateTime.now()));
        decisionLogRepository.save(new DecisionLog("No presence detected", "Turn off lights", LocalDateTime.now()));
    }
}

@Entity
class DecisionLog {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    private String condition;
    private String action;
    private LocalDateTime timestamp;

    // Constructeurs, getters et setters
    public DecisionLog() {}

    public DecisionLog(String condition, String action, LocalDateTime timestamp) {
        this.condition = condition;
        this.action = action;
        this.timestamp = timestamp;
    }

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getCondition() {
        return condition;
    }

    public void setCondition(String condition) {
        this.condition = condition;
    }

    public String getAction() {
        return action;
    }

    public void setAction(String action) {
        this.action = action;
    }

    public LocalDateTime getTimestamp() {
        return timestamp;
    }

    public void setTimestamp(LocalDateTime timestamp) {
        this.timestamp = timestamp;
    }
}

@Repository
interface DecisionLogRepository extends JpaRepository<DecisionLog, Long> {}
