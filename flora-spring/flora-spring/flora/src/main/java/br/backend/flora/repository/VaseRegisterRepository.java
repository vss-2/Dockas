package br.backend.flora.repository;

import java.util.UUID;

import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.stereotype.Repository;

import br.backend.flora.model.VaseRegisterModel;

@Repository
public interface VaseRegisterRepository extends MongoRepository<VaseRegisterModel, UUID> {
    // public Iterable<VaseRegisterModel> findVaseById(UUID id);

    // public ResponseEntity<?> saveRegister();
}
