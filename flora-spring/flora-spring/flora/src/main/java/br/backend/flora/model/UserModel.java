package br.backend.flora.model;

import java.util.Date;
import java.util.UUID;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;

import lombok.Getter;
import lombok.Setter;

@Entity
@Getter
@Setter
public class UserModel {
	@Id
	@GeneratedValue(strategy = GenerationType.AUTO)
    private UUID id;
	
	private String name;
	private String surname;
	private String email;
	private String password;
	private String loginMethod;
	private UUID profileImg;
	private Date lastLogin;
	private Date registerDate;
	private String location;
	private String language;
	private UUID vaseId;
}
