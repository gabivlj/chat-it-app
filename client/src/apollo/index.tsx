import { ApolloClient } from 'apollo-client';
import { InMemoryCache, NormalizedCacheObject } from 'apollo-cache-inmemory';
import { HttpLink } from 'apollo-link-http';
import { WebSocketLink } from 'apollo-link-ws';
import { split } from 'apollo-link';
import { getMainDefinition } from 'apollo-utilities';
import { setContext } from 'apollo-link-context';
import { mutationResolvers } from './mutation_resolvers';

const wsLink = new WebSocketLink({
  uri: `ws://localhost:8080/query`,
  options: {
    reconnect: true,
    connectionParams: () => {
      const token = localStorage.getItem('token');
      return {
        headers: {
          Authorization: token
        }
      };
    }
  }
});

export const cache = new InMemoryCache();
const httpLink = new HttpLink({
  uri: 'http://localhost:8080/query'
});

const linkUnifier = split(
  // split based on operation type
  ({ query }) => {
    const definition = getMainDefinition(query);
    return (
      definition.kind === 'OperationDefinition' &&
      definition.operation === 'subscription'
    );
  },
  wsLink,
  httpLink
);

const authLink = setContext((_, { headers }) => {
  // get the authentication token from local storage if it exists
  const token = localStorage.getItem('token');
  // return the headers to the context so httpLink can read them
  return {
    headers: {
      ...headers,
      authorization: token ? token : ''
    }
  };
});

export const client: ApolloClient<NormalizedCacheObject> = new ApolloClient({
  cache,
  link: authLink.concat(linkUnifier),
  resolvers: {
    Mutation: {
      ...mutationResolvers
    }
  }
});
