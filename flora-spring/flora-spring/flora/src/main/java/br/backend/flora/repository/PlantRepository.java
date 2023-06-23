package br.backend.flora.repository;

import java.util.UUID;

import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

import br.backend.flora.model.PlantModel;

@Repository
public interface PlantRepository extends CrudRepository<PlantModel, UUID>{
	
}
