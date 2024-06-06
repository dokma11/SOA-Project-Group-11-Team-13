package com.soa.tour_search.config;

import java.util.ArrayList;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import io.swagger.v3.oas.models.OpenAPI;
import io.swagger.v3.oas.models.info.Info;
import io.swagger.v3.oas.models.servers.Server;

@Configuration
public class OpenAPIConfig {
    
    @Bean
    public OpenAPI openAPI() {
        ArrayList<Server> servers = new ArrayList<>(3);
        servers.add(new Server().url("http://localhost:8080/").description("development server"));

        Info info = new Info()
                        .title("Tour Search Service")
                        .description("Tour Search API")
                        .version("v1.0");

        OpenAPI openAPI = new OpenAPI().info(info).servers(servers);

        return openAPI;
    }
    
}
