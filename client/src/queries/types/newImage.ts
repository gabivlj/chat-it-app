/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL mutation operation: newImage
// ====================================================

export interface newImage_newProfileImage_profileImage {
  __typename: "Image";
  urlXL: string;
}

export interface newImage_newProfileImage {
  __typename: "User";
  id: string;
  profileImage: newImage_newProfileImage_profileImage | null;
}

export interface newImage {
  newProfileImage: newImage_newProfileImage;
}

export interface newImageVariables {
  image: any;
}
