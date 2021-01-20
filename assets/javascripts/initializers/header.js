import Header, {
  DEFAULT_SELECTOR as HEADER_SELECTOR,
} from "../components/Header";

document.addEventListener("DOMContentLoaded", () => {
  const header = document.querySelector(HEADER_SELECTOR);

  header && new Header(header);
});
