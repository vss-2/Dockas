package br.backend.flora.service;

import java.util.ArrayList;
import java.util.List;
import java.util.UUID;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;

import br.backend.flora.model.VaseRegisterModel;
import br.backend.flora.repository.VaseRegisterRepository;

@Service
public class VaseRegisterService {

    @Autowired
    private VaseRegisterRepository vaserepo;

    public Iterable<VaseRegisterModel> findVaseById(UUID id){
        List<UUID> ids = new ArrayList<UUID>();
        ids.add(id);
        return vaserepo.findAllById(ids);
    }

    public ResponseEntity<?> saveStatus(VaseRegisterModel register){
        try {
            vaserepo.save(register);
        } catch (Exception e) {
            System.out.print("\n\n\n----EXCEPTION----\n\n\n");
            System.out.print(e.toString());
            System.out.print("\n\n\n----EXCEPTION----\n\n\n");
            return new ResponseEntity<HttpStatus>(HttpStatus.NOT_ACCEPTABLE);
        }
        return new ResponseEntity<VaseRegisterModel>(register, HttpStatus.ACCEPTED);
    }

}
