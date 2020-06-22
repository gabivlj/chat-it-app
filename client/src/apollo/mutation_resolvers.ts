import { isLogged_loged_user, isLogged } from '../queries/types/isLogged';
import { logUser_logUser } from '../queries/types/logUser';

export const mutationResolvers = {
  /**
   * log is a mutation query that inserts the desired user into the client data
   * @param User user
   */
  log: (
    _root: any,
    variables: { user: isLogged_loged_user | null },
    { client }: any
  ) => {
    const data: isLogged = {
      // We build the user object
      loged: {
        user: variables.user ? { ...variables.user, __typename: 'User' } : null,
        __typename: 'Loged'
      }
    };
    client.writeData({ data });

    return null;
  },

  /**
   * log session stores in the localStorage the token session and saves in the local state the user.
   *
   * @param user UserSession
   */
  logSession: (
    _root: any,
    variables: { user: logUser_logUser | null },
    { client }: any
  ) => {
    const data: isLogged = {
      // We build the user object
      loged: {
        user: variables.user
          ? { ...variables.user.user, __typename: 'User' }
          : null,
        __typename: 'Loged'
      }
    };
    console.log('ewee');
    client.writeData({ data });
    if (variables.user) localStorage.setItem('token', variables.user.session);
    return null;
  }
};
