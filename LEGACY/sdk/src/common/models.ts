export interface RequestOptions {
  signal?: AbortSignal;
}

export interface Result<T> {
  hasError: () => boolean;
  hasValue: () => boolean;
  getValue: () => T;
  getError: () => string;
}
