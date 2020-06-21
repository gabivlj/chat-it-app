/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL subscription operation: onMessageAdded
// ====================================================

export interface onMessageAdded_newMessage_user {
  __typename: "User";
  username: string;
  id: string;
}

export interface onMessageAdded_newMessage {
  __typename: "Message";
  id: string;
  createdAt: number;
  text: string;
  user: onMessageAdded_newMessage_user;
}

export interface onMessageAdded {
  newMessage: onMessageAdded_newMessage;
}

export interface onMessageAddedVariables {
  postId: string;
}
