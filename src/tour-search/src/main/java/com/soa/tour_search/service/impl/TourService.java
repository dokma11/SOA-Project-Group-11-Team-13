package com.soa.tour_search.service.impl;

import com.soa.tour_search.core.service.impl.CRUDService;
import com.soa.tour_search.model.Tour;
import com.soa.tour_search.repository.TourRepository;

public class TourService extends CRUDService<Tour, Integer> {

    public TourService(TourRepository repository) {
        super(repository);
    }
    
}
