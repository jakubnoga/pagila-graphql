import { readFileSync } from "node:fs";
import { ApolloServer } from "apollo-server";
import { Resolvers } from "./resolvers/resolvers-types";
import { createRepository } from "./repository";

const typeDefs = readFileSync("./schema.graphql", "utf8");

(async () => {
  const repo = await createRepository();

  const resolvers: Resolvers = {
    Query: {
      film: (_, args) => repo.getFilm(args.title),
      films: (_, args) => repo.getFilms(args.first),
    },
    Film: {
      actors: async (parent) => repo.getActors(parent.film_id),
    },
    Actor: {
      firstName: (parent) => parent.first_name,
      lastName: (parent) => parent.last_name,
    },
    Mutation: {
      changeLength: (_, args) =>
        repo
          .changeLength(args.title, args.newLength)
          .then(({ film_id }) => film_id),
    },
  };

  const server = new ApolloServer({ typeDefs, resolvers, introspection: true });
  // The `listen` method launches a web server
  server.listen().then(({ url }) => {
    console.log(`🚀  Server ready at ${url}`);
  });
})();
