import express from "express";
import path from "path";
import { createLogger, format, transports } from "winston";

const __pathPrefix = "/assets";
const app = express();
const __dirmodule = path.join(__dirname, "node_modules");
const { combine, timestamp, label, simple } = format;
const logger = createLogger({
  level: "info",
  format: combine(timestamp(), label({ label: "assets" }), simple()),
  transports: [
    new transports.Console(),
    new transports.File({
      filename: `${__dirname}/log/asserts.log`,
    }),
  ],
});

app.use((req, _, next) => {
  logger.info(`Received a request for ${req.url}`);
  next();
});

app.use(__pathPrefix, express.static(path.join(__dirname, "generated")));

app.use(
  `${__pathPrefix}/_/htmx.org`,
  express.static(path.join(__dirmodule, "htmx.org/dist")),
);

app.use(
  `${__pathPrefix}/_/tailwindcss`,
  express.static(path.join(__dirmodule, "tailwindcss")),
);

app.use(
  `${__pathPrefix}/_/daisyui`,
  express.static(path.join(__dirmodule, "daisyui/dist")),
);

app.listen(45537);
logger.info("Server started");
