package com.soa.tour_search.model;

import org.springframework.data.annotation.Id;
import org.springframework.data.elasticsearch.annotations.Document;

import lombok.Getter;
import lombok.Setter;

@Document(indexName = "exhibition")
@Getter @Setter
public class Tour {

    @Id
    private String id;    

    private String name;

    private String description;
    
}
