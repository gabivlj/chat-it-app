import { GraphQLError } from 'graphql';
import { ApolloError } from 'apollo-client';

export class ParseError {
  static parseLoginError(err: ApolloError | GraphQLError) {
    const { message } = err;
    const split = message.split(':');
    const [, messageError] = split;
    const errorObject: LoginError = {
      username: '',
      password: ''
    };
    if (messageError === '') {
      return errorObject;
    }
    if (messageError.includes('password')) {
      errorObject.password = messageError.trim();
      return errorObject;
    }
    if (messageError.includes('User')) {
      errorObject.username = messageError.trim();
      return errorObject;
    }
    return errorObject;
  }
}

export type LoginError = {
  username: string;
  password: string;
};
