import { Controller } from "stimulus";
import { useClickOutside } from "stimulus-use";

export default class extends Controller {
  static targets = ["profileMenu"];
  static classes = ["profileMenuShow"];

  connect() {
    useClickOutside(this);
  }

  toggleProfileMenu() {
    this.profileMenuTarget.classList.toggle(this.profileMenuShowClass);
  }

  clickOutside(e) {
    this.profileMenuTarget.classList.remove(this.profileMenuShowClass);
  }
}
