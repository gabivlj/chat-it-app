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

/**
 * Todo user image must be stored in cache
 */
export const SUBSCRIPTION_COMMENTS = gql`
  subscription onMessageAdded($postId: ID!) {
    newMessage(postId: $postId) {
      id
      createdAt
      text
      user {
        username
        id
      }
    }
  }
`;

export const SEND_MESSAGE = gql`
  mutation sendMessage($text: String!, $postId: ID!) {
    sendMessage(text: $text, postId: $postId, userId: "") {
      text
    }
  }
`;

export const POST_FRONTPAGE = gql`
  query posts($params: Params) {
    posts(params: $params) {
      text
      id
      title
      image {
        urlXL
      }
      user {
        username
        profileImage {
          urlXL
        }
      }
    }
  }
`;
