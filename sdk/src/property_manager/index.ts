import axios, { AxiosInstance } from "axios";
import { Property } from "./models";
import { buildAxiosOptions } from "../common";
import { RequestOptions } from "../common/models";

export type { Property };

export function PropertyManager(url: string) {
  const client: AxiosInstance = axios.create({
    baseURL: url,
    timeout: 2000,
  });

  async function getProperty(
    key: string,
    options?: RequestOptions
  ): Promise<string | null> {
    const config = buildAxiosOptions(options);
    try {
      const resp = await client.get<string>(`/property?key=${key}`, config);
      return resp.data;
    } catch (err) {
      return null;
    }
  }

  async function setProperty(
    key: string,
    value: string,
    options?: RequestOptions
  ): Promise<string | null> {
    const config = buildAxiosOptions(options);
    try {
      const resp = await client.post<string>(
        `/property?key=${key}&value=${value}`,
        config
      );
      return resp.data;
    } catch (err) {
      return null;
    }
  }

  async function deleteProperty(
    key: string,
    options?: RequestOptions
  ): Promise<string | null> {
    const config = buildAxiosOptions(options);
    try {
      const resp = await client.delete<string>(`/property?key=${key}`, config);
      return resp.data;
    } catch (err) {
      return null;
    }
  }

  async function getAllProperties(
    options?: RequestOptions
  ): Promise<Property[] | null> {
    const config = buildAxiosOptions(options);
    try {
      const resp = await client.get<Property[]>("/", config);
      return resp.data;
    } catch (err) {
      return null;
    }
  }

  return {
    getProperty,
    setProperty,
    deleteProperty,
    getAllProperties,
  };
}

export default PropertyManager;
