import React, { useEffect, useState } from 'react';
import PostCard from '../components/Posts/PostCard';
import { useQuery } from '@apollo/react-hooks';
import {
  getPostAndMessages,
  getPostAndMessagesVariables,
} from '../queries/types/getPostAndMessages';
import {
  GET_POST_AND_MESSAGES,
  SUBSCRIPTION_COMMENTS,
} from '../queries/post_queries';
import Message from '../components/Chat/Message';
import { CHECK_LOGED_LOCAL } from '../queries/user_queries';
import { isLogged } from '../queries/types/isLogged';
import Loading from '../components/Utils/Loading';
import NotFound from '../components/Utils/NotFound';
import {
  onMessageAdded,
  onMessageAddedVariables,
} from '../queries/types/onMessageAdded';
import InputChat from '../components/Chat/InputChat';
import WarningNotLoged from '../components/Utils/WarningNotLoged';
import useIsBottom from '../utils/useIsBottom';

type Props = {
  match: {
    params: {
      id: string;
    };
  };
};
let dataLoaded = false;
export default function Post({
  match: {
    params: { id },
  },
}: Props) {
  const [loadingMore, setLoadingMore] = useState(false);
  const [reachedEnd, setReachedEnd] = useState(false);
  const isBottom = useIsBottom();
  const { data, loading, error, fetchMore, subscribeToMore } = useQuery<
    getPostAndMessages,
    getPostAndMessagesVariables
  >(GET_POST_AND_MESSAGES, {
    variables: {
      id,
      cursor: {
        limit: 10,
      },
    },
  });
  useEffect(() => {
    if (!isBottom || !data) return;
    let reached;
    // Get reached end in a cool way
    setReachedEnd((prev) => {
      reached = prev;
      return prev;
    });
    if (reached) return;
    const before = data.messagesPost.length
      ? data.messagesPost[data.messagesPost.length - 1].id
      : null;
    const limit = 10;
    setLoadingMore(true);
    fetchMore({
      variables: {
        id,
        cursor: {
          limit,
          before,
        },
      },
      updateQuery: (prev, { fetchMoreResult }) => {
        setLoadingMore(false);
        if (!fetchMoreResult) {
          return prev;
        }
        if (!fetchMoreResult.messagesPost.length) {
          setReachedEnd(true);
          return prev;
        }
        // Check if there are errors for pagination
        if (
          prev.messagesPost.filter(
            (p) => p.id === fetchMoreResult.messagesPost[0].id
          ).length > 0
        ) {
          console.warn(
            'THERE ARE STILL ERRORS IN PAGINATION.',
            prev.messagesPost,
            fetchMoreResult.messagesPost
          );
          return Object.assign({}, prev, {
            messagesPost: [...prev.messagesPost],
          });
        }
        return Object.assign({}, prev, {
          messagesPost: [...prev.messagesPost, ...fetchMoreResult.messagesPost],
        });
      },
    });
  }, [isBottom, data]);
  useEffect(
    () => () => {
      dataLoaded = false;
    },
    []
  );
  useEffect(() => {
    if (!data) return;
    // If we already subscribed, don't subscribe more
    if (dataLoaded) return;
    subscribeToMore<onMessageAdded, onMessageAddedVariables>({
      document: SUBSCRIPTION_COMMENTS,
      variables: {
        postId: id,
      },
      updateQuery: (prev, { subscriptionData }) => {
        if (!subscriptionData.data) return prev;
        const newMessageItem = subscriptionData.data.newMessage;
        return Object.assign({}, prev, {
          messagesPost: [newMessageItem, ...prev.messagesPost],
        });
      },
    });
    dataLoaded = true;
  }, [data]);
  const resultIsLogged = useQuery<isLogged>(CHECK_LOGED_LOCAL);
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
      <div>
        <PostCard post={data.post} />
        {/* <div className="bg-indigo-900 text-center py-4 lg:px-4"> */}
        <h1 className="font-bold pt-8 pl-3 text-5xl text-center mb-10">
          Live Comments
        </h1>
        <WarningNotLoged />
        <div className="container pt-3">
          {resultIsLogged.data && resultIsLogged.data.loged.user && (
            <InputChat postId={id} />
          )}

          <div className="mt-3">
            {data.messagesPost.map((message) => (
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
      {loadingMore && !reachedEnd && <Loading className={'float-right mr-2'} />}
    </div>
  );
}
