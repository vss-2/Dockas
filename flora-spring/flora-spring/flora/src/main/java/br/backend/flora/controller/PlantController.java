package br.backend.flora.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

import br.backend.flora.model.PlantModel;
import br.backend.flora.service.PlantService;

@RestController
public class PlantController {
    
    @Autowired
    private PlantService plantsvc;

    @GetMapping("/plant/list/all")
    public Iterable<PlantModel> listAll(){
        return plantsvc.listAll();
    }

    @PostMapping("/plant/register")
    public ResponseEntity<?> register(@RequestBody PlantModel plant){
        return plantsvc.register(plant);
    }

}
