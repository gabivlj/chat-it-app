import React, { useEffect, useState } from 'react';
import PostCard from '../components/Posts/PostCard';
import { useQuery } from '@apollo/react-hooks';
import {
  getPostAndMessages,
  getPostAndMessagesVariables
} from '../queries/types/getPostAndMessages';
import {
  GET_POST_AND_MESSAGES,
  SUBSCRIPTION_COMMENTS
} from '../queries/post_queries';
import Message from '../components/Chat/Message';
import { IS_LOGED_QUERY, CHECK_LOGED_LOCAL } from '../queries/user_queries';
import { isLogged } from '../queries/types/isLogged';
import Loading from '../components/Utils/Loading';
import NotFound from '../components/Utils/NotFound';
import {
  onMessageAdded,
  onMessageAddedVariables
} from '../queries/types/onMessageAdded';
import { userInfo } from 'os';
import InputChat from '../components/Chat/InputChat';
import WarningNotLoged from '../components/Utils/WarningNotLoged';

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
    params: { id }
  }
}: Props) {
  const [loadingMore, setLoadingMore] = useState(false);
  const [reachedEnd, setReachedEnd] = useState(false);
  const { data, loading, error, fetchMore, subscribeToMore } = useQuery<
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
  useEffect(
    () => () => {
      dataLoaded = false;
    },
    []
  );
  useEffect(() => {
    if (!data) return;
    // If we already subscribed, don't subscribe more
    if (!dataLoaded) {
      console.log('[SUBSCRIBED]');
      subscribeToMore<onMessageAdded, onMessageAddedVariables>({
        document: SUBSCRIPTION_COMMENTS,
        variables: {
          postId: id
        },
        updateQuery: (prev, { subscriptionData }) => {
          if (!subscriptionData.data) return prev;
          const newMessageItem = subscriptionData.data.newMessage;
          return Object.assign({}, prev, {
            messagesPost: [newMessageItem, ...prev.messagesPost]
          });
        }
      });
    }
    dataLoaded = true;
    const fetchMoreScroll = () => {
      // Check if we hitted bottom
      const d = document.documentElement;
      const offset = d.scrollTop + window.innerHeight;
      const height = d.offsetHeight;
      if (offset !== height || loading || !data) {
        return;
      }
      let reached;
      // Get reached end in a cool way
      setReachedEnd(prev => {
        reached = prev;
        return prev;
      });
      if (reached) return;
      setLoadingMore(true);
      fetchMore({
        variables: {
          id: id,
          cursor: {
            limit: 10,
            before:
              data && data.messagesPost.length
                ? data.messagesPost[data.messagesPost.length - 1].id
                : null
          }
        },
        updateQuery: (prev, { fetchMoreResult, variables }) => {
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
              p => p.id === fetchMoreResult.messagesPost[0].id
            ).length > 0
          ) {
            console.warn(
              'THERE ARE STILL ERRORS IN PAGINATION.',
              prev.messagesPost,
              fetchMoreResult.messagesPost
            );
            return Object.assign({}, prev, {
              messagesPost: [...prev.messagesPost]
            });
          }
          return Object.assign({}, prev, {
            messagesPost: [
              ...prev.messagesPost,
              ...fetchMoreResult.messagesPost
            ]
          });
        }
      });
    };
    window.onscroll = fetchMoreScroll;
    return () => {
      window.onscroll = () => {};
    };
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
      {loadingMore && !reachedEnd && <Loading className={'float-right mr-2'} />}
    </div>
  );
}
