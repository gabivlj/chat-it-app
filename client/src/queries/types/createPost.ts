/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

import { PostForm } from "./../../types/graphql-global-types";

// ====================================================
// GraphQL mutation operation: createPost
// ====================================================

export interface createPost_newPost_image {
  __typename: "Image";
  urlXL: string;
}

export interface createPost_newPost_user_profileImage {
  __typename: "Image";
  urlXL: string;
}

export interface createPost_newPost_user {
  __typename: "User";
  username: string;
  profileImage: createPost_newPost_user_profileImage | null;
}

export interface createPost_newPost {
  __typename: "Post";
  text: string;
  title: string;
  id: string;
  image: createPost_newPost_image | null;
  user: createPost_newPost_user;
  numberOfComments: number;
}

export interface createPost {
  newPost: createPost_newPost;
}

export interface createPostVariables {
  form: PostForm;
}
