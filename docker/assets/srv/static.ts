import express, { Application } from "express";
import { join } from "path";
import { Logger, createLogger, format, transports } from "winston";
import { readConfigFile } from "./config";
import { __assetsdir, __logdir, __vendordir } from "./constants";

const startServer = async (app: Application, logger: Logger) => {
  const config = await readConfigFile();

  const __pathPrefix = config.prefix;

  app.use((req, _, next) => {
    logger.info(`Received a request for ${req.url}`);
    next();
  });

  app.use(__pathPrefix, express.static(join(__assetsdir, config.default)));
  logger.info(`Default route pointing to ${config.default}`);

  config.vendor.forEach((item) => {
    const uri = `${__pathPrefix}/_/${item.uri}`;

    app.use(uri, express.static(join(__vendordir, item.dir)));

    logger.info(`${uri} loaded !`);
  });

  app.listen(config.port, () => {
    logger.info(`Server started on port ${config.port}`);
  });
};

const app = express();

const { combine, timestamp, label, simple } = format;
const logger = createLogger({
  level: "info",
  format: combine(timestamp(), label({ label: "assets" }), simple()),
  transports: [
    new transports.Console(),
    new transports.File({
      filename: `${__logdir}/asserts.log`,
    }),
  ],
});

startServer(app, logger);
