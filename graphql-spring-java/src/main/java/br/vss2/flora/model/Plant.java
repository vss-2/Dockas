package br.vss2.flora.model;

import java.util.UUID;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;
import javax.persistence.Table;

import org.hibernate.annotations.GenericGenerator;

import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter
@NoArgsConstructor
@Entity
@Table(name = "PLANTS")
public class Plant {

    @Id
    @GeneratedValue(generator = "uuid2")
    @GenericGenerator(name = "uuid2", strategy = "org.hibernate.id.UUIDGenerator")
    @Column(name = "id", columnDefinition = "VARCHAR(255)")
    private UUID id;
    
    @Column(name = "PLANT_TYPE")
    private String type;

    @Column(name = "USER_ID")
    private String owner_user_id;

    @Column(name = "int_PLANTED")
    private int planted;
    
    @Column(name = "intS_HARVESTS")
    private int harvests;
    
    @Column(name = "int_DIED")
    private int died;
}
