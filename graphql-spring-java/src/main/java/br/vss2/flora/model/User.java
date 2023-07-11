package br.vss2.flora.model;

import java.util.Date;
import java.util.UUID;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;
import javax.persistence.Table;

import org.hibernate.annotations.GenericGenerator;

import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
@Entity
@Table(name = "USERS")
public class User {

    public User(String username, String email) {
        this.id = UUID.randomUUID();
        this.username = username;
        this.email = email;
        Date date = new Date();
        this.signup = date.hashCode();
    }

    public User() {
        this.id = UUID.randomUUID();
        this.username = "Flora";
        this.email = "user@flora.com";
        this.signup = 1;
    }

    @Id
    @GeneratedValue(generator = "uuid2")
    @GenericGenerator(name = "uuid2", strategy = "org.hibernate.id.UUIDGenerator")
    @Column(name = "id", columnDefinition = "VARCHAR(255)")
    private UUID id;

    @Column(name = "USERNAME")
    private String username = "Flora";

    @Column(name = "EMAIL")
    private String email = "user@flora.com";

    @Column(name = "SIGNUP")
    private int signup = 1;

}
