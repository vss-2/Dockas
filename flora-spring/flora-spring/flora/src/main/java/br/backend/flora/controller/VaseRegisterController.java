package br.backend.flora.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

import br.backend.flora.model.VaseRegisterModel;
import br.backend.flora.service.VaseRegisterService;

@RestController
public class VaseRegisterController {
    
    @Autowired
    private VaseRegisterService registersvc;

    @PostMapping("/vase/save")
    public ResponseEntity<?> saveRegister(@RequestBody VaseRegisterModel register){
        System.out.print(register.getTemperature());
        return registersvc.saveStatus(register);
    }

}
