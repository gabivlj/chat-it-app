import { ApolloClient } from 'apollo-client';
import { InMemoryCache, NormalizedCacheObject } from 'apollo-cache-inmemory';
import { HttpLink } from 'apollo-link-http';
import { WebSocketLink } from 'apollo-link-ws';
import { split } from 'apollo-link';
import { getMainDefinition } from 'apollo-utilities';
import { setContext } from 'apollo-link-context';
import { createUploadLink } from 'apollo-upload-client';
import { mutationResolvers } from './mutation_resolvers';

const uri = `localhost:8070`;
const uriBuild =
  process.env.NODE_ENV === 'development' ? uri : process.env.REACT_APP_URI_API;
const wsLink = new WebSocketLink({
  uri: `ws://${uriBuild || uri}/query`,
  options: {
    reconnect: true,
    connectionParams: () => {
      const token = localStorage.getItem('token');
      return {
        fetchOptions: {
          mode: 'cors',
        },
        headers: {
          Authorization: token,
        },
      };
    },
  },
});

export const cache = new InMemoryCache();
const httpLink = createUploadLink({
  uri: `http://${uriBuild || uri}/query`,
  fetchOptions: {
    mode: 'cors',
  },
});

const linkUnifier = split(
  // split based on operation type
  ({ query }) => {
    const definition = getMainDefinition(query);
    console.log(definition);
    return (
      definition.kind === 'OperationDefinition' &&
      definition.operation === 'subscription'
    );
  },
  wsLink,
  httpLink
);

const authLink = setContext((operation, { headers, credentials }) => {
  // get the authentication token from local storage if it exists
  const token = localStorage.getItem('token');
  // return the headers to the context so httpLink can read them
  return {
    headers: {
      ...headers,
      authorization: token ? token : '',
    },
    fetchOptions: {
      mode: 'cors',
    },
  };
});

export const client: ApolloClient<NormalizedCacheObject> = new ApolloClient({
  cache,
  link: authLink.concat(linkUnifier),
  resolvers: {
    Mutation: {
      ...mutationResolvers,
    },
  },
});
