import axios, { AxiosInstance } from "axios";
import { Property } from "./models";

export type { Property };

export function PropertyManager(url: string) {
  const client: AxiosInstance = axios.create({
    baseURL: url,
    timeout: 2000,
  });

  async function getProperty(key: string): Promise<string | null> {
    try {
      const resp = await client.get<string>(`/property?key=${key}`);
      return resp.data;
    } catch (err) {
      return null;
    }
  }

  async function setProperty(
    key: string,
    value: string
  ): Promise<string | null> {
    try {
      const resp = await client.post<string>(
        `/property?key=${key}&value=${value}`
      );
      return resp.data;
    } catch (err) {
      return null;
    }
  }

  async function deleteProperty(key: string): Promise<string | null> {
    console.log("deleting key", key);
    try {
      const resp = await client.delete<string>(`/property?key=${key}`);
      return resp.data;
    } catch (err) {
      return null;
    }
  }

  async function getAllProperties(): Promise<Property[] | null> {
    try {
      const resp = await client.get<Property[]>("/");
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
