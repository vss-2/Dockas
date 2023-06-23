package br.backend.flora.repository;

import java.util.UUID;

import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

import br.backend.flora.model.UserModel;

@Repository
public interface UserRepository extends CrudRepository<UserModel, UUID>{
    
    // public Iterable<UserModel> listAll();

    // public UserModel register(UserModel user);

}
