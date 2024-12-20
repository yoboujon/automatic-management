package insa.application.app.rest;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.cloud.netflix.eureka.EnableEurekaClient;

@SpringBootApplication
@RestController
@EnableEurekaClient
public class DecisionServiceApplication {

    public static void main(String[] args) {
        SpringApplication.run(DecisionServiceApplication.class, args);
    }

    @GetMapping("/hello")
    public String hello() {
        return "Hello from REST!";
    }
}
