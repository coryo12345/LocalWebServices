export interface FileMeta {
  name: string;
  size: number;
  isDir: boolean;
  lastModified: string; // maybe Date?
  relativePath: string;
  realFile: boolean;
}
