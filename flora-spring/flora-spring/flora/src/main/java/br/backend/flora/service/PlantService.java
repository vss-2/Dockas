package br.backend.flora.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;

import br.backend.flora.model.PlantModel;
import br.backend.flora.model.ResponseModel;
import br.backend.flora.repository.PlantRepository;

@Service
public class PlantService {
    
    @Autowired
    private PlantRepository plantrepo;

    @Autowired
    private ResponseModel response;

    public Iterable<PlantModel> listAll(){
        return plantrepo.findAll();
    }

    public ResponseEntity<?> register(PlantModel plant){
        if(plant.getSpecie().equals("")){
            response.setMessage("Plant's specie must be specified!");
            return new ResponseEntity<ResponseModel>(response, HttpStatus.BAD_REQUEST);
        }

        return new ResponseEntity<PlantModel>(plantrepo.save(plant), HttpStatus.CREATED);
    }

}
