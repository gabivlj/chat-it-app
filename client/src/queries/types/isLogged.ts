/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: isLogged
// ====================================================

export interface isLogged_loged_user {
  __typename: "User";
  username: string;
  id: string;
}

export interface isLogged_loged {
  __typename: "Loged";
  user: isLogged_loged_user | null;
}

export interface isLogged {
  loged: isLogged_loged;
}
