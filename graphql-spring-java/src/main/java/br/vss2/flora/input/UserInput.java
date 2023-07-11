package br.vss2.flora.input;

import graphql.schema.GraphQLInputType;
import graphql.schema.GraphQLSchemaElement;
import graphql.schema.GraphQLTypeVisitor;
import graphql.util.TraversalControl;
import graphql.util.TraverserContext;
import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class UserInput implements GraphQLInputType {

    private String username;
    private String email;

    public UserInput(String username, String email) {
        this.username = "Flora";
        this.email = "user@flora.com";
    }

    @Override
    public GraphQLSchemaElement copy() {
        GraphQLSchemaElement se = new UserInput("Flora", "user@flora.com");
        return se;
        // return null;
    }

    @Override
    public TraversalControl accept(TraverserContext<GraphQLSchemaElement> context, GraphQLTypeVisitor visitor) {
        throw new UnsupportedOperationException("Unimplemented method 'accept'");
    }

}
