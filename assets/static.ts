import express from "express";
import path from "path";

const app = express();
const __dirmodule = path.join(__dirname, "node_modules");

app.use(express.static(path.join(__dirname, "generated")));

app.use("/_/htmx.org", app.use(path.join(__dirmodule, "htmx.org/dist")));

app.use("/_/tailwindcss", app.use(path.join(__dirmodule, "tailwindcss")));

app.use("/_/daisyui", app.use(path.join(__dirmodule, "daisyui/dist")));

app.listen(8080);
