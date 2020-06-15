/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: checkLoged
// ====================================================

export interface checkLoged_loged_user {
  __typename: "User";
  username: string;
}

export interface checkLoged_loged {
  __typename: "Loged";
  user: checkLoged_loged_user | null;
}

export interface checkLoged {
  loged: checkLoged_loged;
}
