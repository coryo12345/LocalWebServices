import axios, { AxiosInstance, AxiosResponse } from "axios";
import { RequestOptions, Result } from "../common/models";
import { FileMeta } from "./models";
import { SdkResult, buildAxiosOptions } from "../common";
import { AxiosError } from "axios";
import { constants } from "../constants";

export { FileMeta };

export function FileStorage(url: string) {
  const client: AxiosInstance = axios.create({
    baseURL: url,
    timeout: 2000,
  });

  async function listFiles(
    path: string,
    options?: RequestOptions
  ): Promise<Result<FileMeta[]>> {
    const config = buildAxiosOptions(options);
    const pathStr = encodeURIComponent(path);

    try {
      const resp = await client.get<FileMeta[]>(`/dir?path=${pathStr}`, config);

      if (typeof resp.data === "object") {
        return new SdkResult<FileMeta[]>(resp.data);
      }
      return new SdkResult<FileMeta[]>(undefined, "Unable to fetch files");
    } catch (err) {
      return new SdkResult<FileMeta[]>(undefined, "Unable to fetch files");
    }
  }

  async function downloadFile(
    path: string,
    downloadName: string,
    options?: RequestOptions
  ): Promise<Result<boolean>> {
    const config = buildAxiosOptions(options);
    config.responseType = "blob";
    const pathStr = encodeURIComponent(path);

    let resp: AxiosResponse;

    try {
      resp = await client.get(`/file?path=${pathStr}`, config);

      const href = URL.createObjectURL(resp.data);

      const link = document.createElement("a");
      link.href = href;
      link.setAttribute("download", downloadName);
      document.body.appendChild(link);
      link.click();

      document.body.removeChild(link);
      URL.revokeObjectURL(href);

      return new SdkResult(true);
    } catch (err) {
      let msg = constants.UNKNOWN_ERROR;

      if ((err as AxiosError<Blob>).response?.data.type === "text/plain") {
        const responseMsg = await (
          err as AxiosError<Blob>
        ).response?.data.text();
        if (responseMsg) {
          msg = responseMsg;
        }
      }

      return new SdkResult<boolean>(undefined, msg);
    }
  }

  async function deleteFile(path: string, options?: RequestOptions) {
    return;
  }

  async function uploadFile(
    path: string,
    filedata: unknown /* idk what this should be yet */,
    options?: RequestOptions
  ) {
    return;
  }

  return {
    listFiles,
    downloadFile,
    uploadFile,
    deleteFile,
  };
}

export default FileStorage;
