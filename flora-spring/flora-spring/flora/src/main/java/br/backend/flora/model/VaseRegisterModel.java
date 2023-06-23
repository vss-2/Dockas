package br.backend.flora.model;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;

import org.springframework.data.mongodb.core.mapping.Document;

import lombok.Getter;
import lombok.Setter;

@Entity
@Document(collection = "vase_status")
@Getter
@Setter
public class VaseRegisterModel {
    
    @Id
    @GeneratedValue()
    private String id;
    private short temperature;
    private short luminosity;
    private short umidity;

}
