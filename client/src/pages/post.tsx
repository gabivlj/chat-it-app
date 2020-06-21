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
import { IS_LOGED_QUERY } from '../queries/user_queries';
import { isLogged } from '../queries/types/isLogged';
import Loading from '../components/Utils/Loading';
import NotFound from '../components/Utils/NotFound';
import {
  onMessageAdded,
  onMessageAddedVariables
} from '../queries/types/onMessageAdded';
import { userInfo } from 'os';
import InputChat from '../components/Chat/InputChat';

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
  const resultIsLogged = useQuery<isLogged>(IS_LOGED_QUERY);
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
        {(!resultIsLogged.data || !resultIsLogged.data.loged.user) && (
          <div
            className={`rounded-lg p-2 bg-indigo-800 items-center text-indigo-100 leading-none  m-3 max-w-lg h-auto float-right clear-both w-lg sm:w-2/3 xl:w-1/3 md:w-1/3 lg:w-2/3 w-2/3`}
            role="alert"
          >
            <span className="rounded-full bg-indigo-500 uppercase px-2 py-1 text-xs font-bold mr-3 w-auto inline-block">
              Warning!
            </span>
            <span
              style={{ wordWrap: 'break-word' }}
              className="font-semibold mt-3 mr-2 text-left max-w-lg"
            >
              You are not logged in! If you want real-time chat you must be
              logged in!
            </span>
            <svg
              className="fill-current opacity-75 h-4 w-4"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 20 20"
            ></svg>
          </div>
        )}
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
