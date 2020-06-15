/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

import { FormLogInRegister } from "./../../types/graphql-global-types";

// ====================================================
// GraphQL mutation operation: logUser
// ====================================================

export interface logUser_logUser_user {
  __typename: "User";
  username: string;
  id: string;
}

export interface logUser_logUser {
  __typename: "UserSession";
  user: logUser_logUser_user;
  session: string;
}

export interface logUser {
  logUser: logUser_logUser;
}

export interface logUserVariables {
  formParameters?: FormLogInRegister | null;
}
