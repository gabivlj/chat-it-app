const scroll = function(
  element: any,
  top: number,
  targetY: number,
  startingPoint: number
) {
  startingPoint++;
  if (startingPoint > 30) return;
  element.scrollTop = top + ((targetY - top) / 30) * startingPoint;
  setTimeout(function() {
    scroll(element, top, targetY, startingPoint);
  }, 20);
};

export function smoothScroll(target: any) {
  let scrollContainer = target;
  do {
    //find scroll container
    scrollContainer = scrollContainer.parentNode;
    if (!scrollContainer) return;
    scrollContainer.scrollTop += 1;
  } while (scrollContainer.scrollTop == 0);

  let targetY = 0;
  do {
    //find the top of target relatively to the container
    if (target == scrollContainer) break;
    targetY += target.offsetTop;
  } while ((target = target.offsetParent));

  // start scrolling
  scroll(scrollContainer, scrollContainer.scrollTop, targetY, 0);
}
