import gql from 'graphql-tag';

// apollo schema:download
// apollo codegen:generate --localSchemaFile=schema.json --target=typescript --includes=src/**/*.ts --tagName=gql --addTypename --globalTypesFile=src/types/graphql-global-types.ts types

export const CHECK_LOGED_LOCAL = gql`
  query checkLoged {
    loged @client {
      user {
        username
      }
    }
  }
`;

export const IS_LOGED_QUERY = gql`
  query isLogged {
    loged {
      user {
        username
        id
      }
    }
  }
`;

export const LOG_USER_MUTATION = gql`
  mutation logUser($formParameters: FormLogInRegister) {
    logUser(parameters: $formParameters) {
      user {
        username
        id
      }
      session
    }
  }
`;

export const LOG_USER_LOCAL = gql`
  mutation LogUserLocal($user: Loged) {
    log(user: $user) @client
  }
`;

export const LOG_USER_LOCAL_SESSION = gql`
  mutation LogUserLocalSession($user: UserSession) {
    logSession(user: $user) @client
  }
`;

export const REGISTER_USER_MUTATION = gql`
  mutation registerUser($username: String, $password: String) {
    newUser(parameters: { username: $username, password: $password }) {
      user {
        username
        id
      }
      session
    }
  }
`;

export const GET_USER = gql`
  query GetUser($query: UserQuery!) {
    user(id: $query) {
      username
      id
      profileImage {
        urlXL
      }
      posts {
        text
        image {
          urlXL
        }
        title
        id
      }
    }
  }
`;

export const UPLOAD_PROFILE_IMAGE = gql`
  mutation newImage($image: Upload!) {
    newProfileImage(image: $image) {
      id
      profileImage {
        urlXL
      }
    }
  }
`;
