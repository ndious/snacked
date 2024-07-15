import { __assetsdir, __defaultport } from "./constants";
import { Config, LogLevel } from "./types";
import { join } from "path";

const defaultConfig: Config = {
  prefix: "",
  log: LogLevel.info,
  default: "css",
  assets: [],
  vendor: [],
  port: __defaultport,
};

export const readConfigFile = async (): Promise<Config> => {
  const file = Bun.file(join(__assetsdir, "/assets.json"));
  let userConfig = {} as Config;

  if (await file.exists()) {
    userConfig = await file.json();
  }

  if (userConfig.log) {
    userConfig.log = LogLevel[userConfig.log];
  }

  return {
    ...defaultConfig,
    ...userConfig,
  } as Config;
};
