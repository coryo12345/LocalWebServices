import { AxiosRequestConfig } from "axios";
import { RequestOptions, Result } from "./models";

export function buildAxiosOptions(
  options?: RequestOptions
): AxiosRequestConfig {
  const arc: AxiosRequestConfig = {};

  if (options?.signal) arc.signal = options.signal;

  return arc;
}

export class SdkResult<T> implements Result<T> {
  private value?: T;
  private error?: string;

  constructor(value?: T, error?: string) {
    this.value = value;
    this.error = error;
  }

  hasError() {
    return typeof this.error === "string";
  }

  hasValue() {
    return this.value !== undefined;
  }

  getValue() {
    if (this.hasError()) {
      throw new Error(
        "Attempted to access value on result with error: " + this.error ??
          "no error message provided to result"
      );
    }
    return this.value as T;
  }

  getError() {
    return this.error ?? "";
  }
}
