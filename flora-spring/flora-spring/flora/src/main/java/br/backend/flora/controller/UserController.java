package br.backend.flora.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

import br.backend.flora.model.UserModel;
import br.backend.flora.service.UserService;

@RestController
public class UserController {
    
    @Autowired
    private UserService usersvc;

    @GetMapping("/user/list/all")
    public Iterable<UserModel> listAll(){
        return usersvc.listAll();
    }

    @GetMapping("/")
    public String home(){
        return "Hello Flora World!";
    }

    @PostMapping("/user/register")
    public ResponseEntity<?> register(@RequestBody UserModel plant){
        return usersvc.register(plant);
    }

}
