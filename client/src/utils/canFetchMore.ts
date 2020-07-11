type PropsMoreOnScroll = {
  data: any;
  loading: boolean;
};
export default function canFetchMore({ data, loading }: PropsMoreOnScroll) {
  const d = document.documentElement;
  const offset = d.scrollTop + window.innerHeight;
  const height = d.offsetHeight;
  if (offset !== height || loading || !data) {
    return false;
  }
  return true;
}
