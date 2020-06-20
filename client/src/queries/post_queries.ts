import gql from 'graphql-tag';

export const GET_POST_AND_MESSAGES = gql`
  query getPostAndMessages($id: ID!, $cursor: Params) {
    post(id: $id) {
      text
      title
      image {
        urlXL
      }
      user {
        username
        profileImage {
          urlXL
        }
        id
      }
    }
    messagesPost(id: $id, params: $cursor) {
      id
      user {
        username
        id
      }
      createdAt
      text
    }
  }
`;
