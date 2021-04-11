// CSS
import "../stylesheets/app-base.scss";
import "../stylesheets/app-components.scss";
import "../stylesheets/app-utilities.scss";

// Stimulus
import { Application } from "stimulus";
import { definitionsFromContext } from "stimulus/webpack-helpers";
import { Tabs } from "tailwindcss-stimulus-components";

const application = Application.start();
const context = require.context("./controllers", true, /\.js$/);
application.load(definitionsFromContext(context));
application.register("tabs", Tabs);
