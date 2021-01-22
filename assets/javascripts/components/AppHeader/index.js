export const DEFAULT_SELECTOR = "header.app-header";
const SELECTORS = {
  profileMenu: ".profile-menu",
  profileMenuButton: ".profile-icon__button",
};

const CLASSES = {
  profileMenuShow: "profile-menu--show",
};

class AppHeader {
  constructor(elementRef) {
    this.elementRef = elementRef;
    this.profileMenu = this.elementRef.querySelector(SELECTORS.profileMenu);
    this.profileMenuButton = this.elementRef.querySelector(
      SELECTORS.profileMenuButton
    );

    this._bind();
    this._addEventListeners();
  }

  // Event handler

  onProfileMenuClick(e) {
    this._toogleProfileMenu();
  }

  // Private

  _bind() {
    this.onProfileMenuClick = this.onProfileMenuClick.bind(this);
  }

  _addEventListeners() {
    this.profileMenuButton.addEventListener("click", this.onProfileMenuClick);
  }

  _toogleProfileMenu() {
    this.profileMenu.classList.toggle(CLASSES.profileMenuShow);
  }
}

export default AppHeader;
