import React, { useEffect } from 'react';
import PostCard from '../components/Posts/PostCard';
import { useQuery } from '@apollo/react-hooks';
import {
  getPostAndMessages,
  getPostAndMessagesVariables
} from '../queries/types/getPostAndMessages';
import { GET_POST_AND_MESSAGES } from '../queries/post_queries';
import Message from '../components/Chat/Message';
import { IS_LOGED_QUERY } from '../queries/user_queries';
import { isLogged } from '../queries/types/isLogged';
import Loading from '../components/Utils/Loading';
import NotFound from '../components/Utils/NotFound';

type Props = {
  match: {
    params: {
      id: string;
    };
  };
};

export default function Post({
  match: {
    params: { id }
  }
}: Props) {
  const resultIsLogged = useQuery<isLogged>(IS_LOGED_QUERY);
  const { data, loading, error, fetchMore } = useQuery<
    getPostAndMessages,
    getPostAndMessagesVariables
  >(GET_POST_AND_MESSAGES, {
    variables: {
      id,
      cursor: {
        limit: 10
      }
    }
  });
  // useEffect(() => {
  //   if (!data) return;
  //   setTimeout(() => {
  //     fetchMore({
  //       variables: {
  //         id: match.params.id,
  //         cursor: {
  //           limit: 2,
  //           before: data && data.messagesPost[data.messagesPost.length - 1].id
  //         }
  //       },
  //       updateQuery: (prev, { fetchMoreResult, variables }) => {
  //         console.log(variables);
  //         if (!fetchMoreResult) return prev;
  //         console.log(fetchMoreResult, prev);
  //         return Object.assign({}, prev, {
  //           messagesPost: [
  //             ...prev.messagesPost,
  //             ...fetchMoreResult.messagesPost
  //           ]
  //         });
  //       }
  //     });
  //   }, 2000);
  // }, [data]);
  if (loading) {
    return <Loading />;
  }
  if (error) {
    return <NotFound />;
  }
  if (!data) {
    return <></>;
  }
  return (
    <div>
      <PostCard post={data.post} />
      {/* <div className="bg-indigo-900 text-center py-4 lg:px-4"> */}
      <h1 className="text-3xl font-bold pt-8 pl-3">Comments</h1>
      <div className="container pt-3">
        <div className="mt-3">
          {data.messagesPost.map(message => (
            <Message
              message={message}
              key={message.id}
              currentUser={
                resultIsLogged.data && resultIsLogged.data.loged.user
              }
            />
          ))}
        </div>
      </div>
    </div>
  );
}
