package com.soa.tour_search.service.impl;

import org.springframework.stereotype.Service;

import com.soa.tour_search.core.service.impl.CRUDService;
import com.soa.tour_search.model.Tour;
import com.soa.tour_search.repository.TourRepository;
import com.soa.tour_search.service.ITourService;

@Service
public class TourService extends CRUDService<Tour, String> implements ITourService {

    public TourService(TourRepository repository) {
        super(repository);
    }
    
}
