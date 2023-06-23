package br.backend.flora.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;

import br.backend.flora.model.UserModel;
import br.backend.flora.repository.UserRepository;

@Service
public class UserService {

    @Autowired
    private UserRepository userrepo;

    public Iterable<UserModel> listAll(){
        return userrepo.findAll();
    }

    public ResponseEntity<?> register(UserModel user){
        return new ResponseEntity<UserModel>(userrepo.save(user), HttpStatus.CREATED);
    }
    
}
