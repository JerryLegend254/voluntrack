function scroll(section) {
  element = document.getElementById(section);
  window.scrollTo(element.offsetLeft, element.offsetTop, {
    behavior: "smooth",
  });
}
