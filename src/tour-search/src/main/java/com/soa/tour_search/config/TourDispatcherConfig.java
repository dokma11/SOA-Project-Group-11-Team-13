package com.soa.tour_search.config;

import java.nio.charset.StandardCharsets;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.soa.tour_search.model.Tour;
import com.soa.tour_search.service.impl.TourService;

import io.nats.client.Connection;
import io.nats.client.Dispatcher;
import lombok.RequiredArgsConstructor;

@Configuration
@RequiredArgsConstructor
public class TourDispatcherConfig {
    
    private final Connection natsConnection;

    private final TourService tourService;

    @Bean
    public Dispatcher tourDispatcher() {
        System.out.println("DISPECER JBMUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUU");
        var dispatcher = natsConnection.createDispatcher(msg -> {
            var data = new String(msg.getData(), StandardCharsets.UTF_8);
            data = data.replace("Id", "id");
            data = data.replace("Name", "name");
            data = data.replace("Description", "description");
            ObjectMapper objectMapper = new ObjectMapper();
            try {
                Tour tour = objectMapper.readValue(data, Tour.class);
                tourService.save(tour);
            } catch (Exception e) {
                e.printStackTrace();
            }
        });
        dispatcher.subscribe("com.tours");
        return dispatcher;
    }

}
