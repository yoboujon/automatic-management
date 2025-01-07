## Installation

```
mvn compile 
```

Une fois le projet compilé, lancer les différents services SOAP et REST via `Spring Boot` :

```
mvn spring-boot:run -pl rest-service
```


```
mvn spring-boot:run -pl soap-service
```

Vérifier le bon fonctionnement en accédant à `localhost:8080/hello` pour REST et `localhost:8081/ws` pour SOAP.
