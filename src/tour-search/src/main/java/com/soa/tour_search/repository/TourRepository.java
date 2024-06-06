package com.soa.tour_search.repository;

import org.springframework.data.elasticsearch.repository.ElasticsearchRepository;
import org.springframework.stereotype.Repository;

import com.soa.tour_search.model.Tour;

@Repository
public interface TourRepository extends ElasticsearchRepository<Tour, String> {
    
}
