package com.soa.tour_search.config;

import java.io.IOException;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import io.nats.client.Connection;
import io.nats.client.Nats;

@Configuration
public class NATSConfig {
    
    @Value("${application.nats.host}") private String natsHost;
    @Value("${application.nats.port}") private String natsPort;

    @Bean
    public Connection natsConnection() throws IOException, InterruptedException {
        var url = "nats://" + natsHost + ':' + natsPort;
        System.out.println("NATS JBMUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUU    " + url);
        return Nats.connect("nats://" + natsHost + ':' + natsPort);
    }
    
}
