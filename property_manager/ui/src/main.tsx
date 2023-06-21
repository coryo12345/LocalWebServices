import "preact/debug";
import { render } from "preact";
import { App } from "./app.tsx";
import "./index.css";

// dropdown
import "@szhsin/react-menu/dist/index.css";
import "@szhsin/react-menu/dist/transitions/slide.css";

render(<App />, document.getElementById("app") as HTMLElement);
