import AppHeader, {
  DEFAULT_SELECTOR as HEADER_SELECTOR,
} from "../components/AppHeader";

document.addEventListener("DOMContentLoaded", () => {
  const appHeader = document.querySelector(HEADER_SELECTOR);

  appHeader && new AppHeader(appHeader);
});
