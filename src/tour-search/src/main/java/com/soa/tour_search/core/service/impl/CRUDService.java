package com.soa.tour_search.core.service.impl;

import java.util.NoSuchElementException;

import org.springframework.data.repository.CrudRepository;

import com.soa.tour_search.core.service.ICRUDService;

public abstract class CRUDService<T, ID> implements ICRUDService<T, ID> {

    private CrudRepository<T, ID> repository;
    
    public CRUDService(CrudRepository<T, ID> repository) {
        super();
        this.repository = repository;
    }

    public T findById(ID id) throws NoSuchElementException {
        return repository.findById(id).orElseThrow();
    }

    public Iterable<T> findAll() {
        return repository.findAll();
    }

    public Iterable<T> findAllByIds(Iterable<ID> ids) {
        return repository.findAllById(ids);
    }

    public T save(T entity) {
        return repository.save(entity);
    }

    public void deleteById(ID id) {
        repository.deleteById(id);
    }

    public void delete(T entity) {
        repository.delete(entity);
    }
    
}
