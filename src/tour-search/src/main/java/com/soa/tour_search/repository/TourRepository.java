package com.soa.tour_search.repository;

import org.springframework.data.elasticsearch.repository.ElasticsearchRepository;

import com.soa.tour_search.model.Tour;

public interface TourRepository extends ElasticsearchRepository<Tour, Integer> {
    
}
