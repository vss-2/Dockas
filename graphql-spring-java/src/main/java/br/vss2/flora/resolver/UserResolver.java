package br.vss2.flora.resolver;

import java.util.List;
import java.util.Optional;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import br.vss2.flora.input.UserInput;
import br.vss2.flora.model.User;
import br.vss2.flora.repository.UserRepository;
import graphql.kickstart.tools.GraphQLMutationResolver;
import graphql.kickstart.tools.GraphQLQueryResolver;

@Component
public class UserResolver implements GraphQLQueryResolver, GraphQLMutationResolver {

    // Dependency injection
    @Autowired
    private UserRepository repository;

    public Optional<User> findUserById(String user_id) {
        return repository.findById(user_id);
    }

    public List<User> findAllUsers() {
        return repository.findAll();
    }

    public User findOneUser() {
        User user = new User();
        return user;
    }

    public User saveUser(UserInput new_user) {
        return repository.save(new User(new_user.getUsername(), new_user.getEmail()));
    }

}
