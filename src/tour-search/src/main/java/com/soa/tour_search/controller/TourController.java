package com.soa.tour_search.controller;

import java.util.ArrayList;
import java.util.Collection;
import java.util.List;
import java.util.NoSuchElementException;
import java.util.stream.Collectors;

import org.modelmapper.ModelMapper;
import org.springframework.web.bind.annotation.RestController;

import com.soa.tour_search.dto.TourRequestDTO;
import com.soa.tour_search.dto.TourResponseDTO;
import com.soa.tour_search.model.Tour;
import com.soa.tour_search.service.ITourService;

import lombok.RequiredArgsConstructor;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;


@RestController
@RequestMapping(value = "tours")
@RequiredArgsConstructor
public class TourController {

    private final ITourService tourService;

    private final ModelMapper modelMapper;

    @GetMapping
    public ResponseEntity<Collection<TourResponseDTO>> findAll() {
        var tours = convertToList(tourService.findAll());
        var dtos = tours.stream().map(tour -> modelMapper.map(tour, TourResponseDTO.class)).collect(Collectors.toList());
        return ResponseEntity.ok(dtos);
    }

    @GetMapping("{id}")
    public ResponseEntity<?> findById(@PathVariable Integer id) {
        try {
            var tour = tourService.findById(id);
            var dto = modelMapper.map(tour, TourResponseDTO.class);
            return ResponseEntity.ok().body(dto);
        } catch (NoSuchElementException e) {
            return new ResponseEntity<>(HttpStatus.NOT_FOUND);
        }
    }

    @PostMapping
    public void create(@RequestBody TourRequestDTO request) {
        var tour = modelMapper.map(request, Tour.class);
        tourService.save(tour);
    }

    @DeleteMapping("{id}")
    public void deleteById(@PathVariable Integer id) {
        tourService.deleteById(id);
    }

    private List<Tour> convertToList(Iterable<Tour> tours) {
        List<Tour> toursList = new ArrayList<>();
        tours.forEach(toursList::add);
        return toursList;
    }
    
}
