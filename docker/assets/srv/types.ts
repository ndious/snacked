type Vendor = {
  uri: string;
  dir: string;
};

type Asset = {
  prefix: string;
  dir: string
}

export enum LogLevel {
  dev = "dev",
  info = "info",
  warn = "warn",
  fatal = "fatal",
}

export type Config = {
  prefix: string;
  log: LogLevel;
  default: string;
  vendor: Vendor[];
  asset: Asset[];
  port: number;
};
