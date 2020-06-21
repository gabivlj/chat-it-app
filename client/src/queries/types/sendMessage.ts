/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL mutation operation: sendMessage
// ====================================================

export interface sendMessage_sendMessage {
  __typename: "Message";
  text: string;
}

export interface sendMessage {
  sendMessage: sendMessage_sendMessage;
}

export interface sendMessageVariables {
  text: string;
  postId: string;
}
