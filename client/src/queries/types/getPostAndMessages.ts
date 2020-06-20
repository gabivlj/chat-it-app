/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

import { Params } from "./../../types/graphql-global-types";

// ====================================================
// GraphQL query operation: getPostAndMessages
// ====================================================

export interface getPostAndMessages_post_image {
  __typename: "Image";
  urlXL: string;
}

export interface getPostAndMessages_post_user_profileImage {
  __typename: "Image";
  urlXL: string;
}

export interface getPostAndMessages_post_user {
  __typename: "User";
  username: string;
  profileImage: getPostAndMessages_post_user_profileImage | null;
  id: string;
}

export interface getPostAndMessages_post {
  __typename: "Post";
  text: string;
  title: string;
  image: getPostAndMessages_post_image | null;
  user: getPostAndMessages_post_user;
}

export interface getPostAndMessages_messagesPost_user {
  __typename: "User";
  username: string;
  id: string;
}

export interface getPostAndMessages_messagesPost {
  __typename: "Message";
  id: string;
  user: getPostAndMessages_messagesPost_user;
  createdAt: number;
  text: string;
}

export interface getPostAndMessages {
  post: getPostAndMessages_post;
  messagesPost: getPostAndMessages_messagesPost[];
}

export interface getPostAndMessagesVariables {
  id: string;
  cursor?: Params | null;
}
